package css

type SpecialSelector string

func (selector SpecialSelector) Selector() string {
	return string(selector)
}

const FontFaceSelector SpecialSelector = "@font-face"
