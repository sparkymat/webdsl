package css

import (
	"fmt"

	"github.com/sparkymat/webdsl/css/position"
	"github.com/sparkymat/webdsl/css/size"
)

func Position(position position.Type) Property {
	property := Property{propertyType: "position"}
	property.values = append(property.values, string(position))

	return property
}

func Top(size size.Size) Property {
	property := Property{propertyType: "top"}
	property.values = append(property.values, fmt.Sprintf("%v", size))

	return property
}

func Left(size size.Size) Property {
	property := Property{propertyType: "left"}
	property.values = append(property.values, fmt.Sprintf("%v", size))

	return property
}

func Bottom(size size.Size) Property {
	property := Property{propertyType: "bottom"}
	property.values = append(property.values, fmt.Sprintf("%v", size))

	return property
}

func Right(size size.Size) Property {
	property := Property{propertyType: "right"}
	property.values = append(property.values, fmt.Sprintf("%v", size))

	return property
}

func ZIndex(depth uint32) Property {
	property := Property{propertyType: "z-index"}
	property.values = append(property.values, fmt.Sprintf("%v", depth))

	return property
}
