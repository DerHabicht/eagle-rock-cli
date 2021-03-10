package document

import "github.com/derhabicht/eagle-rock-lib/lib"

type IHeader interface {
	DocumentCN() lib.ControlNumber
	DocumentDate() *lib.Date
	DocumentTitle() string
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

type IDocumentIndex interface {
	GetByClass(class lib.ControlNumberClass) map[int][]lib.ControlNumber
}
