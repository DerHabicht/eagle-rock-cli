package document

import (
	"fmt"
	"github.com/derhabicht/eagle-rock-lib/lib"
	"strings"
)

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

type HeaderLogo struct {
	Image string `json:"image" yaml:"image"`
	Scale float64 `json:"scale" yaml:"scale"`
}

type MemoHeader struct {
	Logo          HeaderLogo        `json:"logo" yaml:"logo"`
	Address       string            `json:"address" yaml:"address"`
	Tlp           lib.Tlp           `json:"tlp" yaml:"tlp"`
	ControlNumber lib.ControlNumber `json:"control_number" yaml:"control_number"`
	Date          *lib.Date         `json:"date" yaml:"date"`
	Attachments   []string          `json:"attachments,omitempty" yaml:"attachments,omitempty"`
	Cc            []string          `json:"cc,omitempty" yaml:"cc,omitempty"`
}

func (mh MemoHeader) HeaderFieldMap() map[string]interface{} {
	// TODO: Find a more generic way to preserve linebreaks on multi-line strings
	address := strings.Join(strings.Split(mh.Address, "\n"), ` \\ `)
	return map[string]interface{}{
		"LOGO_IMAGE":     mh.Logo.Image,
		"LOGO_SCALE":     fmt.Sprintf("%f", mh.Logo.Scale),
		"ADDRESS":        address,
		"TLP":            mh.Tlp,
		"CONTROL_NUMBER": strings.ReplaceAll(mh.ControlNumber.String(), "-", "--"),
		"DATE":           mh.Date.FormatFormal(),
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
