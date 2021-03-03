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
	Logo          string                  `json:"logo" yaml:"logo"`
	Address       string                  `json:"address" yaml:"address"`
	Tlp           documents.Tlp           `json:"tlp" yaml:"tlp"`
	ControlNumber documents.ControlNumber `json:"control_number" yaml:"control_number"`
	Date          documents.Date          `json:"date" yaml:"date"`
	Track         MrTrack                 `json:"track" yaml:"track"`
	Subject       string                  `json:"subject" yaml:"subject"`
	Attachments   []string                `json:"attachments" yaml:"attachments"`
	Cc            []string                `json:"cc" yaml:"cc"`
	StatusHistory MrStatusHistory         `json:"status_history" yaml:"status_history"`
}

func (mh MrHeader) HeaderFieldMap() map[string]interface{} {
	if mh.Track == STANDARDS {
		return map[string]interface{}{
			"LOGO":           mh.Logo,
			"ADDRESS":        mh.Address,
			"TLP":            mh.Tlp.String(),
			"CONTROL_NUMBER": mh.ControlNumber.String(),
			"DATE":           mh.Date.String(),
			"TRACK":          mh.Track.String(),
			"SUBJECT":        mh.Subject,
			"ATTACHMENTS":    mh.Attachments,
			"CC":             mh.Cc,
			"ADOPTED":        mh.StatusHistory.Adopted,
			"REJECTED":	      mh.StatusHistory.Rejected,
		}
	} else {
		return map[string]interface{}{
			"LOGO":           mh.Logo,
			"ADDRESS":        mh.Address,
			"TLP":            mh.Tlp.String(),
			"CONTROL_NUMBER": mh.ControlNumber.String(),
			"DATE":           mh.Date.String(),
			"TRACK":          mh.Track.String(),
			"SUBJECT":        mh.Subject,
			"ATTACHMENTS":    mh.Attachments,
			"CC":             mh.Cc,
		}
	}
}

type MrSignature struct {
	Name string `json:"name" yaml:"name"`
	Signature string `json:"signature" yaml:"signature"`
}

func (ms MrSignature) SignatureFieldMap() map[string]interface{} {
	return map[string]interface{}{
		"NAME": ms.Name,
		"SIGNATURE": ms.Signature,
	}
}

type Mr struct {
	header MrHeader
	body []byte
	signature MrSignature
}

func NewMr(header MrHeader, body []byte, signature MrSignature) Mr {
	return Mr{
		header: header,
		body: body,
		signature: signature,
	}
}

func (m Mr) Header() IHeader {
	return m.header
}

func (m Mr) Body() []byte {
	return m.body
}

func (m Mr) Signature() ISignature {
	return m.signature
}