package size

import "fmt"

type unit string

const px unit = "px"
const em unit = "em"
const percent unit = "%"

type Size struct {
	intValue   int64
	floatValue float64
	unit       unit
}

func Px(value int64) Size {
	return Size{intValue: value, unit: px}
}

func Percent(value int64) Size {
	return Size{intValue: value, unit: percent}
}

func Em(value float64) Size {
	return Size{floatValue: value, unit: em}
}

func (size Size) String() string {
	switch size.unit {
	case px, percent:
		return fmt.Sprintf("%v%v", size.intValue, size.unit)
	case em:
		return fmt.Sprintf("%v%v", size.floatValue, size.unit)
	default:
		panic("unsupported unit")
	}
}
