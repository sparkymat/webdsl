package css

type Selector interface {
	Selector() string
}

func selectorsToStrings(selectorsSelectors []Selector) []string {
	var selectors []string
	for i := range selectorsSelectors {
		selectors = append(selectors, selectorsSelectors[i].Selector())
	}
	return selectors
}
