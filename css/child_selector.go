package css

type ChildSelector struct {
	Element Selector
}

func (child ChildSelector) Selector() string {
	return "> " + child.Element.Selector()
}
