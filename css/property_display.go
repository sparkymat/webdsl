package css

import (
	"fmt"

	"github.com/sparkymat/webdsl/css/display"
)

func Display(value display.Type) Property {
	property := Property{propertyType: "display"}
	property.values = append(property.values, string(value))

	return property
}

type flexDirection string

const Column flexDirection = "column"
const Row flexDirection = "row"

func FlexDirection(value flexDirection) Property {
	property := Property{propertyType: "flex-direction"}
	property.values = append(property.values, string(value))

	return property
}

func FlexGrow(value int) Property {
	property := Property{propertyType: "flex-grow"}
	property.values = append(property.values, fmt.Sprintf("%v", value))

	return property
}

type itemAlign string

const Center itemAlign = "center"

func AlignItems(value itemAlign) Property {
	property := Property{propertyType: "align-items"}
	property.values = append(property.values, string(value))

	return property
}

type justifyContent string

const FlexEnd justifyContent = "flex-end"

func JustifyContent(value justifyContent) Property {
	property := Property{propertyType: "justify-content"}
	property.values = append(property.values, string(value))

	return property
}
