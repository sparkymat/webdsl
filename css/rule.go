package css

import (
	"fmt"
	"strings"
)

type RuleSet struct {
	selectors  []Selector
	properties []Property
}

func For(selectors ...Selector) RuleSet {
	ruleset := RuleSet{}
	ruleset.selectors = append(ruleset.selectors, selectors...)
	return ruleset
}

func (rule RuleSet) Set(properties ...Property) RuleSet {
	rule.properties = append(rule.properties, properties...)
	return rule
}

func (rule RuleSet) String() string {

	selectors := strings.Join(selectorsToStrings(rule.selectors), ",")
	properties := make([]string, 0, 1)
	for _, property := range rule.properties {
		properties = append(properties, property.String())
	}
	propertiesString := strings.Join(properties, "\n")

	return fmt.Sprintf("%v {\n%v\n}\n", selectors, propertiesString)
}
