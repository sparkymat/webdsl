package css

import "strings"

type AdjacentSelectors []Selector

func Adjacents(selectors ...Selector) AdjacentSelectors {
	return selectors
}

func (adjacents AdjacentSelectors) Selector() string {
	var selectors []string
	for _, selector := range adjacents {
		selectors = append(selectors, selector.Selector())
	}

	return strings.Join(selectors, " + ")
}

func (adjacents AdjacentSelectors) Style(properties ...Property) RuleSet {
	return Rule().For(adjacents).Set(properties...)
}
