package css

type wordWrapType string

const BreakWord wordWrapType = "break-word"

func WordWrap(wordWrap wordWrapType) Property {
	property := Property{propertyType: "word-wrap"}
	property.values = append(property.values, string(wordWrap))
	return property
}
