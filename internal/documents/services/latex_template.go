package services

import (
	"bytes"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/models"
	"github.com/pkg/errors"
	"reflect"
	"strconv"
	"strings"
)

type LatexTemplate struct {
	template []byte
}

func NewLatexTemplate(template []byte) LatexTemplate {
	return LatexTemplate{
		template: template,
	}
}

// TODO: Handle the status table for Standards-Track MRs
func (lt LatexTemplate) Inject(header models.Header, text []byte) ([]byte, error) {
	h, err := processHeaderFields(header)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to prepare header information for template")
	}

	// TODO: Find a more elegant way to handle this.
	h["%CONTROL_NUMBER%"] = bytes.ReplaceAll(h["%CONTROL_NUMBER%"], []byte("-"), []byte("--"))

	templ := lt.template
	for k, v := range h {
		templ = bytes.ReplaceAll(templ, []byte(k), v)
	}

	templ = bytes.ReplaceAll(templ, []byte("%TEXT%"), text)

	return templ, nil
}

func processListHeaderField(command string, field []string) []byte {
	if field == nil {
		return nil
	}

	list := `\` + command + "{%\n"

	for _, v := range field {
		list += `\item ` + v + "\n"
	}

	list += "}"

	return []byte(list)
}

func processHeaderFields(header models.Header) (map[string][]byte, error) {
	p := make(map[string][]byte)

	t := reflect.TypeOf(header)
	v := reflect.ValueOf(header)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("latex") != "-" {
			// TODO: Review and clean this disaster up
			if t.Field(i).Type == reflect.TypeOf([]string{}) {
				s := strings.Split(t.Field(i).Tag.Get("latex"), ",")
				key := "%" + s[0] + "%"
				p[key] = processListHeaderField(s[1], v.Field(i).Interface().([]string))
			} else {
				key := "%" + t.Field(i).Tag.Get("latex") + "%"
				switch reflect.TypeOf(v.Field(i).Interface()) {
				case reflect.TypeOf(27):
					p[key] = []byte(strconv.FormatInt(v.Field(i).Int(), 10))
				case reflect.TypeOf(""):
					p[key] = []byte(v.Field(i).String())
				case reflect.TypeOf(models.RED):
					t := v.Field(i).Interface().(models.Tlp)
					p[key] = []byte(strings.ToLower(t.String()))
				case reflect.TypeOf(&models.Date{}):
					d := v.Field(i).Interface().(*models.Date)
					p[key] = []byte(d.Format("02 January 2006"))
				default:
					return nil, errors.Errorf("%s does not have a supported type", t.Field(i).Name)

				}
			}
		}
	}

	return p, nil
}