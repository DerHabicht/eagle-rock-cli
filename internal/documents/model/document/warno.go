package document

import (
	"fmt"
	lib "github.com/derhabicht/eagle-rock-lib/lib"
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

type WarnoHeader struct {
	Tlp               lib.Tlp
	DateTime          lib.Dtg
	Issuer            string
	ControlNumber     lib.ControlNumber
	MissionNumber     string
	TimeZone          string
	IncidentCommander string
}

func ParseWarnoHeader(s string) (WarnoHeader, error) {
	re := regexp.MustCompile(`(.+)\n(\d{8}T\d{4}(?:(?:Z)|(?:(?:\+|-)\d{4})))\s+(.+)\n(WARNO-\d{2}-\d{3})\s+(.+)\nTZ\s+(UTC(?:\+|-)\d{4})\s+IC\s+(.+)`)

	m := re.FindStringSubmatch(s)

	tlp, err := lib.ParseTlp(m[1])
	if err != nil {
		return WarnoHeader{}, errors.WithMessage(err, "failed to parse WARNO header")
	}
	dtg, err := lib.ParseDtg(m[2])
	if err != nil {
		return WarnoHeader{}, errors.WithMessage(err, "failed to parse WARNO header")
	}
	controlNumber, err := lib.ParseControlNumber(m[4])
	if err != nil {
		return WarnoHeader{}, errors.WithMessage(err, "failed to parse WARNO header")
	}

	return WarnoHeader{
			Tlp:               tlp,
			DateTime:          dtg,
			Issuer:            m[3],
			ControlNumber:     controlNumber,
			MissionNumber:     m[5],
			TimeZone:          m[6],
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

func (wh WarnoHeader) DocumentCN() lib.ControlNumber {
	return wh.ControlNumber
}

func (wh WarnoHeader) DocumentDate() *lib.Date {
	d, err := lib.ParseDate(wh.DateTime.String())
	if err != nil {
		panic(errors.WithStack(err))
	}

	return &d
}

func (wh WarnoHeader) DocumentTitle() string {
	return wh.MissionNumber
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
	header    WarnoHeader
	body      string
	signature WarnoSignature
}

func NewWarno(header WarnoHeader, body string, signature WarnoSignature) Warno {
	return Warno{
		header:    header,
		body:      body,
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
