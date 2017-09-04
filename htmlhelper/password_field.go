package htmlhelper

import (
	. "github.com/kirillrdy/nadeshiko/html"
)

func PasswordField() Node {
	return Input().Type("password").Name("password").Placeholder("Password")
}
