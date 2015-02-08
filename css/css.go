package css

import "strings"

type CssContainer struct {
	rules []RuleSet
}

func Stylesheet(rules ...RuleSet) CssContainer {
	container := CssContainer{}
	container.rules = append(container.rules, rules...)
	return container
}

func (css CssContainer) String() string {
	var rulesStrings []string
	for _, rule := range css.rules {
		rulesStrings = append(rulesStrings, rule.String())
	}
	return strings.Join(rulesStrings, "\n")
}
