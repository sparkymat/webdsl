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

func FontWeight(value uint32) Property {
	property := Property{propertyType: "font-weight"}
	property.values = append(property.values, fmt.Sprintf("%v", value))

	return property
}

func FontWeightNormal() Property {
	property := Property{propertyType: "font-weight"}
	property.values = append(property.values, "normal")

	return property
}

func FontWeightBold() Property {
	property := Property{propertyType: "font-weight"}
	property.values = append(property.values, "bold")

	return property
}

func FontWeightLighter() Property {
	property := Property{propertyType: "font-weight"}
	property.values = append(property.values, "lighter")

	return property
}

func FontWeightInherit() Property {
	property := Property{propertyType: "font-weight"}
	property.values = append(property.values, "inherit")

	return property
}

func FontWeightInitial() Property {
	property := Property{propertyType: "font-weight"}
	property.values = append(property.values, "initial")

	return property
}

func FontWeightBolder() Property {
	property := Property{propertyType: "font-weight"}
	property.values = append(property.values, "bolder")

	return property
}
