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
	order          *Order
	limit          *int64
}

func Select(fields ...SelectedColumn) *Query {
	query := Query{}

	for _, column := range fields {
		query.SelectedFields = append(query.SelectedFields, column)
	}

	return &query
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

func (q *Query) Order(column SelectedColumn, direction Direction) *Query {
	value := Order{Column: column, Direction: direction}
	q.order = &value
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
	if q.order != nil {
		orderString = fmt.Sprintf("\nORDER BY %v %v", q.order.Column.String(), q.order.Direction)
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

	return fmt.Sprintf(sql, strings.Join(selectedColumnStrings, ",\n"), q.FromTable.Name(), strings.Join(joinStrings, ""), whereString, orderString, limitString)
}
