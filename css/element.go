package css

//TODO possibly bad name, but the point is shat Element will implement Selector interface
type Element string

const Html Element = "html"
const Body Element = "body"
const Head Element = "head"
const All Element = "*"

func (selector Element) Selector() string {
	return string(selector)
}

func (selector Element) Style(properties ...Property) RuleSet {
	return Rule().For(selector).Set(properties...)
}
