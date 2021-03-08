package template

import (
	"fmt"
	"github.com/derhabicht/eagle-rock-lib/lib"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"strings"
)

type LatexTemplate struct {
	initialized bool
	keyFormat   string
	template    string
}

func (lt *LatexTemplate) Init(template string) {
	lt.keyFormat = "%%%s%%"
	lt.template = template
	lt.initialized = true
}

func (lt LatexTemplate) Inject(values map[string]interface{}) (string, error) {
	if !lt.initialized {
		return "", errors.New("cannot inject text into uninitialized LaTeX template")
	}

	injected := lt.template
	for k, v := range values {
		if s, ok := v.(string); ok {
			injected = strings.ReplaceAll(injected, fmt.Sprintf(lt.keyFormat, k), s)
		} else if l, ok := v.([]string); ok {
			injected = strings.ReplaceAll(injected, fmt.Sprintf(lt.keyFormat, k), buildLatexList(k, l))
		} else if t, ok := v.(lib.Tlp); ok {
			injected = strings.ReplaceAll(
				injected,
				fmt.Sprintf(lt.keyFormat, "TLP"),
				strings.ToLower(t.LevelString()[4:]),
			)
			caveats := t.CaveatsCsv()
			if caveats != "" {
				injected = strings.ReplaceAll(
					injected,
					fmt.Sprintf(lt.keyFormat, "CAVEATS"),
					fmt.Sprintf(",compartments={%s}", t.CaveatsCsv()),
				)
			}
		} else {
			log.Warn().Msgf("%s: %v is neither a string nor a slice of strings, it will be ignored", k, v)
		}
	}

	return injected, nil
}

func (lt LatexTemplate) Type() string {
	return "latex"
}

func buildLatexList(envName string, items []string) string {
	list := fmt.Sprintf("\\%s{%%\n", strings.ToLower(envName))
	for _, v := range items {
		list += "    \\item " + v + "\n"
	}
	list += "}\n"

	return list
}
