package css

import "strings"

type SelectorChain []Selector

func NestedSelector(selectors ...Selector) SelectorChain {
	return selectors
}

func (set SelectorChain) Selector() string {
	var selectors []string
	for _, selector := range set {
		selectors = append(selectors, selector.Selector())
	}

	return strings.Join(selectors, " ")
}

func (set SelectorChain) Style(properties ...Property) RuleSet {
	return Rule().For(set).Set(properties...)
}
