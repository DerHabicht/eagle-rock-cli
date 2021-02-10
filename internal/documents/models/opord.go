package models

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/yaml.v2"
)

type OpordIndexEntry struct {
	ControlNumber string                        `yaml:"control_number"`
	Uuid          uuid.UUID                     `yaml:"uuid"`
	Date          *Date                         `yaml:"date"`
	Tlp           Tlp                           `yaml:"tlp"`
	Caveats       []string                      `yaml:"caveats"`
	Mission       string                        `yaml:"mission"`
	StatusHistory map[string]StatusHistoryEntry `yaml:"status_history"`
}

func (om OpordIndexEntry) DocumentControlNumber() string {
	return om.ControlNumber
}

func (om OpordIndexEntry) DocumentUuid() uuid.UUID {
	return om.Uuid
}

func (om OpordIndexEntry) DocumentStatusHistory() map[string]StatusHistoryEntry {
	return om.StatusHistory
}

type OpordHeader struct {
	Tlp               Tlp      `yaml:"tlp"`
	Caveats           []string `yaml:"caveats"`
	ControlNumber     string   `yaml:"control_number"`
	Date              *Date    `yaml:"date"`
	Mission           string   `yaml:"mission"`
	Timezone          string   `yaml:"timezone"`
	IncidentCommander string   `yaml:"incident_commander"`
	References        []string `yaml:"references"`
	Cc                []string `yaml:"cc"`
}

func (oh OpordHeader) FullTlp() string {
	return BuildFullTlpString(oh.Tlp, oh.Caveats)
}

func (oh OpordHeader) Marshal() (string, error) {
	b, err := yaml.Marshal(oh)

	return string(b), err
}

type Opord struct {
	uuid    uuid.UUID
	header  OpordHeader
	Content string
}

func (o Opord) Uuid() uuid.UUID {
	return o.uuid
}

func (o Opord) ControlNumber() string {
	return o.header.ControlNumber
}

func (o Opord) DocumentType() DocumentType {
	return OPORD
}

func (o Opord) Header() Header {
	return o.header
}
