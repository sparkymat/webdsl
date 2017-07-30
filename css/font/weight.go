package font

type Weight string

const WeightNormal = "normal"
const WeightBold = "bold"
const WeightBolder = "bolder"
const WeightLighter = "lighter"
const WeightInitial = "initial"
const WeightInherit = "inherit"

const Weight100 = "100"
const Weight200 = "200"
const Weight300 = "300"
const Weight400 = "400"
const Weight500 = "500"
const Weight600 = "600"
const Weight700 = "700"
const Weight800 = "800"
const Weight900 = "900"

func (w Weight) String() string {
	return string(w)
}
