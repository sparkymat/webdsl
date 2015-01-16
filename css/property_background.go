package css

import (
	"fmt"

	"github.com/sparkymat/webdsl/css/background"
	"github.com/sparkymat/webdsl/css/color"
)

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

func BackgroundRepeat(repeatType background.RepeatType) Property {
	property := Property{propertyType: "background-repeat"}
	property.values = append(property.values, string(repeatType))

	return property
}

func BackgroundPosition(position background.PositionType) Property {
	property := Property{propertyType: "background-position"}
	property.values = append(property.values, string(position))

	return property
}
