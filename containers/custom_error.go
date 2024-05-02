package containers

import "strings"

type CustomError struct {
	Msg  string
	Item any
}

func (e *CustomError) Error() string {
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
