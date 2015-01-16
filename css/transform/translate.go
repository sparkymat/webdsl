package transform

import (
	"fmt"

	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/size"
)

func TranslateX(distance size.Size) css.Property {
	return css.Property{}.WithPropertyType("transform").WithValues(fmt.Sprintf("translateX(%v)", distance))
}
