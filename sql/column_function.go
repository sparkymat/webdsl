package sql

import "fmt"

type ColumnFunction struct {
	Column   Column
	Function string
}

func (f ColumnFunction) SelectionString() string {
	return fmt.Sprintf("%v(%v)", f.Function, f.Column.String())
}
