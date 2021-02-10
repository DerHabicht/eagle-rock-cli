package services

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type FileContentReader struct {
	contentDirectory string
}

func NewFileContentReader(contentDirectory string) FileContentReader {
	return FileContentReader{
		contentDirectory: contentDirectory,
	}
}

func (fcr FileContentReader) Read(controlNumber string) ([]byte, error) {
	i := strings.IndexRune(controlNumber, '-')
	year := fmt.Sprintf("20%s", controlNumber[i+1:i+3])

	return ioutil.ReadFile(filepath.Join(fcr.contentDirectory, year, fmt.Sprintf("%s.md", controlNumber)))
}