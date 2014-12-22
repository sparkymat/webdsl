package form

import "github.com/sparkymat/webdsl/form/rest"

type Form struct {
	action rest.Action

	value interface{}
}

func (form Form) Path() string {
	switch form.action {
	case rest.Create:
		return "/cars"
	default:
		panic("unknown action")
	}
}
