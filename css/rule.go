package css

import (
	"fmt"
	"strings"
)

type RuleSet struct {
	selectors  []string
	properties []Property
}

func Rule() RuleSet {
	return RuleSet{}
}

func (rule RuleSet) For(selector Selector) RuleSet {
	rule.selectors = append(rule.selectors, selector.Selector())
	return rule
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
