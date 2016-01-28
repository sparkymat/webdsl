package sql

type Table string

func (t Table) AllColumns() SelectedColumn {
	return SelectedColumn{Table: t, Column: "*"}
}

func (t Table) Column(columnName string) SelectedColumn {
	return SelectedColumn{Table: t, Column: columnName}
}
