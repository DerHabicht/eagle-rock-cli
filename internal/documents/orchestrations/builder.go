package orchestrations

type Builder interface {
	BuildDocument(year string, controlNumber string, content []byte) error
}
