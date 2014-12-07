package css

import (
	"fmt"
	"strings"
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
	property.values = make([]string, 0, 1)
	property.values = append(property.values, value)

	return property
}

func BackgroundImage(url string) Property {
	property := Property{propertyType: "background-image"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, fmt.Sprintf("url(\"%v\")", url))

	return property
}

func BackgroundRepeat(value string) Property {
	property := Property{propertyType: "background-repeat"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, value)

	return property
}

func BorderRadius(size uint32) Property {
	property := Property{propertyType: "border-radius"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, fmt.Sprintf("%vpx", size))

	return property
}

func Color(value string) Property {
	property := Property{propertyType: "color"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, value)

	return property
}

func FontFamily(values ...string) Property {
	property := Property{propertyType: "font-family"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, values...)

	return property
}

func FontSizeEm(size uint32) Property {
	property := Property{propertyType: "font-size"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, fmt.Sprintf("%vem", size))

	return property
}

func Padding1a(top uint32) Property {
	property := Property{propertyType: "padding"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, fmt.Sprintf("%vpx", top))

	return property
}

func Padding2a(top uint32, right uint32) Property {
	property := Property{propertyType: "padding"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, fmt.Sprintf("%vpx %vpx", top, right))

	return property
}

func Padding3a(top uint32, right uint32, bottom uint32) Property {
	property := Property{propertyType: "padding"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, fmt.Sprintf("%vpx %vpx %vpx", top, right, bottom))

	return property
}

func Padding4a(top uint32, right uint32, bottom uint32, left uint32) Property {
	property := Property{propertyType: "padding"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, fmt.Sprintf("%vpx %vpx %vpx %vpx", top, right, bottom, left))

	return property
}

func PaddingLeft(distance uint32) Property {
	property := Property{propertyType: "padding-left"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, fmt.Sprintf("%vpx", distance))

	return property
}

func PaddingRight(distance uint32) Property {
	property := Property{propertyType: "padding-right"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, fmt.Sprintf("%vpx", distance))

	return property
}

func PaddingTop(distance uint32) Property {
	property := Property{propertyType: "padding-top"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, fmt.Sprintf("%vpx", distance))

	return property
}

func PaddingBottom(distance uint32) Property {
	property := Property{propertyType: "padding-bottom"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, fmt.Sprintf("%vpx", distance))

	return property
}

func Margin1a(top uint32) Property {
	property := Property{propertyType: "margin"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, fmt.Sprintf("%vpx", top))

	return property
}

func Margin2a(top uint32, right uint32) Property {
	property := Property{propertyType: "margin"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, fmt.Sprintf("%vpx %vpx", top, right))

	return property
}

func Margin3a(top uint32, right uint32, bottom uint32) Property {
	property := Property{propertyType: "margin"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, fmt.Sprintf("%vpx %vpx %vpx", top, right, bottom))

	return property
}

func Margin4a(top uint32, right uint32, bottom uint32, left uint32) Property {
	property := Property{propertyType: "margin"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, fmt.Sprintf("%vpx %vpx %vpx %vpx", top, right, bottom, left))

	return property
}

func MarginLeft(distance uint32) Property {
	property := Property{propertyType: "margin-left"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, fmt.Sprintf("%vpx", distance))

	return property
}

func MarginRight(distance uint32) Property {
	property := Property{propertyType: "margin-right"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, fmt.Sprintf("%vpx", distance))

	return property
}

func MarginTop(distance uint32) Property {
	property := Property{propertyType: "margin-top"}
	property.values = make([]string, 0, 1)
	property.values = append(property.values, fmt.Sprintf("%vpx", distance))

	return property
}

func MarginBottom(distance uint32) Property {
	property := Property{propertyType: "margin-bottom"}
	property.values = make([]string, 0, 1)
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
