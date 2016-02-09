package sql

import (
	"fmt"
	"strings"
)

type ComplexRelation struct {
	Relations []Relation
	Operation Operation
}

type Operation string

const Or Operation = "OR"
const And Operation = "AND"

func (q ComplexRelation) Or(relation Relation) ComplexRelation {
	if q.Operation == Or {
		q.Relations = append(q.Relations, relation)
		return q
	}

	return ComplexRelation{Operation: Or, Relations: []Relation{q, relation}}
}

func (q ComplexRelation) And(relation Relation) ComplexRelation {
	if q.Operation == And {
		q.Relations = append(q.Relations, relation)
		return q
	}

	return ComplexRelation{Operation: And, Relations: []Relation{q, relation}}
}

func (q ComplexRelation) QueryFragment() string {
	var formattedRelations []string

	for _, relation := range q.Relations {
		formattedRelations = append(formattedRelations, fmt.Sprintf("(%v)", relation.QueryFragment()))
	}

	return strings.Join(formattedRelations, fmt.Sprintf(" %v ", q.Operation))
}
