package css

type Selector interface {
	Selector() string
}

//TODO bad name
type RawSelector string

func (selector RawSelector) Selector() string {
	return string(selector)
}

func (selector RawSelector) Style(properties ...Property) RuleSet {
	return Rule().For(selector).Set(properties...)
}
