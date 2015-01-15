package css

type PseudoClass string

const FirstChild PseudoClass = "first-child"
const LastChild PseudoClass = "last-child"
const After PseudoClass = "after"
const Before PseudoClass = "before"

type SelectorWithPseudoClass struct {
	Element     Selector
	PseudoClass PseudoClass
}

func (selector SelectorWithPseudoClass) Selector() string {
	return selector.Element.Selector() + ":" + string(selector.PseudoClass)
}

func (selector SelectorWithPseudoClass) Style(properties ...Property) RuleSet {
	return Rule().For(selector).Set(properties...)
}
