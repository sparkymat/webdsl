package css

import "fmt"

type Element string
type ElementWithAttribute string

func (selector Element) Selector() string {
	return string(selector)
}

func (selector Element) WithPseudoClass(pseudoClass PseudoClass) SelectorWithPseudoClass {
	return SelectorWithPseudoClass{Element: selector, PseudoClass: pseudoClass}
}

func (selector Element) Style(properties ...Property) RuleSet {
	return For(selector).Set(properties...)
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
	return For(selector).Set(properties...)
}
