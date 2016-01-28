package sql

type Relation interface {
	QueryFragment() string
}
