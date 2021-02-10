package models

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/yaml.v2"
)

type FragoIndexEntry struct {
	ControlNumber string                        `yaml:"control_number"`
	Uuid          uuid.UUID                     `yaml:"uudi"`
	Date          *Date                         `yaml:"date"`
	Tlp           Tlp                           `yaml:"tlp"`
	Caveats       []string                      `yaml:"caveats"`
	OpordModified uuid.UUID                     `yaml:"opord_modified"`
	StatusHistory map[string]StatusHistoryEntry `yaml:"status_history"`
}

func (fm FragoIndexEntry) DocumentControlNumber() string {
	return fm.ControlNumber
}

func (fm FragoIndexEntry) DocumentUuid() uuid.UUID {
	return fm.Uuid
}

func (fm FragoIndexEntry) DocumentStatusHistory() map[string]StatusHistoryEntry {
	return fm.StatusHistory
}

type FragoHeader struct {
	Tlp               Tlp      `yaml:"tlp"`
	Caveats           []string `yaml:"caveats"`
	ControlNumber     string   `yaml:"control_number"`
	Date              *Date    `yaml:"date"`
	OriginalOpord     string   `yaml:"original_opord"`
	Timezone          string   `yaml:"timezone"`
	IncidentCommander string   `yaml:"incident_commander"`
	References        []string `yaml:"references"`
	Cc                []string `yaml:"cc"`
}

func (fh FragoHeader) FullTlp() string {
	return BuildFullTlpString(fh.Tlp, fh.Caveats)
}

func (fh FragoHeader) Marshal() (string, error) {
	b, err := yaml.Marshal(fh)

	return string(b), err
}

type Frago struct {
	uuid    uuid.UUID
	header  OpordHeader
	Content string
}

func (f Frago) Uuid() uuid.UUID {
	return f.uuid
}

func (f Frago) ControlNumber() string {
	return f.header.ControlNumber
}

func (f Frago) DocumentType() DocumentType {
	return FRAGO
}

func (f Frago) Header() Header {
	return f.header
}
