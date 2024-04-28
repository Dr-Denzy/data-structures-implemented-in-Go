package stack

import "strings"

type CustomStackError struct {
	Msg  string
	Item any
}

func (e *CustomStackError) Error() string {
	return e.Msg
}

func ErrorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}
