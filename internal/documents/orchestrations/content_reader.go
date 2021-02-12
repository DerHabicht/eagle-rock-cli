package orchestrations

type ContentReader interface {
	// Read looks up the document by its control number, and slurps up the file's contents.
	Read(year string, controlNumber string) ([]byte, error)
}
