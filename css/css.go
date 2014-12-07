package css

import "strings"

type CssContainer struct {
	fileName string
	rules    []RuleSet
}

func Css(Css string) CssContainer {
	css := CssContainer{fileName: Css}
	css.rules = make([]RuleSet, 0, 4)
	return css
}

func (css CssContainer) String() string {
	rulesStrings := make([]string, 0, 1)
	for _, rule := range css.rules {
		rulesStrings = append(rulesStrings, rule.String())
	}
	return strings.Join(rulesStrings, "\n")
}

func (css CssContainer) Rules(rules ...RuleSet) CssContainer {
	css.rules = append(css.rules, rules...)
	return css
}
