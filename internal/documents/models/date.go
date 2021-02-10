// TODO: Refactor this to eagle-rock-api/pkg
package models

import (
	"github.com/pkg/errors"
	"time"
)

type Date struct {
	time.Time
}

func (d Date) MarshalYAML() (interface{}, error) {
	return d.Format("2006-01-02"), nil
}

func (d *Date) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var buf string
	err := unmarshal(&buf)
	if err != nil {
		return errors.WithStack(err)
	}

	t, err := time.Parse("2006-01-02", buf)
	if err != nil {
		return errors.WithStack(err)
	}

	*d = Date{t}
	return nil
}

type Dtg struct {
	time.Time
}

func (d Dtg) MarshalYAML() (interface{}, error) {
	return d.Format("2006-01-02T15:04Z"), nil
}

func (d *Dtg) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var buf string
	err := unmarshal(&buf)
	if err != nil {
		return errors.WithStack(err)
	}

	t, err := time.Parse("2006-01-02T15:04Z", buf)
	if err != nil {
		return errors.WithStack(err)
	}

	*d = Dtg{t}
	return nil
}
