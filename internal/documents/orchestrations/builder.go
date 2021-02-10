package orchestrations

type Builder interface {
	BuildDocument(controlNumber string, content []byte) error
}
