package css

type displayType string

//TODO Feel like display should be own package
const Block displayType = "block"
const Flex displayType = "flex"

func Display(value displayType) Property {
	property := Property{propertyType: "display"}
	property.values = append(property.values, string(value))

	return property
}
