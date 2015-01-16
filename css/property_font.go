package css

import (
	"fmt"

	"github.com/sparkymat/webdsl/css/size"
)

func FontSize(value size.Size) Property {
	property := Property{propertyType: "font-size"}
	property.values = append(property.values, fmt.Sprintf("%v", value))

	return property
}

func LineHeight(value size.Size) Property {
	property := Property{propertyType: "line-height"}
	property.values = append(property.values, fmt.Sprintf("%v", value))

	return property
}
