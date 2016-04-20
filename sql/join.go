package sql

type Join struct {
	Type        string
	LeftColumn  Column
	RightColumn Column
}
