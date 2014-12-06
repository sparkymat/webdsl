package css

import (
	"fmt"
	"strings"
)

type CssProperty struct {
	propertyType   string
	propertyValues []string
	isImportant    bool
}

const CssBackgroundRepeat = "repeat"
const CssBackgroundRepeatX = "repeat-x"
const CssBackgroundRepeatY = "repeat-y"
const CssBackgroundNoRepeat = "no-repeat"
const CssBackgroundInitial = "initial"
const CssBackgroundInherit = "inherit"

// Properties
func BackgroundColor(value string) CssProperty {
	property := CssProperty{propertyType: "background-color"}
	property.propertyValues = make([]string, 0, 1)
	property.propertyValues = append(property.propertyValues, value)

	return property
}

func BackgroundImage(url string) CssProperty {
	property := CssProperty{propertyType: "background-image"}
	property.propertyValues = make([]string, 0, 1)
	property.propertyValues = append(property.propertyValues, fmt.Sprintf("url(\"%v\")", url))

	return property
}

func BackgroundRepeat(value string) CssProperty {
	property := CssProperty{propertyType: "background-repeat"}
	property.propertyValues = make([]string, 0, 1)
	property.propertyValues = append(property.propertyValues, value)

	return property
}

func BorderRadius(size uint32) CssProperty {
	property := CssProperty{propertyType: "border-radius"}
	property.propertyValues = make([]string, 0, 1)
	property.propertyValues = append(property.propertyValues, fmt.Sprintf("%vpx", size))

	return property
}

func Color(value string) CssProperty {
	property := CssProperty{propertyType: "color"}
	property.propertyValues = make([]string, 0, 1)
	property.propertyValues = append(property.propertyValues, value)

	return property
}

func FontFamily(values ...string) CssProperty {
	property := CssProperty{propertyType: "font-family"}
	property.propertyValues = make([]string, 0, 1)
	property.propertyValues = append(property.propertyValues, values...)

	return property
}

func FontSizeEm(size uint32) CssProperty {
	property := CssProperty{propertyType: "font-size"}
	property.propertyValues = make([]string, 0, 1)
	property.propertyValues = append(property.propertyValues, fmt.Sprintf("%vem", size))

	return property
}

func Margin1a(top uint32) CssProperty {
	property := CssProperty{propertyType: "margin"}
	property.propertyValues = make([]string, 0, 1)
	property.propertyValues = append(property.propertyValues, fmt.Sprintf("%vpx", top))

	return property
}

func Margin2a(top uint32, right uint32) CssProperty {
	property := CssProperty{propertyType: "margin"}
	property.propertyValues = make([]string, 0, 1)
	property.propertyValues = append(property.propertyValues, fmt.Sprintf("%vpx %vpx", top, right))

	return property
}

func Margin3a(top uint32, right uint32, bottom uint32) CssProperty {
	property := CssProperty{propertyType: "margin"}
	property.propertyValues = make([]string, 0, 1)
	property.propertyValues = append(property.propertyValues, fmt.Sprintf("%vpx %vpx %vpx", top, right, bottom))

	return property
}

func Margin4a(top uint32, right uint32, bottom uint32, left uint32) CssProperty {
	property := CssProperty{propertyType: "margin"}
	property.propertyValues = make([]string, 0, 1)
	property.propertyValues = append(property.propertyValues, fmt.Sprintf("%vpx %vpx %vpx %vpx", top, right, bottom, left))

	return property
}

func MarginLeft(distance uint32) CssProperty {
	property := CssProperty{propertyType: "margin-left"}
	property.propertyValues = make([]string, 0, 1)
	property.propertyValues = append(property.propertyValues, fmt.Sprintf("%vpx", distance))

	return property
}

func MarginRight(distance uint32) CssProperty {
	property := CssProperty{propertyType: "margin-right"}
	property.propertyValues = make([]string, 0, 1)
	property.propertyValues = append(property.propertyValues, fmt.Sprintf("%vpx", distance))

	return property
}

func MarginTop(distance uint32) CssProperty {
	property := CssProperty{propertyType: "margin-top"}
	property.propertyValues = make([]string, 0, 1)
	property.propertyValues = append(property.propertyValues, fmt.Sprintf("%vpx", distance))

	return property
}

func MarginBottom(distance uint32) CssProperty {
	property := CssProperty{propertyType: "margin-bottom"}
	property.propertyValues = make([]string, 0, 1)
	property.propertyValues = append(property.propertyValues, fmt.Sprintf("%vpx", distance))

	return property
}

// Modifiers
func (property CssProperty) Important() CssProperty {
	property.isImportant = true
	return property
}

func (property CssProperty) AddValue(value string) CssProperty {
	property.propertyValues = append(property.propertyValues, value)
	return property
}

// Utils
func (property CssProperty) String() string {
	values := strings.Join(property.propertyValues, ", ")
	if property.isImportant {
		return fmt.Sprintf("%v: %v !important;", property.propertyType, values)
	} else {
		return fmt.Sprintf("%v: %v;", property.propertyType, values)
	}
}
