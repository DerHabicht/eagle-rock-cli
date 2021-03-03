// TODO: Refactor this to eagle-rock-api/pkg
package documents

import (
	"github.com/pkg/errors"
	"time"
)

type Date struct {
	time.Time
}

func ParseDate(s string) (Date, error) {
	t, err := time.Parse("2006-01-02", s)

	return Date{t}, err
}

func (d Date) String() string {
	return d.Format("2006-01-02")
}

func (d Date) MarshalYAML() (interface{}, error) {
	return d.String(), nil
}

func (d *Date) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var buf string
	err := unmarshal(&buf)
	if err != nil {
		return errors.WithStack(err)
	}

	t, err := ParseDate(buf)
	if err != nil {
		return errors.WithStack(err)
	}

	*d = t
	return nil
}

type Dtg struct {
	time.Time
}

func ParseDtg(s string) (Dtg, error) {
	t, err := time.Parse("2006-01-02T15:04Z", s)

	return Dtg{t}, err
}

func (d Dtg) String() string {
	return d.Format("2006-01-02T15:04Z")
}

func (d Dtg) MarshalYAML() (interface{}, error) {
	return d.String(), nil
}

func (d *Dtg) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var buf string
	err := unmarshal(&buf)
	if err != nil {
		return errors.WithStack(err)
	}

	t, err := ParseDtg(buf)
	if err != nil {
		return errors.WithStack(err)
	}

	*d = t
	return nil
}
