package form

import (
	"errors"
	"regexp"
	"strings"

	"github.com/sparkymat/webdsl/form/rest"
)

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

func (form Form) Method() string {
	switch form.action {
	case rest.Create:
		return "POST"
	default:
		panic("unknown action")
	}
}

func CamelToUnderscore(camel string) (string, error) {
	regex, err := regexp.Compile("^[a-zA-Z]+[a-zA-Z0-9]*$")

	if err != nil {
		return "", err
	}

	if regex.MatchString(camel) != true {
		return "", errors.New("Invalid input string")
	}

	regex, err = regexp.Compile("[A-Z]")

	if err != nil {
		return "", err
	}

	stage1 := string(regex.ReplaceAll([]byte(camel), []byte("_$0")))

	regex, err = regexp.Compile("^_+")

	if err != nil {
		return "", err
	}

	stage2 := regex.ReplaceAllString(stage1, "")

	return strings.ToLower(stage2), nil
}
