package services

import (
	"fmt"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/document"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/repository"
	"github.com/derhabicht/eagle-rock-lib/lib"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"path/filepath"
	"strconv"
	"strings"
)

type Indexer struct {
	repository repository.IRepository
}

func NewIndexer(repo repository.IRepository) Indexer {
	return Indexer{
		repository: repo,
	}
}

// TODO: Consider a less ugly data structure
func (i Indexer) IndexDocuments() (map[string]map[int]string, error) {
	docList, err := i.repository.ListDocuments()
	if err != nil {
		return nil, errors.WithStack(err)
	}

}

func loadDocumentData(repo repository.IRepository, docList document.IDocumentIndex) (string, error) {

}

func createDocumentClassEntry(class lib.ControlNumberClass, hdrs map[int][]document.IHeader, path string) {
	var section string
	switch class {
	case lib.MR:
		section = "## Memoranda for Record\n\n"
	case lib.WARNO:
		section = "## Warn Orders\n\n"
	case lib.OPORD:
		section = "## Operation Orders\n\n"
	case lib.FRAGO:
		section = "## Fragmentary Operation Orders\n\n"
	default:
		panic(errors.Errorf("%d is not a vaild ControlNumberClass", class))
	}

	for year, docs := range hdrs {
		section += createYearEntry(year, docs, filepath.Join(path, strconv.Itoa(year)))
	}
}

func createYearEntry(year int, hdrs []document.IHeader, path string) string {
	section := "### " + strconv.Itoa(year) + "\n"

	for _, hdr := range hdrs {
		section += createDocumentEntry(hdr, filepath.Join(path, strings.ToLower(hdr.DocumentCN().String()) + ".md"))
	}

}

func createDocumentEntry(hdr document.IHeader, path string) string {
	format := "- [%s](%s): %s (%s)\n"

	if hdr.DocumentDate() != nil {
		return fmt.Sprintf(format, hdr.DocumentCN(), path, hdr.DocumentTitle(), "UNPUBLISHED")
	} else {
		return fmt.Sprintf(format, hdr.DocumentCN(), path, hdr.DocumentTitle(), *hdr.DocumentDate())
	}
}