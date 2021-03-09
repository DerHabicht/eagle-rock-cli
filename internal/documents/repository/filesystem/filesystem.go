package filesystem

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/artifact"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/document"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/template"
	"github.com/derhabicht/eagle-rock-lib/lib"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type FileRepository struct {
	sourceDirectories   map[string]string
	templateDirectories map[string]string
	outputDirectories   map[string]string
}

func NewFileRepository(srcDirs map[string]string, templDirs map[string]string, outDirs map[string]string) FileRepository {
	return FileRepository{
		sourceDirectories:   srcDirs,
		templateDirectories: templDirs,
		outputDirectories:   outDirs,
	}
}

func (fr FileRepository) NewDocument(doc document.IDocument) error {
	return errors.New("NewDocument() not implemented")
}

func (fr FileRepository) LoadDocument(controlNumber lib.ControlNumber) (document.IDocument, error) {
	filename := filepath.Join(
		fr.sourceDirectories[strings.ToLower(controlNumber.Class.String())],
		strconv.Itoa(controlNumber.Year),
		strings.ToLower(controlNumber.String()) + ".md",
	)
	log.Debug().Msgf("Attempting to load document source: %s", filename)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to load document: %s", controlNumber.String())
	}

	switch controlNumber.Class {
	case lib.MR:
		return memoParser(lib.MR, content)
	case lib.OPORD, lib.FRAGO:
		return memoParser(lib.OPORD, content)
	case lib.WARNO:
		return warnoParser(content)
	default:
		return nil, errors.Errorf("unsupported document type: %s", controlNumber.Class.String())
	}
}

func (fr FileRepository) SaveCompiledDocument(controlNumber lib.ControlNumber, doc artifact.BuildArtifact) (string, error) {
	path := filepath.Join(
		fr.outputDirectories[strings.ToLower(controlNumber.Class.String())],
		strconv.Itoa(controlNumber.Year),
	)
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return "", errors.WithMessage(err, "failed to create output directory")
	}
	// TODO: Make this more generic
	filename := strings.ToLower(controlNumber.String() + "." + doc.Type())

	err = ioutil.WriteFile(filepath.Join(path, filename), doc.Artifact(), 0644)
	if err != nil {
		return "", errors.WithMessagef(err, "failed to write %s", controlNumber)
	}

	return filename, nil
}

func (fr FileRepository) LoadTemplate(name string, template template.ITemplate) error {
	filename := filepath.Join(fr.templateDirectories[template.Type()], name + ".template")
	log.Debug().Msgf("Attempting to load template file: %s", filename)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.WithMessagef(err, "failed to load %s template: %s", template.Type(), name)
	}

	template.Init(string(content))

	return nil
}

// FIXME: Handle tables in the Markdown better
func memoParser(class lib.ControlNumberClass, content []byte) (document.IDocument, error) {
	signature := document.MemoSignature{}

	re, err := regexp.Compile(`(?ms)^---(.+?)\.\.\.`)
	if err != nil {
		panic(errors.New("invalid regex for parsing memos"))
	}
	parts := re.FindAllSubmatch(content, -1)
	headerContent := parts[0][1]
	sigContent := parts[1][1]

	err = yaml.Unmarshal(sigContent, &signature)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to parse memo signature")
	}
	body := re.ReplaceAll(content, []byte{})
	switch class {
	case lib.MR:
		header := document.MrHeader{}
		err = yaml.Unmarshal(headerContent, &header)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to parse memo header")
		}
		return document.NewMr(header, string(body), signature), nil
	case lib.OPORD:
		header := document.OpordHeader{}
		err = yaml.Unmarshal(headerContent, &header)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to parse memo header")
		}
		return document.NewOpord(header, string(body), signature), nil
	default:
		return nil, errors.Errorf("memoParser() does not support parsing of %s lib", class.String())
	}
}

func warnoParser(content []byte) (document.IDocument, error) {
	reHdr, err := regexp.Compile(`(?s)(.+?)\n\n`)
	if err != nil {
		panic(errors.New("invalid regex for parsing WARNO headers"))
	}

	headerContent := reHdr.FindSubmatch(content)
	header, err := document.ParseWarnoHeader(strings.Trim(string(headerContent[1]), "\n"))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	reSig, err := regexp.Compile(`//\s+(.+)\s+//`)
	if err != nil {
		panic(errors.New("invalid regex for parsing WARNO signatures"))
	}
	sigContent := reSig.FindSubmatch(content)
	signature := document.WarnoSignature{Name: string(sigContent[1])}

	paragraphs := strings.Split(string(reSig.ReplaceAll(content, []byte{})), "\n\n")

	body := strings.Join(paragraphs[1:], "\n\n")

	return document.NewWarno(header, body, signature), nil
}
