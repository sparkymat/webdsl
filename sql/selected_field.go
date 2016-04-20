package sql

type SelectedField interface {
	SelectionString() string
}
