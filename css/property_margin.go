package css

import (
	"fmt"
	"strings"

	"github.com/sparkymat/webdsl/css/size"
)

func Margin(sizes ...size.Size) Property {
	var values []string

	for _, value := range sizes {
		values = append(values, value.String())
	}

	property := Property{propertyType: "margin"}
	property.values = append(property.values, strings.Join(values, " "))

	return property
}

func MarginLeft(distance size.Size) Property {
	property := Property{propertyType: "margin-left"}
	property.values = append(property.values, fmt.Sprintf("%v", distance))

	return property
}

func MarginRight(distance size.Size) Property {
	property := Property{propertyType: "margin-right"}
	property.values = append(property.values, fmt.Sprintf("%v", distance))

	return property
}

func MarginTop(distance size.Size) Property {
	property := Property{propertyType: "margin-top"}
	property.values = append(property.values, fmt.Sprintf("%v", distance))

	return property
}

func MarginBottom(distance size.Size) Property {
	property := Property{propertyType: "margin-bottom"}
	property.values = append(property.values, fmt.Sprintf("%v", distance))

	return property
}
