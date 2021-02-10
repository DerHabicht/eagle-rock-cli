package models

import (
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/yaml.v2"
	"strings"
)

type MemoForRecordTrack int

const (
	STANDARDS MemoForRecordTrack = iota + 1
	PROGRAM
	PROJECT
	ADVISORY
)

func (mfrt MemoForRecordTrack) String() string {
	switch mfrt {
	case STANDARDS:
		return "STANDARDS"
	case PROGRAM:
		return "PROGRAM"
	case PROJECT:
		return "PROJECT"
	case ADVISORY:
		return "ADVISORY"
	default:
		panic(errors.Errorf("invalid MemoForRecordTrack value: %d", mfrt))
	}
}

func ParseMemoForRecordTrack(s string) (MemoForRecordTrack, error) {
	switch strings.ToLower(s) {
	case "standards":
		return STANDARDS, nil
	case "program":
		return PROGRAM, nil
	case "project":
		return PROJECT, nil
	case "advisory":
		return ADVISORY, nil
	default:
		return 0, errors.Errorf("invalid value for MemoForRecordsTrack: %s", s)
	}
}

func (mfrt MemoForRecordTrack) MarshalYAML() (interface{}, error) {
	return mfrt.String(), nil
}

func (mfrt *MemoForRecordTrack) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var buf string
	err := unmarshal(&buf)
	if err != nil {
		return errors.WithStack(err)
	}

	t, err := ParseMemoForRecordTrack(buf)
	if err != nil {
		return errors.WithStack(err)
	}

	*mfrt = t
	return nil
}

type MemoForRecordIndexEntry struct {
	Uuid          uuid.UUID                     `yaml:"uuid"`
	ControlNumber string                        `yaml:"control_number"`
	Date          *Date                         `yaml:"date"`
	Subject       string                        `yaml:"subject"`
	Tlp           Tlp                           `yaml:"tlp"`
	Caveats       []string                      `yaml:"caveats"`
	Track         MemoForRecordTrack            `yaml:"track"`
	StatusHistory map[string]StatusHistoryEntry `yaml:"status_history"`
}

func (mfrm MemoForRecordIndexEntry) DocumentUuid() uuid.UUID {
	return mfrm.Uuid
}

func (mfrm MemoForRecordIndexEntry) DocumentControlNumber() string {
	return mfrm.ControlNumber
}

func (mfrm MemoForRecordIndexEntry) DocumentStatusHistory() map[string]StatusHistoryEntry {
	return mfrm.StatusHistory
}

type MemoForRecordHeader struct {
	Tlp           Tlp      `yaml:"tlp" latex:"TLP"`
	Caveats       []string `yaml:"caveats" latex:"-"`
	ControlNumber string   `yaml:"control_number" latex:"CONTROL_NUMBER"`
	Date          *Date    `yaml:"date" latex:"DATE"`
	MemoFor       string   `yaml:"memo_for" latex:"MEMO_FOR"`
	MemoFrom      string   `yaml:"memo_from" latex:"MEMO_FROM"`
	Subject       string   `yaml:"subject" latex:"SUBJECT"`
	Attachments   []string `yaml:"attachments" latex:"ATTACHMENTS,attachments"`
	Cc            []string `yaml:"cc" latex:"CC,cc"`
}

func (mrh MemoForRecordHeader) FullTlp() string {
	return BuildFullTlpString(mrh.Tlp, mrh.Caveats)
}

func (mrh MemoForRecordHeader) Marshal() (string, error) {
	b, err := yaml.Marshal(mrh)

	return string(b), err
}

type MemoForRecord struct {
	uuid    uuid.UUID
	header  MemoForRecordHeader
	Content string
}

func (mr MemoForRecord) Uuid() uuid.UUID {
	return mr.uuid
}

func (mr MemoForRecord) ControlNumber() string {
	return mr.header.ControlNumber
}

func (mr MemoForRecord) DocumentType() DocumentType {
	return MR
}

func (mr MemoForRecord) Header() Header {
	return mr.header
}
