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

func (d Date) FormatFormal() string {
	return d.Format("02 January 2006")
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

	temp, err := ParseDate(buf)
	if err != nil {
		return errors.WithStack(err)
	}

	*d = temp
	return nil
}

type Dtg struct {
	time.Time
}

func ParseDtg(s string) (Dtg, error) {
	var t Dtg

	t, err := ParseLongDtg(s)
	if err != nil {
		t, err = ParseShortDtg(s)
		if err != nil {
			return Dtg{}, errors.WithMessage(err, "%s is neither in the long nor the short format")
		}
	}

	return t, nil
}

func ParseLongDtg(s string) (Dtg, error) {
	var t time.Time
	var err error

	if s[len(s)-1:] == "Z" {
		t, err = time.Parse("2006-01-02T15:04Z", s)
	} else {
		t, err = time.Parse("2006-01-02T15:04-07:00", s)
	}

	return Dtg{t}, err
}

func ParseShortDtg(s string) (Dtg, error) {
	var t time.Time
	var err error

	if s[len(s)-1:] == "Z" {
		t, err = time.Parse("20060102T1504Z", s)
	} else {
		t, err = time.Parse("20060102T1504-0700", s)
	}

	return Dtg{t}, err
}

func (d Dtg) FormatLong() string {
	_, tz := d.Zone()
	if tz == 0 {
		return d.Format("2006-01-02T15:04Z")
	} else {
		return d.Format("2006-01-02T15:04-07:00")
	}
}

func (d Dtg) FormatShort() string {
	_, tz := d.Zone()
	if tz == 0 {
		return d.Format("20060102T1504Z")
	} else {
		return d.Format("20060102T1504-0700")
	}
}

func (d Dtg) String() string {
	return d.FormatLong()
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

	temp, err := ParseDtg(buf)
	if err != nil {
		return errors.WithStack(err)
	}

	*d = temp
	return nil
}
