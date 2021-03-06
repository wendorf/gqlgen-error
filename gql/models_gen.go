// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"fmt"
	"io"
	"strconv"
)

type Setting string

const (
	SettingDefault Setting = "DEFAULT"
)

var AllSetting = []Setting{
	SettingDefault,
}

func (e Setting) IsValid() bool {
	switch e {
	case SettingDefault:
		return true
	}
	return false
}

func (e Setting) String() string {
	return string(e)
}

func (e *Setting) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Setting(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Setting", str)
	}
	return nil
}

func (e Setting) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
