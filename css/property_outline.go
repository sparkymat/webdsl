package css

import (
	"fmt"
	"image/color"

	"github.com/sparkymat/webdsl/css/outline"
	"github.com/sparkymat/webdsl/css/size"
)

func Outline(color color.Color, style outline.StyleType, width size.Size) Property {
	property := Property{propertyType: "outline"}
	property.values = append(property.values, fmt.Sprintf("%v %v %v", color, style, width))

	return property
}

func OutlineNone() Property {
	property := Property{propertyType: "outline"}
	property.values = append(property.values, "none")

	return property
}

func OutlineInitial() Property {
	property := Property{propertyType: "outline"}
	property.values = append(property.values, "initial")

	return property
}

func OutlineInherit() Property {
	property := Property{propertyType: "outline"}
	property.values = append(property.values, "inherit")

	return property
}
