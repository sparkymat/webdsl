package css

import "strings"

type AdjacentSelectors []Selector

func Adjacents(selectors ...Selector) AdjacentSelectors {
	return selectors
}

func (adjacents AdjacentSelectors) Selector() string {
	return strings.Join(selectorsToStrings(adjacents), " + ")
}

func (adjacents AdjacentSelectors) Style(properties ...Property) RuleSet {
	return Rule().For(adjacents).Set(properties...)
}
