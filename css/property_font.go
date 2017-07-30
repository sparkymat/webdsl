package css

import (
	"fmt"

	"github.com/sparkymat/webdsl/css/font"
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

func FontWeight(weight font.Weight) Property {
	property := Property{propertyType: "font-weight"}
	property.values = append(property.values, weight.String())

	return property
}

func FontStyle(style font.Style) Property {
	property := Property{propertyType: "font-style"}
	property.values = append(property.values, style.String())

	return property
}

func Src(url string) Property {
	property := Property{propertyType: "src"}
	property.values = append(property.values, url)

	return property
}
