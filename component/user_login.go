package component

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/size"
	"github.com/sparkymat/webdsl/http"
)

const userLoginCardClass css.Class = "user-login-card"

type UserLoginCard struct {
	Action string
	Method http.Method
}

func (card UserLoginCard) Style() []css.RuleSet {
	var rules []css.RuleSet

	rules = append(rules, css.Rule().For(userLoginCardClass).Set(
		css.Width(card.Width()),
		css.Height(card.Height()),
	))

	return rules
}

func (card UserLoginCard) Html() html.Node {
	return html.Div().Class(userLoginCardClass).Children(
		html.Form().Action(card.Action).Method(string(card.Method)).Children(
			html.Input().Name("email").Placeholder("E-mail").Type("email"),
			html.Input().Name("password").Placeholder("Password").Type("password"),
			html.Input().Type("submit").Value("Login"),
		),
	)
}

func (card UserLoginCard) Type() ComponentType {
	return FixedSize
}

func (card UserLoginCard) Width() size.Size {
	return size.Px(640)
}

func (card UserLoginCard) Height() size.Size {
	return size.Px(360)
}
