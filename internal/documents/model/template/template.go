package template

import "github.com/derhabicht/eagle-rock-cli/internal/documents/model/document"

type ITemplate interface {
	Inject(document.IDocument) ([]byte, error)
}
