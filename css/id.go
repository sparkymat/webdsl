package css

type Id string

func (id Id) Selector() string {
	return "#" + string(id)
}
