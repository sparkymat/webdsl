package css

import (
	"fmt"

	"github.com/sparkymat/webdsl/css/display"
	"github.com/sparkymat/webdsl/css/flex"
)

func Display(value display.Type) Property {
	property := Property{propertyType: "display"}
	property.values = append(property.values, string(value))

	return property
}

func FlexDirection(value flex.Direction) Property {
	property := Property{propertyType: "flex-direction"}
	property.values = append(property.values, string(value))

	return property
}

func FlexWrap(value flex.WrapType) Property {
	property := Property{propertyType: "flex-wrap"}
	property.values = append(property.values, string(value))

	return property
}

func FlexGrow(value int) Property {
	property := Property{propertyType: "flex-grow"}
	property.values = append(property.values, fmt.Sprintf("%v", value))

	return property
}

//TODO this is used for flex box only. move this somewhere
type itemAlign string

const Center itemAlign = "center"
const Stretch itemAlign = "stretch"

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
