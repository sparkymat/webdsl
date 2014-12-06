package css

import (
	"fmt"
	"strings"
)

type CssProperty struct {
	propertyType   string
	propertyValues []string
}

// Properties
func Color(value string) CssProperty {
	property := CssProperty{propertyType: "color"}
	property.propertyValues = make([]string, 0, 1)
	property.propertyValues = append(property.propertyValues, value)

	return property
}

func FontSizeEm(size uint32) CssProperty {
	property := CssProperty{propertyType: "font-size"}
	property.propertyValues = make([]string, 0, 1)
	property.propertyValues = append(property.propertyValues, fmt.Sprintf("%vem", size))

	return property
}

// Modifiers
func (property CssProperty) Important() CssProperty {
	property.propertyValues = append(property.propertyValues, "!important")
	return property
}

func (property CssProperty) AddValue(value string) CssProperty {
	property.propertyValues = append(property.propertyValues, value)
	return property
}

// Utils
func (property CssProperty) String() string {
	values := strings.Join(property.propertyValues, ",")
	return fmt.Sprintf("%v: %v;", property.propertyType, values)
}
