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

type BackgroundPositionType string

const LeftTop BackgroundPositionType = "left top"
const LeftCenter BackgroundPositionType = "left center"
const LeftBottom BackgroundPositionType = "left bottom"
const RightTop BackgroundPositionType = "right top"
const RightCenter BackgroundPositionType = "right center"
const RightBottom BackgroundPositionType = "right bottom"
const CenterTop BackgroundPositionType = "center top"
const CenterCenter BackgroundPositionType = "center center"
const CenterBottom BackgroundPositionType = "center bottom"

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

func BackgroundPosition(position BackgroundPositionType) Property {
	property := Property{propertyType: "background-position"}
	property.values = append(property.values, string(position))

	return property
}
