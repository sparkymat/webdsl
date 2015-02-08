package css

type PseudoClass string

const FirstChild PseudoClass = "first-child"
const LastChild PseudoClass = "last-child"
const After PseudoClass = "after"
const Before PseudoClass = "before"
const Hover PseudoClass = "hover"
const Visited PseudoClass = "visited"
const Active PseudoClass = "active"
const Link PseudoClass = "link"
const MozFocusRing PseudoClass = "-moz-focusring"

type SelectorWithPseudoClass struct {
	Element     Selector
	PseudoClass PseudoClass
}

func (selector SelectorWithPseudoClass) Selector() string {
	return selector.Element.Selector() + ":" + string(selector.PseudoClass)
}

func (selector SelectorWithPseudoClass) Style(properties ...Property) RuleSet {
	return For(selector).Set(properties...)
}
