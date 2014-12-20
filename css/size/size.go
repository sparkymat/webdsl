package size

import (
	"fmt"

	"github.com/sparkymat/webdsl/types"
)

type Size struct {
	intValue   int64
	floatValue float64
	unit       string
	valueType  types.Type
}

func Px(value int64) Size {
	return Size{intValue: value, unit: "px", valueType: types.Int64}
}

func Percent(value int64) Size {
	return Size{intValue: value, unit: "%", valueType: types.Int64}
}

func Em(value float64) Size {
	return Size{floatValue: value, unit: "em", valueType: types.Float64}
}

func (size Size) String() string {
	switch size.valueType {
	case types.Int8, types.Int16, types.Int32, types.Int64, types.Uint8, types.Uint16, types.Uint32, types.Uint64:
		return fmt.Sprintf("%v%v", size.intValue, size.unit)
	case types.Float32, types.Float64:
		return fmt.Sprintf("%v%v", size.floatValue, size.unit)
	default:
		return fmt.Sprintf("%v%v", size.intValue, size.unit) // FIXME: Better way to handle defaut case?
	}
}
