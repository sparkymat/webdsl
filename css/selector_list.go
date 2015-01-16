package css

import "strings"

type SelectorList []Selector

func AllSelectors(selectors ...Selector) SelectorList {
	return selectors
}

func (set SelectorList) Selector() string {
	var selectors []string
	for _, selector := range set {
		selectors = append(selectors, selector.Selector())
	}

	return strings.Join(selectors, ", ")
}

func (set SelectorList) Style(properties ...Property) RuleSet {
	return Rule().For(set).Set(properties...)
}

func (set SelectorList) Nest(selector Selector) SelectorList {
	var nestedSelectors []Selector

	for _, eachSelector := range set {
		nestedSelectors = append(nestedSelectors, NestedSelector(eachSelector, selector))
	}

	return nestedSelectors
}

func (set SelectorList) NestAll(nestedSet SelectorList) SelectorList {
	var nestedSelectors []Selector

	for _, eachSelector := range set {
		for _, eachNestedSelector := range nestedSet {
			nestedSelectors = append(nestedSelectors, NestedSelector(eachSelector, eachNestedSelector))
		}
	}

	return nestedSelectors
}
