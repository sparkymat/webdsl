package color

type namedColor string

const Initial namedColor = "initial"
const Inherit namedColor = "inherit"
const Aqua namedColor = "aqua"
const Black namedColor = "black"
const Blue namedColor = "blue"
const Fuchsia namedColor = "fuchsia"
const Gray namedColor = "gray"
const Green namedColor = "green"
const Lime namedColor = "lime"
const Maroon namedColor = "maroon"
const Navy namedColor = "navy"
const Olive namedColor = "olive"
const Orange namedColor = "orange"
const Purple namedColor = "purple"
const Red namedColor = "red"
const Silver namedColor = "silver"
const Teal namedColor = "teal"
const Transparent namedColor = "transparent"
const White namedColor = "white"
const Yellow namedColor = "yellow"

func (color namedColor) ColorString() string {
	return string(color)
}
