package css

import (
	"fmt"

	"github.com/sparkymat/webdsl/css/color"
	"github.com/sparkymat/webdsl/css/size"
	"github.com/sparkymat/webdsl/css/textdecoration"
)

// Letter spacing
func LetterSpacingNormal() Property {
	property := Property{propertyType: "letter-spacing"}
	property.values = append(property.values, "normal")

	return property
}

func LetterSpacingInitial() Property {
	property := Property{propertyType: "letter-spacing"}
	property.values = append(property.values, "initial")

	return property
}

func LetterSpacingInherit() Property {
	property := Property{propertyType: "letter-spacing"}
	property.values = append(property.values, "inherit")

	return property
}

func LetterSpacing(distance size.Size) Property {
	property := Property{propertyType: "letter-spacing"}
	property.values = append(property.values, distance.String())

	return property
}

// Direction functions
func DirectionLtr() Property {
	property := Property{propertyType: "direction"}
	property.values = append(property.values, "ltr")

	return property
}

func DirectionRtl() Property {
	property := Property{propertyType: "direction"}
	property.values = append(property.values, "rtl")

	return property
}

func DirectionInitial() Property {
	property := Property{propertyType: "direction"}
	property.values = append(property.values, "initial")

	return property
}

func DirectionInherit() Property {
	property := Property{propertyType: "direction"}
	property.values = append(property.values, "inherit")

	return property
}

// Color functions
func Color(value color.Color) Property {
	property := Property{propertyType: "color"}
	property.values = append(property.values, value.ColorString())

	return property
}

// Text decoration
func TextDecoration(value textdecoration.DecorationType) Property {
	property := Property{propertyType: "text-decoration"}
	property.values = append(property.values, string(value))

	return property
}

// Text shadow
func TextShadow(horizontal size.Size, vertical size.Size, blur size.Size, shadowColor color.Color) Property {
	property := Property{propertyType: "text-shadow"}
	property.values = append(property.values, fmt.Sprintf("%v %v %v %v", horizontal, vertical, blur, shadowColor))

	return property
}
