package filesystem

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/document"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/template"
	"github.com/derhabicht/eagle-rock-cli/pkg/documents"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"regexp"
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

func (fr FileRepository) LoadDocument(controlNumber documents.ControlNumber) (document.IDocument, error) {
	filename := fr.sourceDirectories[controlNumber.Class.String()] + "/" + controlNumber.String() + ".md"

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to load document: %s", controlNumber.String())
	}

	switch controlNumber.Class {
	case documents.MR:
		return memoParser(documents.MR, content)
	case documents.OPORD, documents.FRAGO:
		return memoParser(documents.OPORD, content)
	case documents.WARNO:
		return warnoParser(content)
	default:
		return nil, errors.Errorf("unsupported document type: %s", controlNumber.Class.String())
	}
}

func (fr FileRepository) SaveCompiledDocument(controlNumber documents.ControlNumber, artifact []byte) error {
	return errors.New("SaveDocument() not implemented")
}

func (fr FileRepository) LoadTemplate(name string, template template.ITemplate) error {
	filename := fr.templateDirectories[template.Type()] + "/" + name + ".template"

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.WithMessagef(err, "failed to load %s template: %s", template.Type(), name)
	}

	template.Init(string(content))

	return nil
}

func memoParser(class documents.ControlNumberClass, content []byte) (document.IDocument, error) {
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
	case documents.MR:
		header := document.MrHeader{}
		err = yaml.Unmarshal(headerContent, &header)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to parse memo header")
		}
		return document.NewMr(header, string(body), signature), nil
	case documents.OPORD:
		header := document.OpordHeader{}
		err = yaml.Unmarshal(headerContent, &header)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to parse memo header")
		}
		return document.NewOpord(header, string(body), signature), nil
	default:
		return nil, errors.Errorf("memoParser() does not support parsing of %s documents", class.String())
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
