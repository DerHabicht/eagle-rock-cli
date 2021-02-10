package models

import (
	"fmt"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"regexp"
	"strings"
	"time"
)

type WarnoIndexEntry struct {
	ControlNumber string                        `yaml:"control_number"`
	Uuid          uuid.UUID                     `yaml:"uuid"`
	Dtg           *Dtg                          `yaml:"dtg"`
	Tlp           Tlp                           `yaml:"tlp"`
	Caveats       []string                      `yaml:"caveats"`
	Mission       string                        `yaml:"mission"`
	StatusHistory map[string]StatusHistoryEntry `yaml:"status_history"`
}

func (wm WarnoIndexEntry) DocumentUuid() uuid.UUID {
	return wm.Uuid
}

func (wm WarnoIndexEntry) DocumentControlNumber() string {
	return wm.ControlNumber
}

func (wm WarnoIndexEntry) DocumentStatusHistory() map[string]StatusHistoryEntry {
	return wm.StatusHistory
}

type WarnoHeader struct {
	Tlp               Tlp
	Caveats           []string
	DateTime          time.Time
	Issuer            string
	ControlNumber     string
	Mission           string
	TimeZone          string
	IncidentCommander string
}

func (wh WarnoHeader) FullTlp() string {
	return BuildFullTlpString(wh.Tlp, wh.Caveats)
}

func (wh WarnoHeader) Marshal() (string, error) {
	s := fmt.Sprintf(
		"%s\n%s %s\n%s %s\nTZ %s IC %s",
		BuildFullTlpString(wh.Tlp, wh.Caveats),
		wh.DateTime.Format("20060102T1504Z"),
		wh.Issuer,
		wh.ControlNumber,
		wh.Mission,
		wh.TimeZone,
		strings.ToUpper(wh.IncidentCommander),
	)

	return s, nil
}

func UnmarshalWarnoHeader(s string) (WarnoHeader, error) {
	re := regexp.MustCompile(`(.+)\n(\d{8}T\d{4}Z)\s(.+)\n(.+)\s(.+)\nTZ (.+) IC (.+)`)

	matches := re.FindStringSubmatch(s)

	if len(matches) != 8 {
		return WarnoHeader{}, errors.Errorf(
			"WARNO header should have 7 fields, but %d were found:\n%#v\nRaw header was:\n%s)",
			len(matches)-1,
			matches,
			s,
		)
	}

	tlp, caveats, err := ParseFullTlpString(matches[1])
	if err != nil {
		return WarnoHeader{}, errors.WithMessagef(err, "invalid TLP field in WARNO header: %s", s)
	}

	dt, err := time.Parse("20060102T1504Z", matches[2])
	if err != nil {
		return WarnoHeader{}, errors.WithMessagef(err, "invalid DTG in WARNO header: %s", s)
	}

	return WarnoHeader{
			Tlp:               tlp,
			Caveats:           caveats,
			DateTime:          dt,
			Issuer:            matches[3],
			ControlNumber:     matches[4],
			Mission:           matches[5],
			TimeZone:          matches[6],
			IncidentCommander: strings.Title(strings.ToLower(matches[7])),
		},
		nil
}

type Warno struct {
	uuid   uuid.UUID
	header WarnoHeader
}

func (w Warno) Uuid() uuid.UUID {
	return w.uuid
}

func (w Warno) ControlNumber() string {
	return w.header.ControlNumber
}

func (w Warno) DocumentType() DocumentType {
	return WARNO
}

func (w Warno) Header() Header {
	return w.header
}
