package css

import "strings"

type CssContainer struct {
	rules []RuleSet
}

func Css(Css string) CssContainer {
	return CssContainer{}
}

func (css CssContainer) String() string {
	var rulesStrings []string
	for _, rule := range css.rules {
		rulesStrings = append(rulesStrings, rule.String())
	}
	return strings.Join(rulesStrings, "\n")
}

func (css CssContainer) Rules(rules ...RuleSet) CssContainer {
	css.rules = append(css.rules, rules...)
	return css
}
