package css

import (
	"fmt"

	"github.com/sparkymat/webdsl/css/size"
)

type PositionType string

const Absolute PositionType = "absolute"
const Fixed PositionType = "fixed"
const Relative PositionType = "relative"
const PositionStatic PositionType = "static"
const PositionInherit PositionType = "inherit"

func Position(position PositionType) Property {
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
