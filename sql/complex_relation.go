package sql

import "fmt"

type ComplexRelation struct {
	LeftQuery  Relation
	Operation  Operation
	RightQuery Relation
}

type Operation string

const Or Operation = "OR"
const And Operation = "AND"

func (q ComplexRelation) Or(rightQuery Relation) ComplexRelation {
	return ComplexRelation{LeftQuery: q, Operation: Or, RightQuery: rightQuery}

}

func (q ComplexRelation) And(rightQuery Relation) ComplexRelation {
	return ComplexRelation{LeftQuery: q, Operation: And, RightQuery: rightQuery}
}

func (q ComplexRelation) QueryFragment() string {
	return fmt.Sprintf("(%v) %v (%v)", q.LeftQuery.QueryFragment(), q.Operation, q.RightQuery.QueryFragment())
}
