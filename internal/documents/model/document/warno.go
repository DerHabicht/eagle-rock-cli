package document

import (
	"fmt"
	"github.com/derhabicht/eagle-rock-cli/pkg/documents"
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

type WarnoHeader struct {
	Tlp documents.Tlp
	DateTime documents.Dtg
	Issuer string
	ControlNumber documents.ControlNumber
	MissionNumber string
	TimeZone string
	IncidentCommander string
}

func ParseWarnoHeader(s string) (WarnoHeader, error) {
	re := regexp.MustCompile(`(.+)\n(\d{8}T\d{4}(?:(?:Z)|(?:(?:\+|-)\d{4})))\s+(.+)\n(WARNO-\d{2}-\d{3})\s+(.+)\nTZ\s+(UTC(?:\+|-)\d{4})\s+IC\s+(.+)`)

	m := re.FindStringSubmatch(s)

	tlp, err := documents.ParseTlp(m[1])
	if err != nil {
		return WarnoHeader{}, errors.WithMessage(err, "failed to parse WARNO header")
	}
	dtg, err := documents.ParseDtg(m[2])
	if err != nil {
		return WarnoHeader{}, errors.WithMessage(err, "failed to parse WARNO header")
	}
	controlNumber, err := documents.ParseControlNumber(m[4])
	if err != nil {
		return WarnoHeader{}, errors.WithMessage(err, "failed to parse WARNO header")
	}

	return WarnoHeader{
		Tlp: tlp,
		DateTime: dtg,
		Issuer: m[3],
		ControlNumber: controlNumber,
		MissionNumber: m[5],
		TimeZone: m[6],
		IncidentCommander: m[7],
	},
	nil
}

func (wh WarnoHeader) String() string {
	return strings.ToUpper(
		fmt.Sprintf(
			"%s\n%s %s\n%s %s\nTZ %s IC %s",
			wh.Tlp.String(),
			wh.DateTime.FormatShort(),
			wh.Issuer,
			wh.ControlNumber.String(),
			wh.MissionNumber,
			wh.TimeZone,
			wh.IncidentCommander,
		),
	)
}

func (wh WarnoHeader) HeaderFieldMap() map[string]interface{} {
	return map[string]interface{}{
		"HEADER": wh.String(),
	}
}

type WarnoSignature struct {
	Name string
}

func (ws WarnoSignature) SignatureFieldMap() map[string]interface{} {
	return map[string]interface{}{
		"SIGNED": "// " + strings.ToUpper(ws.Name) + " //",
	}
}

type Warno struct {
	header WarnoHeader
	body string
	signature WarnoSignature
}

func NewWarno(header WarnoHeader, body string, signature WarnoSignature) Warno {
	return Warno {
		header: header,
		body: body,
		signature: signature,
	}
}

func (w Warno) Header() IHeader {
	return w.header
}

func (w Warno) Body() string {
	return w.body
}

func (w Warno) Signature() ISignature {
	return w.signature
}