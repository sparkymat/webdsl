package css

import (
	"fmt"
	"strings"

	"github.com/sparkymat/webdsl/css/size"
)

type Property struct {
	propertyType string
	values       []string
	isImportant  bool
}

const CssBackgroundRepeat = "repeat"
const CssBackgroundRepeatX = "repeat-x"
const CssBackgroundRepeatY = "repeat-y"
const CssBackgroundNoRepeat = "no-repeat"
const CssBackgroundInitial = "initial"
const CssBackgroundInherit = "inherit"

// Properties
func BackgroundColor(value string) Property {
	property := Property{propertyType: "background-color"}
	property.values = append(property.values, value)

	return property
}

func BackgroundImage(url string) Property {
	property := Property{propertyType: "background-image"}
	property.values = append(property.values, fmt.Sprintf("url(\"%v\")", url))

	return property
}

func BackgroundRepeat(value string) Property {
	property := Property{propertyType: "background-repeat"}
	property.values = append(property.values, value)

	return property
}

func BorderRadius(size size.Size) Property {
	property := Property{propertyType: "border-radius"}
	property.values = append(property.values, fmt.Sprintf("%v", size))

	return property
}

func FontFamily(values ...string) Property {
	property := Property{propertyType: "font-family"}
	property.values = append(property.values, values...)

	return property
}

func FontSizeEm(size uint32) Property {
	property := Property{propertyType: "font-size"}
	property.values = append(property.values, fmt.Sprintf("%vem", size))

	return property
}

func Padding1i(top int32) Property {
	property := Property{propertyType: "padding"}
	property.values = append(property.values, fmt.Sprintf("%vpx", top))

	return property
}

func Padding2i(top int32, right int32) Property {
	property := Property{propertyType: "padding"}
	property.values = append(property.values, fmt.Sprintf("%vpx %vpx", top, right))

	return property
}

func Padding3i(top int32, right int32, bottom int32) Property {
	property := Property{propertyType: "padding"}
	property.values = append(property.values, fmt.Sprintf("%vpx %vpx %vpx", top, right, bottom))

	return property
}

func Padding4i(top int32, right int32, bottom int32, left int32) Property {
	property := Property{propertyType: "padding"}
	property.values = append(property.values, fmt.Sprintf("%vpx %vpx %vpx %vpx", top, right, bottom, left))

	return property
}

func PaddingLeft(distance int32) Property {
	property := Property{propertyType: "padding-left"}
	property.values = append(property.values, fmt.Sprintf("%vpx", distance))

	return property
}

func PaddingRight(distance int32) Property {
	property := Property{propertyType: "padding-right"}
	property.values = append(property.values, fmt.Sprintf("%vpx", distance))

	return property
}

func PaddingTop(distance int32) Property {
	property := Property{propertyType: "padding-top"}
	property.values = append(property.values, fmt.Sprintf("%vpx", distance))

	return property
}

func PaddingBottom(distance int32) Property {
	property := Property{propertyType: "padding-bottom"}
	property.values = append(property.values, fmt.Sprintf("%vpx", distance))

	return property
}

func Width(distance size.Size) Property {
	property := Property{propertyType: "width"}
	property.values = append(property.values, fmt.Sprintf("%v", distance))

	return property
}

func Height(distance size.Size) Property {
	property := Property{propertyType: "height"}
	property.values = append(property.values, fmt.Sprintf("%v", distance))

	return property
}

func Margin1i(top int32) Property {
	property := Property{propertyType: "margin"}
	property.values = append(property.values, fmt.Sprintf("%vpx", top))

	return property
}

func Margin2i(top int32, right int32) Property {
	property := Property{propertyType: "margin"}
	property.values = append(property.values, fmt.Sprintf("%vpx %vpx", top, right))

	return property
}

func Margin3i(top int32, right int32, bottom int32) Property {
	property := Property{propertyType: "margin"}
	property.values = append(property.values, fmt.Sprintf("%vpx %vpx %vpx", top, right, bottom))

	return property
}

func Margin4i(top int32, right int32, bottom int32, left int32) Property {
	property := Property{propertyType: "margin"}
	property.values = append(property.values, fmt.Sprintf("%vpx %vpx %vpx %vpx", top, right, bottom, left))

	return property
}

func MarginLeft(distance int32) Property {
	property := Property{propertyType: "margin-left"}
	property.values = append(property.values, fmt.Sprintf("%vpx", distance))

	return property
}

func MarginRight(distance int32) Property {
	property := Property{propertyType: "margin-right"}
	property.values = append(property.values, fmt.Sprintf("%vpx", distance))

	return property
}

func MarginTop(distance int32) Property {
	property := Property{propertyType: "margin-top"}
	property.values = append(property.values, fmt.Sprintf("%vpx", distance))

	return property
}

func MarginBottom(distance int32) Property {
	property := Property{propertyType: "margin-bottom"}
	property.values = append(property.values, fmt.Sprintf("%vpx", distance))

	return property
}

// Modifiers
func (property Property) Important() Property {
	property.isImportant = true
	return property
}

func (property Property) AddValue(value string) Property {
	property.values = append(property.values, value)
	return property
}

// Utils
func (property Property) String() string {
	values := strings.Join(property.values, ", ")
	if property.isImportant {
		return fmt.Sprintf("%v: %v !important;", property.propertyType, values)
	} else {
		return fmt.Sprintf("%v: %v;", property.propertyType, values)
	}
}
