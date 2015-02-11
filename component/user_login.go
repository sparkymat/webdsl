package component

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/size"
	"github.com/sparkymat/webdsl/http"
)

const userLoginViewClass css.Class = "user-login-card"

type UserLoginView struct {
	Action string
	Method http.Method
}

func (card UserLoginView) Style() css.CssContainer {
	return css.Stylesheet(
		css.For(userLoginViewClass).Set(
			css.Width(card.Width()),
			css.Height(card.Height()),
		),
	)
}

func (card UserLoginView) Html() html.Node {
	return html.Div().Class(userLoginViewClass).Children(
		html.Form().Action(card.Action).Method(string(card.Method)).Children(
			html.Input().Name("email").Placeholder("E-mail").Type("email"),
			html.Input().Name("password").Placeholder("Password").Type("password"),
			html.Input().Type("submit").Value("Login"),
		),
	)
}

func (card UserLoginView) Width() size.Size {
	return size.Px(640)
}

func (card UserLoginView) Height() size.Size {
	return size.Px(360)
}
