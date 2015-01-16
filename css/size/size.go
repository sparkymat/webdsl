package size

import "fmt"

type unit string

const px unit = "px"
const em unit = "em"
const percent unit = "%"
const none unit = "none"
const times unit = "times"

var Auto = Size{unit: none, stringValue: "auto"}

type Size struct {
	intValue    int64
	floatValue  float64
	stringValue string
	unit        unit
}

func Times(value float64) Size {
	return Size{floatValue: value, unit: times}
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
	case times:
		return fmt.Sprintf("%v", size.floatValue)
	case none:
		return fmt.Sprintf("%v", size.stringValue)
	default:
		panic("unsupported unit")
	}
}
