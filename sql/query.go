package sql

import (
	"fmt"
	"strings"
)

type Query struct {
	SelectedFields []SelectedColumn
	FromTable      table
	Joins          []Join
	WhereRelation  Relation
	orders         []Order
	limit          *int64
	count          *string
}

func Select(fields ...SelectedColumn) *Query {
	query := Query{}

	for _, column := range fields {
		query.SelectedFields = append(query.SelectedFields, column)
	}

	return &query
}

func (q *Query) Count() *Query {
	count := "COUNT(*))"
	q.count = &count
	return q
}

func (q *Query) From(t table) *Query {
	q.FromTable = t

	return q
}

func (q *Query) InnerJoin(leftColumn SelectedColumn, rightColumn SelectedColumn) *Query {
	join := Join{Type: "INNER JOIN", LeftColumn: leftColumn, RightColumn: rightColumn}
	q.Joins = append(q.Joins, join)

	return q
}

func (q *Query) LeftJoin(leftColumn SelectedColumn, rightColumn SelectedColumn) *Query {
	join := Join{Type: "LEFT JOIN", LeftColumn: leftColumn, RightColumn: rightColumn}
	q.Joins = append(q.Joins, join)

	return q
}

func (q *Query) LeftOuterJoin(leftColumn SelectedColumn, rightColumn SelectedColumn) *Query {
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

func (q *Query) Where(relation Relation) *Query {
	q.WhereRelation = relation
	return q
}

func (q Query) ToSQL() string {
	selectedColumnStrings := []string{}

	for _, selectedColumn := range q.SelectedFields {
		selectedColumnStrings = append(selectedColumnStrings, fmt.Sprintf("  %v", selectedColumn.SelectionString()))
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
FROM %v %v%v%v%v;
`

	if q.count != nil {
		return fmt.Sprintf(sql, *q.count, q.FromTable.Name(), strings.Join(joinStrings, ""), whereString, orderString, limitString)
	}

	return fmt.Sprintf(sql, strings.Join(selectedColumnStrings, ",\n"), q.FromTable.Name(), strings.Join(joinStrings, ""), whereString, orderString, limitString)
}
