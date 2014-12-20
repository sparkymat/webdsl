package css

import (
	"github.com/sparkymat/webdsl/css/color"
	"github.com/sparkymat/webdsl/css/size"
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

func ColorInherit() Property {
	property := Property{propertyType: "color"}
	property.values = append(property.values, "inherit")

	return property
}

func ColorInitial() Property {
	property := Property{propertyType: "color"}
	property.values = append(property.values, "initial")

	return property
}

// Standard colors
func ColorAqua() Property {
	property := Property{propertyType: "color"}
	property.values = append(property.values, "aqua")

	return property
}

func ColorBlack() Property {
	property := Property{propertyType: "color"}
	property.values = append(property.values, "black")

	return property
}

func ColorBlue() Property {
	property := Property{propertyType: "color"}
	property.values = append(property.values, "blue")

	return property
}

func ColorFuchsia() Property {
	property := Property{propertyType: "color"}
	property.values = append(property.values, "fuchsia")

	return property
}

func ColorGray() Property {
	property := Property{propertyType: "color"}
	property.values = append(property.values, "gray")

	return property
}

func ColorGreen() Property {
	property := Property{propertyType: "color"}
	property.values = append(property.values, "green")

	return property
}

func ColorLime() Property {
	property := Property{propertyType: "color"}
	property.values = append(property.values, "lime")

	return property
}

func ColorMaroon() Property {
	property := Property{propertyType: "color"}
	property.values = append(property.values, "maroon")

	return property
}

func ColorNavy() Property {
	property := Property{propertyType: "color"}
	property.values = append(property.values, "navy")

	return property
}

func ColorOlive() Property {
	property := Property{propertyType: "color"}
	property.values = append(property.values, "olive")

	return property
}

func ColorOrange() Property {
	property := Property{propertyType: "color"}
	property.values = append(property.values, "orange")

	return property
}

func ColorPurple() Property {
	property := Property{propertyType: "color"}
	property.values = append(property.values, "purple")

	return property
}

func ColorRed() Property {
	property := Property{propertyType: "color"}
	property.values = append(property.values, "red")

	return property
}

func ColorSilver() Property {
	property := Property{propertyType: "color"}
	property.values = append(property.values, "silver")

	return property
}

func ColorTeal() Property {
	property := Property{propertyType: "color"}
	property.values = append(property.values, "teal")

	return property
}

func ColorWhite() Property {
	property := Property{propertyType: "color"}
	property.values = append(property.values, "white")

	return property
}

func ColorYellow() Property {
	property := Property{propertyType: "color"}
	property.values = append(property.values, "yellow")

	return property
}
