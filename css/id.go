package css

type Id string

func (id Id) Selector() string {
	return "#" + string(id)
}

func (id Id) WithPseudoClass(pseudoClass PseudoClass) SelectorWithPseudoClass {
	return SelectorWithPseudoClass{Element: id, PseudoClass: pseudoClass}
}

func (id Id) Style(properties ...Property) RuleSet {
	return For(id).Set(properties...)
}
