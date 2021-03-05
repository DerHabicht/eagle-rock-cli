package document

import "strings"

type OpordHeader struct {
	MemoHeader        `yaml:",inline"`
	MissionNumber     string   `json:"mission_number" yaml:"mission_number"`
	TimeZone          string   `json:"timezone" yaml:"timezone"`
	IncidentCommander string   `json:"ic" yaml:"ic"`
	References        []string `json:"references" yaml:"references"`
}

func (oh OpordHeader) HeaderFieldMap() map[string]interface{} {
	fields := oh.MemoHeader.HeaderFieldMap()

	fields["MISSION_NUMBER"] = strings.ReplaceAll(oh.MissionNumber, "-", "--")
	fields["TIMEZONE"] = oh.TimeZone
	fields["INCIDENT_COMMANDER"] = oh.IncidentCommander

	return fields
}

type Opord struct {
	header    OpordHeader
	body      string
	signature MemoSignature
}

func NewOpord(header OpordHeader, body string, signature MemoSignature) Opord {
	return Opord{
		header:    header,
		body:      body,
		signature: signature,
	}
}

func (o Opord) Header() IHeader {
	return o.header
}

func (o Opord) Body() string {
	return o.body
}

func (o Opord) Signature() ISignature {
	return o.signature
}
