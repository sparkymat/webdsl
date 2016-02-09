package sql

import "fmt"

type SimpleRelation string

func (r SimpleRelation) QueryFragment() string {
	return fmt.Sprintf("%v", r)
}

func (r SimpleRelation) Or(rightQuery Relation) ComplexRelation {
	if complexRightQuery, isComplex := rightQuery.(ComplexRelation); isComplex {
		return complexRightQuery.Or(r)
	}

	return ComplexRelation{Operation: Or, Relations: []Relation{r, rightQuery}}
}

func (r SimpleRelation) And(rightQuery Relation) ComplexRelation {
	if complexRightQuery, isComplex := rightQuery.(ComplexRelation); isComplex {
		return complexRightQuery.And(r)
	}

	return ComplexRelation{Operation: And, Relations: []Relation{r, rightQuery}}
}
