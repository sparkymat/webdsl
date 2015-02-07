package css

import "strings"

type SelectorChild []Selector

func ChildSelector(selectors ...Selector) SelectorChain {
	return selectors
}

func (set SelectorChild) Selector() string {
	return strings.Join(selectorsToStrings(set), " > ")
}

func (set SelectorChild) Style(properties ...Property) RuleSet {
	return For(set).Set(properties...)
}
