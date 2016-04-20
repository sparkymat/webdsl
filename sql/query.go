package sql

import (
	"fmt"
	"strings"
)

type Query struct {
	SelectedFields []SelectedField
	FromTable      table
	Joins          []Join
	WhereRelation  Relation
	orders         []Order
	groups         []Column
	limit          *int64
	count          *string
}

func Select(fields ...SelectedField) *Query {
	query := Query{}

	for _, column := range fields {
		query.SelectedFields = append(query.SelectedFields, column)
	}

	return &query
}

func (q *Query) Count() *Query {
	count := "COUNT(*)"
	q.count = &count
	return q
}

func (q *Query) From(t table) *Query {
	q.FromTable = t

	return q
}

func (q *Query) InnerJoin(leftColumn Column, rightColumn Column) *Query {
	join := Join{Type: "INNER JOIN", LeftColumn: leftColumn, RightColumn: rightColumn}
	q.Joins = append(q.Joins, join)

	return q
}

func (q *Query) LeftJoin(leftColumn Column, rightColumn Column) *Query {
	join := Join{Type: "LEFT JOIN", LeftColumn: leftColumn, RightColumn: rightColumn}
	q.Joins = append(q.Joins, join)

	return q
}

func (q *Query) LeftOuterJoin(leftColumn Column, rightColumn Column) *Query {
	join := Join{Type: "LEFT OUTER JOIN", LeftColumn: leftColumn, RightColumn: rightColumn}
	q.Joins = append(q.Joins, join)

	return q
}

func (q *Query) Limit(value int64) *Query {
	q.limit = &value
	return q
}

func (q *Query) Order(orders ...Order) *Query {
	q.orders = orders
	return q
}

func (q *Query) GroupBy(columns ...Column) *Query {
	q.groups = columns
	return q
}

func (q *Query) Where(relation Relation) *Query {
	q.WhereRelation = relation
	return q
}

func (q Query) ToSQLForSubQuery() string {
	selectedFieldStrings := []string{}

	for _, selectedField := range q.SelectedFields {
		selectedFieldStrings = append(selectedFieldStrings, fmt.Sprintf("  %v", selectedField.SelectionString()))
	}

	joinStrings := []string{}

	for _, join := range q.Joins {
		joinStrings = append(joinStrings, fmt.Sprintf("\n%v %v ON %v = %v", join.Type, join.LeftColumn.Table.Name(), join.LeftColumn.String(), join.RightColumn.String()))
	}

	orderString := ""
	eachOrderStrings := []string{}
	for _, eachOrder := range q.orders {
		eachOrderStrings = append(eachOrderStrings, fmt.Sprintf("%v %v", eachOrder.Column.String(), eachOrder.Direction))
	}
	if len(eachOrderStrings) > 0 {
		orderString = fmt.Sprintf("\nORDER BY %v", strings.Join(eachOrderStrings, ", "))
	}

	groupString := ""
	eachGroupStrings := []string{}
	for _, eachGroup := range q.groups {
		eachGroupStrings = append(eachGroupStrings, eachGroup.String())
	}
	if len(eachGroupStrings) > 0 {
		groupString = fmt.Sprintf("\nGROUP BY %v", strings.Join(eachGroupStrings, ", "))
	}

	limitString := ""
	if q.limit != nil {
		limitString = fmt.Sprintf("\nLIMIT %v", *q.limit)
	}

	whereString := ""
	if q.WhereRelation != nil {
		whereString = fmt.Sprintf("\nWHERE %v", q.WhereRelation.QueryFragment())
	}

	sql := `
SELECT
%v
FROM %v %v%v%v%v%v
`

	if q.count != nil {
		return fmt.Sprintf(sql, *q.count, q.FromTable.Name(), strings.Join(joinStrings, ""), whereString, groupString, orderString, limitString)
	}

	return fmt.Sprintf(sql, strings.Join(selectedFieldStrings, ",\n"), q.FromTable.Name(), strings.Join(joinStrings, ""), whereString, groupString, orderString, limitString)
}

func (q Query) ToSQL() string {
	return fmt.Sprintf("%v;", q.ToSQLForSubQuery())
}
