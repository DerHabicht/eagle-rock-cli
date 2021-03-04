package document

import "github.com/derhabicht/eagle-rock-cli/pkg/documents"

type IHeader interface {
	HeaderFieldMap() map[string]interface{}
}

type ISignature interface {
	SignatureFieldMap() map[string]interface{}
}

type IDocument interface {
	Header() IHeader
	Body() string
	Signature() ISignature
}

type MemoHeader struct {
	Logo          string                  `json:"logo" yaml:"logo"`
	Address       string                  `json:"address" yaml:"address"`
	Tlp           documents.Tlp           `json:"tlp" yaml:"tlp"`
	ControlNumber documents.ControlNumber `json:"control_number" yaml:"control_number"`
	Date          *documents.Date         `json:"date" yaml:"date"`
	Attachments   []string                `json:"attachments" yaml:"attachments"`
	Cc            []string                `json:"cc" yaml:"cc"`
}

func (mh MemoHeader) HeaderFieldMap() map[string]interface{} {
	return map[string]interface{}{
		"LOGO":           mh.Logo,
		"ADDRESS":        mh.Address,
		"TLP":            mh.Tlp.String(),
		"CONTROL_NUMBER": mh.ControlNumber.String(),
		"DATE":           mh.Date.String(),
		"ATTACHMENTS":    mh.Attachments,
		"CC":             mh.Cc,
	}
}

type MemoSignature struct {
	Name      string `json:"name" yaml:"name"`
	Signature string `json:"signature" yaml:"signature"`
}

func (ms MemoSignature) SignatureFieldMap() map[string]interface{} {
	return map[string]interface{}{
		"NAME":      ms.Name,
		"SIGNATURE": ms.Signature,
	}
}

