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

func (property Property) WithPropertyType(propertyType string) Property {
	property.propertyType = propertyType
	return property
}

func (property Property) WithValues(values ...string) Property {
	property.values = values
	return property
}

// Properties
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

func MinHeight(distance size.Size) Property {
	property := Property{propertyType: "min-height"}
	property.values = append(property.values, fmt.Sprintf("%v", distance))

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
