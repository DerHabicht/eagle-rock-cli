package content_readers

import (
	"io/ioutil"
	"path/filepath"
)

type FileContentReader struct {
	contentDirectory string
	fileExtension string
}

func NewFileContentReader(contentDirectory string, fileExtension string) FileContentReader {
	return FileContentReader{
		contentDirectory: contentDirectory,
		fileExtension: fileExtension,
	}
}

func (fcr FileContentReader) Read(year string, controlNumber string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join( fcr.contentDirectory, year, controlNumber + fcr.fileExtension))
}