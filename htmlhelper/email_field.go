package htmlhelper

import (
	. "github.com/kirillrdy/nadeshiko/html"
)

func EmailField() Node {
	return Input().Type("email").Name("email").Placeholder("E-mail")
}
