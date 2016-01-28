package sql

import "fmt"

type SimpleRelation string

func (r SimpleRelation) QueryFragment() string {
	return fmt.Sprintf("%v", r)
}

func (r SimpleRelation) Or(rightQuery Relation) ComplexRelation {
	return ComplexRelation{LeftQuery: r, Operation: Or, RightQuery: rightQuery}
}

func (r SimpleRelation) And(rightQuery Relation) ComplexRelation {
	return ComplexRelation{LeftQuery: r, Operation: And, RightQuery: rightQuery}
}
