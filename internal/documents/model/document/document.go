package document

type IHeader interface {
	HeaderFieldMap() map[string]interface{}
}

type ISignature interface {
	SignatureFieldMap() map[string]interface{}
}

type IDocument interface {
	Header() IHeader
	Body() []byte
	Signature() ISignature
}
