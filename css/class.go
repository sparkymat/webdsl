package css

type Class string

func (class Class) Selector() string {
	return "." + string(class)
}

func (class Class) WithPseudoClass(pseudoClass PseudoClass) SelectorWithPseudoClass {
	return SelectorWithPseudoClass{Element: class, PseudoClass: pseudoClass}
}

func (class Class) Style(properties ...Property) RuleSet {
	return For(class).Set(properties...)
}
