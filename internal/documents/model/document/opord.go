package document

import (
	"github.com/derhabicht/eagle-rock-lib/lib"
	"strings"
)

type OpordHeader struct {
	MemoHeader        `yaml:",inline"`
	MissionNumber     string   `json:"mission_number" yaml:"mission_number"`
	TimeZone          string   `json:"timezone" yaml:"timezone"`
	IncidentCommander string   `json:"ic" yaml:"ic"`
	References        []string `json:"references" yaml:"references"`
}

func (oh OpordHeader) HeaderFieldMap() map[string]interface{} {
	fields := oh.MemoHeader.HeaderFieldMap()

	if oh.ControlNumber.Class == lib.OPORD {
		fields["OPORD"] = "OPERATION ORDER " + strings.ReplaceAll(oh.ControlNumber.String()[6:], "-", "--")
	} else if oh.ControlNumber.Class == lib.FRAGO {
		fields["OPORD"] = "FRAGMENTARY OPERATION ORDER " + oh.ControlNumber.String()[6:]
	}
	fields["MISSION_NUMBER"] = strings.ReplaceAll(oh.MissionNumber, "-", "--")
	fields["TIMEZONE"] = oh.TimeZone
	fields["INCIDENT_COMMANDER"] = oh.IncidentCommander
	fields["REFERENCES"] = oh.References

	return fields
}

func (oh OpordHeader) DocumentTitle() string {
	return oh.MissionNumber
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
