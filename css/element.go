package css

import "fmt"

//TODO possibly bad name, but the point is shat Element will implement Selector interface
type Element string
type ElementWithAttribute string

const Html Element = "html"
const Body Element = "body"
const Head Element = "head"
const All Element = "*"

const H1 Element = "h1"
const H2 Element = "h2"
const H3 Element = "h3"
const H4 Element = "h4"
const H5 Element = "h5"

const A Element = "a"
const P Element = "p"

const Ul Element = "ul"
const Strong Element = "strong"
const Input Element = "input"
const Button Element = "button"
const Select Element = "select"

func (selector Element) Selector() string {
	return string(selector)
}

func (selector Element) WithPseudoClass(pseudoClass PseudoClass) SelectorWithPseudoClass {
	return SelectorWithPseudoClass{Element: selector, PseudoClass: pseudoClass}
}

func (selector Element) Style(properties ...Property) RuleSet {
	return Rule().For(selector).Set(properties...)
}

func (selector Element) WithAttribute(value string) ElementWithAttribute {
	return ElementWithAttribute(fmt.Sprintf("%v[%v]", selector, value))
}

func (selector ElementWithAttribute) Selector() string {
	return string(selector)
}

func (selector ElementWithAttribute) WithPseudoClass(pseudoClass PseudoClass) SelectorWithPseudoClass {
	return SelectorWithPseudoClass{Element: selector, PseudoClass: pseudoClass}
}

func (selector ElementWithAttribute) Style(properties ...Property) RuleSet {
	return Rule().For(selector).Set(properties...)
}
