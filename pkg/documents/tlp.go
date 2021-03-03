package documents

import (
	"github.com/pkg/errors"
	"strings"
)

type TlpLevel int

const (
	RED TlpLevel = iota + 1
	AMBER
	GREEN
	WHITE
)

func ParseTlpLevel(s string) (TlpLevel, error) {
	switch strings.ToLower(s) {
	case "tlp:red":
		return RED, nil
	case "tlp:amber":
		return AMBER, nil
	case "tlp:green":
		return GREEN, nil
	case "tlp:white":
		return WHITE, nil
	default:
		return -1, errors.Errorf("%s is not a valid TLP level", s)
	}
}

func (tl TlpLevel) String() string {
	switch tl {
	case RED:
		return "TLP:RED"
	case AMBER:
		return "TLP:AMBER"
	case GREEN:
		return "TLP:GREEN"
	case WHITE:
		return "TLP:WHITE"
	default:
		panic(errors.Errorf("%d is not a valid TlpLevel", tl))
	}
}

type Tlp struct {
	level   TlpLevel
	caveats []string
}

func ParseTlp(s string) (Tlp, error) {
	components := strings.Split(s, "//")

	if len(components) > 2 {
		return Tlp{}, errors.Errorf("%s is not a valid TLP string", s)
	}

	level, err := ParseTlpLevel(components[0])
	if err != nil {
		return Tlp{}, errors.WithStack(err)
	}

	var caveats []string
	if len(components) == 2 {
		caveats = strings.Split(strings.ToUpper(components[1]), "/")
	}

	return Tlp{level: level, caveats: caveats}, nil
}

func (t Tlp) String() string {
	if t.caveats != nil {
		return t.level.String() + "//" + strings.Join(t.caveats, "/")
	} else {
		return t.level.String()
	}
}
