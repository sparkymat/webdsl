package css

import (
	"fmt"

	"github.com/sparkymat/webdsl/css/color"
)

type RepeatType string

const Repeat RepeatType = "repeat"
const RepeatX RepeatType = "repeat-x"
const RepeatY RepeatType = "repeat-y"
const NoRepeat RepeatType = "no-repeat"
const RepeatInitial RepeatType = "initial"
const RepeatInherit RepeatType = "inherit"

func BackgroundColor(value color.Color) Property {
	property := Property{propertyType: "background-color"}
	property.values = append(property.values, value.ColorString())

	return property
}

func BackgroundImage(url string) Property {
	property := Property{propertyType: "background-image"}
	property.values = append(property.values, fmt.Sprintf("url(\"%v\")", url))

	return property
}

func BackgroundRepeat(repeatType RepeatType) Property {
	property := Property{propertyType: "background-repeat"}
	property.values = append(property.values, string(repeatType))

	return property
}
