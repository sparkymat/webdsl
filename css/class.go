package css

type Class string

func (class Class) Selector() string {
	return "." + string(class)
}
