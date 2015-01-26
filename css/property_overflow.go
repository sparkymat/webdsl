package css

import "github.com/sparkymat/webdsl/css/overflow"

func Overflow(overflowType overflow.Type) Property {
	property := Property{propertyType: "overflow"}
	property.values = append(property.values, string(overflowType))
	return property
}
