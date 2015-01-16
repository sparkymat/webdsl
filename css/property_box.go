package css

import (
	"fmt"

	"github.com/sparkymat/webdsl/css/box"
)

func BoxSizing(sizingType box.BoxSizingType) Property {
	property := Property{propertyType: "box-sizing"}
	property.values = append(property.values, fmt.Sprintf("%v", sizingType))

	return property
}
