package css

import (
	"fmt"
	"strings"

	"github.com/sparkymat/webdsl/css/size"
)

func BorderNone() Property {
	property := Property{propertyType: "border"}
	property.values = append(property.values, "none")

	return property
}

func BorderRadius(sizes ...size.Size) Property {
	property := Property{propertyType: "border-radius"}

	var values []string

	for _, size := range sizes {
		values = append(values, size.String())
	}

	property.values = append(property.values, fmt.Sprintf("%v", strings.Join(values, " ")))

	return property
}
