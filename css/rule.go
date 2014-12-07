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
	rule := RuleSet{}
	rule.selectors = make([]string, 0, 2)
	rule.properties = make([]Property, 0, 4)

	return rule
}

func (rule RuleSet) For(selector string) RuleSet {
	rule.selectors = append(rule.selectors, selector)
	return rule
}

func (rule RuleSet) Set(properties ...Property) RuleSet {
	rule.properties = append(rule.properties, properties...)
	return rule
}

func (rule RuleSet) String() string {
	selectors := strings.Join(rule.selectors, ",")
	properties := make([]string, 0, 1)
	for _, property := range rule.properties {
		properties = append(properties, property.String())
	}
	propertiesString := strings.Join(properties, "\n")

	return fmt.Sprintf("%v {\n%v\n}\n", selectors, propertiesString)
}
