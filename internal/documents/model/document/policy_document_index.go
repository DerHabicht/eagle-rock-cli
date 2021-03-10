package document

import (
	"github.com/derhabicht/eagle-rock-lib/lib"
	"github.com/pkg/errors"
)

type PolicyDocumentIndex struct {
	mrs    map[int][]lib.ControlNumber
	warnos map[int][]lib.ControlNumber
	opords map[int][]lib.ControlNumber
	fragos map[int][]lib.ControlNumber
}

func NewPolicyDocumentIndex(
	mrs map[int][]lib.ControlNumber,
	warnos map[int][]lib.ControlNumber,
	opords map[int][]lib.ControlNumber,
	fragos map[int][]lib.ControlNumber,
) PolicyDocumentIndex {
	return PolicyDocumentIndex{
		mrs:    mrs,
		warnos: warnos,
		opords: opords,
		fragos: fragos,
	}
}

func (pdi PolicyDocumentIndex) GetByClass(class lib.ControlNumberClass) map[int][]lib.ControlNumber {
	switch class {
	case lib.MR:
		return pdi.mrs
	case lib.WARNO:
		return pdi.warnos
	case lib.OPORD:
		return pdi.opords
	case lib.FRAGO:
		return pdi.fragos
	default:
		panic(errors.Errorf("%v is not a valid ControlNumberClass", class))
	}
}
