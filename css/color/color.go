package color

import "fmt"

type Color interface {
	ColorString() string
}

type ColorRGB struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type ColorRGBA struct {
	Red   uint8
	Green uint8
	Blue  uint8
	Alpha float32
}

func (color ColorRGB) ColorString() string {
	return fmt.Sprintf("rgb(%v, %v, %v)", color.Red, color.Green, color.Blue)
}

func (color ColorRGBA) ColorString() string {
	return fmt.Sprintf("rgba(%v, %v, %v, %.2f)", color.Red, color.Green, color.Blue, color.Alpha)
}
