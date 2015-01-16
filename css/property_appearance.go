package css

import "github.com/sparkymat/webdsl/css/appearance"

func Appearance(value appearance.Type) Property {
	property := Property{propertyType: "appearance"}
	property.values = append(property.values, string(value))

	return property
}
