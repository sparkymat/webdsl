package css

import (
	"fmt"

	"github.com/sparkymat/webdsl/css/color"
	"github.com/sparkymat/webdsl/css/size"
)

func BoxShadow(horizontal size.Size, vertical size.Size, blur size.Size, shadowColor color.Color) Property {
	property := Property{propertyType: "box-shadow"}
	property.values = append(property.values, fmt.Sprintf("%v %v %v %v", horizontal, vertical, blur, shadowColor.ColorString()))

	return property
}

func BoxShadowNone() Property {
	property := Property{propertyType: "box-shadow"}
	property.values = append(property.values, "none")

	return property
}
