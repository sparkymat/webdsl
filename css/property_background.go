package css

import (
	"fmt"

	"github.com/sparkymat/webdsl/css/background"
	"github.com/sparkymat/webdsl/css/color"
	"github.com/sparkymat/webdsl/css/size"
)

func BackgroundColor(value color.Color) Property {
	property := Property{propertyType: "background-color"}
	property.values = append(property.values, value.ColorString())

	return property
}

func BackgroundImageNone() Property {
	property := Property{propertyType: "background-image"}
	property.values = append(property.values, "none")

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

func BackgroundPositionX(value size.Size) Property {
	property := Property{propertyType: "background-position-x"}
	property.values = append(property.values, fmt.Sprintf("%v", value))

	return property
}

func BackgroundPositionY(value size.Size) Property {
	property := Property{propertyType: "background-position-y"}
	property.values = append(property.values, fmt.Sprintf("%v", value))

	return property
}

func BackgroundSizeInitial() Property {
	property := Property{propertyType: "background-size"}
	property.values = append(property.values, "initial")

	return property
}

func BackgroundAttachmentScroll() Property {
	property := Property{propertyType: "background-attachment"}
	property.values = append(property.values, "scroll")

	return property
}

func BackgroundOriginInitial() Property {
	property := Property{propertyType: "background-origin"}
	property.values = append(property.values, "initial")

	return property
}

func BackgroundClipInitial() Property {
	property := Property{propertyType: "background-clip"}
	property.values = append(property.values, "initial")

	return property
}
