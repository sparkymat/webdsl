package sql

type Order struct {
	Column    SelectedColumn
	Direction Direction
}

type Direction string

const Ascending Direction = "ASC"
const Descending Direction = "DESC"
