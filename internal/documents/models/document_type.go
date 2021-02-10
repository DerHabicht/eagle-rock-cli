// TODO: Move this into the api/pkg/models
package models

import (
	"github.com/pkg/errors"
	"strings"
)

type DocumentType int

const (
	MR DocumentType = iota + 1
	WARNO
	OPORD
	FRAGO
)

func (dt DocumentType) String() string {
	switch dt {
	case MR:
		return "MR"
	case WARNO:
		return "WARNO"
	case OPORD:
		return "OPORD"
	case FRAGO:
		return "FRAGO"
	default:
		panic(errors.Errorf("%d is not a valid document type", dt))
	}
}

func ParseDocumentType(dt string) (DocumentType, error) {
	switch strings.ToLower(dt) {
	case "mr":
		return MR, nil
	case "warno":
		return WARNO, nil
	case "opord":
		return OPORD, nil
	case "frago":
		return FRAGO, nil
	default:
		return 0, errors.Errorf("%s cannot be parsed to a vaild document type", dt)
	}
}
