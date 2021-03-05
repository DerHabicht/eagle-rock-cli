package template

import (
	"fmt"
	"github.com/pkg/errors"
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
		s, ok := v.(string)
		if ok {
			injected = strings.ReplaceAll(injected, fmt.Sprintf(lt.keyFormat, k), s)
		} else {
			l, ok := v.([]string)
			if ok {
				injected = strings.ReplaceAll(injected, fmt.Sprintf(lt.keyFormat, k), buildLatexList(k, l))
			} else {
				return "", errors.Errorf("%s: %v is neither a string nor a slice of strings", k, v)
			}
		}
	}

	return injected, nil
}

func (lt LatexTemplate) Type() string {
	return "LATEX"
}

func buildLatexList(envName string, items []string) string {
	list := fmt.Sprintf("\\%s{%%\n", strings.ToLower(envName))
	for _, v := range items {
		list += "    \\item " + v + "\n"
	}
	list += "}\n"

	return list
}
