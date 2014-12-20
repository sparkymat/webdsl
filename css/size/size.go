package size

import "fmt"

type Size struct {
	value int64
	unit  string
}

func Px(value int64) Size {
	return Size{value: value, unit: "px"}
}

func Percent(value int64) Size {
	return Size{value: value, unit: "%"}
}

func (size Size) String() string {
	return fmt.Sprintf("%v%v", size.value, size.unit)
}
