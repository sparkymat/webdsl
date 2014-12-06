package css

import (
	"fmt"
	"strings"
)

type CssRule struct {
	selectors  []string
	properties []CssProperty
}

func NewCssRule() CssRule {
	rule := CssRule{}
	rule.selectors = make([]string, 0, 2)
	rule.properties = make([]CssProperty, 0, 4)

	return rule
}

func (rule CssRule) For(selector string) CssRule {
	rule.selectors = append(rule.selectors, selector)
	return rule
}

func (rule CssRule) Set(properties ...CssProperty) CssRule {
	rule.properties = append(rule.properties, properties...)
	return rule
}

func (rule CssRule) String() string {
	selectors := strings.Join(rule.selectors, ",")
	properties := make([]string, 0, 1)
	for _, property := range rule.properties {
		properties = append(properties, property.String())
	}
	propertiesString := strings.Join(properties, "\n")

	return fmt.Sprintf("%v {\n%v\n}\n", selectors, propertiesString)
}
