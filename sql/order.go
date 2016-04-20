package sql

type Order struct {
	Column    Column
	Direction Direction
}

type Direction string

const Ascending Direction = "ASC"
const Descending Direction = "DESC"
