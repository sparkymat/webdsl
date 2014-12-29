package css

//TODO possibly bad name, but the point is shat Element will implement Selector interface
type Element string

const Body Element = "body"
const Head Element = "head"

func (selector Element) Selector() string {
	return string(selector)
}

func (selector Element) Style(properties ...Property) RuleSet {
	return Rule().For(selector).Set(properties...)
}
