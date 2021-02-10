package models

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type LastAssignedNumber struct {
	Year   int `yaml:"year"`
	Number int `yaml:"number"`
}

type StatusHistoryEntry struct {
	Date  Date   `yaml:"date"`
	Notes string `yaml:"notes"`
}

type DocumentIndex struct {
	filename string                                          `yaml:"-"`
	LastAssignedNumbers map[string]LastAssignedNumber        `yaml:"last_assigned_numbers"`
	Mr                  map[string][]MemoForRecordIndexEntry `yaml:"mr"`
	Warno               map[string][]WarnoIndexEntry         `yaml:"warno"`
	Opord               map[string][]OpordIndexEntry         `yaml:"opord"`
	Frago               map[string][]FragoIndexEntry         `yaml:"frago"`
}

func LoadDocumentIndex(filename string) (DocumentIndex, error) {
	di := DocumentIndex{}

	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return DocumentIndex{}, errors.WithMessagef(err, "failed to read document index file")
	}

	err = yaml.Unmarshal(raw, &di)
	if err != nil {
		return DocumentIndex{}, errors.WithMessage(err, "failed to unmarshal document index")
	}

	return di, nil
}

func (di DocumentIndex) Write() error {
	raw, err := yaml.Marshal(di)
	if err != nil {
		return errors.WithMessagef(err, "failed to marshal the document index")
	}

	err = ioutil.WriteFile(di.filename, raw, 0644)
	if err != nil {
		return errors.WithMessagef(err, "failed to write document index")
	}

	return nil
}
