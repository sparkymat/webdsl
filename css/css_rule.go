package css

import (
	"fmt"
	"strings"
)

type CssRuleSet struct {
	selectors  []string
	properties []Property
}

func CssRules() CssRuleSet {
	rule := CssRuleSet{}
	rule.selectors = make([]string, 0, 2)
	rule.properties = make([]Property, 0, 4)

	return rule
}

func (rule CssRuleSet) For(selector string) CssRuleSet {
	rule.selectors = append(rule.selectors, selector)
	return rule
}

func (rule CssRuleSet) Set(properties ...Property) CssRuleSet {
	rule.properties = append(rule.properties, properties...)
	return rule
}

func (rule CssRuleSet) String() string {
	selectors := strings.Join(rule.selectors, ",")
	properties := make([]string, 0, 1)
	for _, property := range rule.properties {
		properties = append(properties, property.String())
	}
	propertiesString := strings.Join(properties, "\n")

	return fmt.Sprintf("%v {\n%v\n}\n", selectors, propertiesString)
}
