package font

type Style string

const StyleNormal Style = "normal"
const StyleItalic Style = "italic"
const StyleOblique Style = "oblique"
const StyleInitial Style = "initial"
const StyleInherit Style = "inherit"

func (s Style) String() string {
	return string(s)
}
