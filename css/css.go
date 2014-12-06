package css

import "strings"

type Css struct {
	fileName string
	rules    []CssRule
}

func NewCss(name string) Css {
	css := Css{fileName: name}
	css.rules = make([]CssRule, 0, 4)
	return css
}

func (css Css) String() string {
	rulesStrings := make([]string, 0, 1)
	for _, rule := range css.rules {
		rulesStrings = append(rulesStrings, rule.String())
	}
	return strings.Join(rulesStrings, "\n")
}

func (css Css) Rules(rules ...CssRule) Css {
	css.rules = append(css.rules, rules...)
	return css
}
