package sql

type Join struct {
	Type        string
	LeftColumn  SelectedColumn
	RightColumn SelectedColumn
}
