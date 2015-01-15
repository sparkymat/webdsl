package css

import (
	"github.com/sparkymat/webdsl/css/clear"
	"github.com/sparkymat/webdsl/css/float"
)

func Float(floatType float.Type) Property {
	property := Property{propertyType: "float"}
	property.values = append(property.values, string(floatType))

	return property
}

func Clear(clearType clear.Type) Property {
	property := Property{propertyType: "clear"}
	property.values = append(property.values, string(clearType))

	return property
}
