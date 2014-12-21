package css

type Class string

func (class Class) Selector() string {
	return "." + string(class)
}

func (class Class) Style(properties ...Property) RuleSet {
	return Rule().For(class).Set(properties...)
}
