package css

import "strings"

type SelectorSet []Selector

func NestedSelector(selectors ...Selector) SelectorSet {
	return selectors
}

func (set SelectorSet) Selector() string {
	var selectors []string
	for _, selector := range set {
		selectors = append(selectors, selector.Selector())
	}

	return strings.Join(selectors, " ")
}

func (set SelectorSet) Style(properties ...Property) RuleSet {
	return Rule().For(set).Set(properties...)
}
