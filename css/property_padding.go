package css

import (
	"fmt"
	"strings"

	"github.com/sparkymat/webdsl/css/size"
)

func Padding(sizes ...size.Size) Property {
	var values []string

	for _, value := range sizes {
		values = append(values, value.String())
	}

	property := Property{propertyType: "padding"}
	property.values = append(property.values, strings.Join(values, " "))

	return property
}

func PaddingLeft(distance size.Size) Property {
	property := Property{propertyType: "padding-left"}
	property.values = append(property.values, fmt.Sprintf("%v", distance))

	return property
}

func PaddingRight(distance size.Size) Property {
	property := Property{propertyType: "padding-right"}
	property.values = append(property.values, fmt.Sprintf("%v", distance))

	return property
}

func PaddingTop(distance size.Size) Property {
	property := Property{propertyType: "padding-top"}
	property.values = append(property.values, fmt.Sprintf("%v", distance))

	return property
}

func PaddingBottom(distance size.Size) Property {
	property := Property{propertyType: "padding-bottom"}
	property.values = append(property.values, fmt.Sprintf("%v", distance))

	return property
}
