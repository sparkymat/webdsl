package css

import (
	"fmt"

	"github.com/sparkymat/webdsl/css/list"
)

func ListStyleType(value list.StyleType) Property {
	property := Property{propertyType: "list-style"}
	property.values = append(property.values, fmt.Sprintf("%v", value))

	return property
}
