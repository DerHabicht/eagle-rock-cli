package document

import "github.com/derhabicht/eagle-rock-cli/pkg/documents"

type MrStatus struct {
	Date  documents.Date `json:"date" yaml:"date"`
	Notes string         `json:"notes" yaml:"notes"`
}

type MrStatusHistory struct {
	Registered MrStatus  `json:"registered" yaml:"registered"`
	Published  *MrStatus `json:"published,omitempty" yaml:"published,omitempty"`
	Adopted    *MrStatus `json:"adopted,omitempty" yaml:"adopted,omitempty"`
	Rejected   *MrStatus `json:"rejected,omitempty" yaml:"rejected,omitempty"`
}

type MrHeader struct {
	MemoHeader    `yaml:",inline"`
	From          string          `json:"from" yaml:"from"`
	Track         MrTrack         `json:"track" yaml:"track"`
	Subject       string          `json:"subject" yaml:"subject"`
	StatusHistory MrStatusHistory `json:"status_history" yaml:"status_history"`
}

func (mh MrHeader) HeaderFieldMap() map[string]interface{} {
	fields := mh.MemoHeader.HeaderFieldMap()
	fields["FROM"] = mh.From
	fields["TRACK"] = mh.Track.String() + " TRACK"
	fields["SUBJECT"] = mh.Subject

	if mh.Track == STANDARDS {
		fields["ADOPTED"] = mh.StatusHistory.Adopted
		fields["REJECTED"] = mh.StatusHistory.Rejected
	}

	return fields
}

type Mr struct {
	header    MrHeader
	body      string
	signature MemoSignature
}

func NewMr(header MrHeader, body string, signature MemoSignature) Mr {
	return Mr{
		header:    header,
		body:      body,
		signature: signature,
	}
}

func (m Mr) Header() IHeader {
	return m.header
}

func (m Mr) Body() string {
	return m.body
}

func (m Mr) Signature() ISignature {
	return m.signature
}
