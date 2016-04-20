package sql

import (
	"fmt"
	"strings"
	"time"
)

type Column struct {
	Table  table
	Column string
	alias  string
}

func (c Column) As(alias string) Column {
	c.alias = alias
	return c
}

func (c Column) String() string {
	return fmt.Sprintf("%v.%v", c.Table.Reference(), c.Column)
}

func (c Column) SelectionString() string {
	if c.alias == "" {
		return c.String()
	}

	return fmt.Sprintf("%v as %v", c.String(), c.alias)
}

func (c Column) IsNull() SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v IS NULL", c.String()))
}

func (c Column) IsNotNull() SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v IS NULL", c.String()))
}

func (c Column) Like(pattern string) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v LIKE '%v'", c.String(), pattern))
}

// EQUALS

func (c Column) EqualsPlaceholder() SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v = ?", c.String()))
}

func (c Column) EqualsString(value string) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v = '%v'", c.String(), value))
}

func (c Column) EqualsInteger(value int64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v = %v", c.String(), value))
}

func (c Column) EqualsFloat(value float64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v = %v", c.String(), value))
}

func (c Column) EqualsDate(value time.Time) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("DATE(%v) = %v", c.String(), value.Format("2006-01-02")))
}

// NOT EQUALS

func (c Column) NotEqualsPlaceholder() SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v != ?", c.String()))
}

func (c Column) NotEqualsString(value string) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v != '%v'", c.String(), value))
}

func (c Column) NotEqualsInteger(value int64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v != %v", c.String(), value))
}

func (c Column) NotEqualsFloat(value float64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v != %v", c.String(), value))
}

// GREATER THAN

func (c Column) GreaterThanPlaceholder() SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v > ?", c.String()))
}

func (c Column) GreaterThanString(value string) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v > '%v'", c.String(), value))
}

func (c Column) GreaterThanInteger(value int64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v > %v", c.String(), value))
}

func (c Column) GreaterThanFloat(value float64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v > %v", c.String(), value))
}

func (c Column) GreaterThanDate(value time.Time) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("DATE(%v) > '%v'", c.String(), value.Format("2006-01-02")))
}

// GREATER THAN OR EQUAL TO

func (c Column) GreaterThanOrEqualsPlaceholder() SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v >= ?", c.String()))
}

func (c Column) GreaterThanOrEqualsString(value string) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v >= '%v'", c.String(), value))
}

func (c Column) GreaterThanOrEqualsInteger(value int64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v >= %v", c.String(), value))
}

func (c Column) GreaterThanOrEqualsFloat(value float64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v >= %v", c.String(), value))
}

func (c Column) GreaterThanOrEqualsDate(value time.Time) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("DATE(%v) >= '%v'", c.String(), value.Format("2006-01-02")))
}

// LESS THAN

func (c Column) LessThanPlaceholder() SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v < ?", c.String()))
}

func (c Column) LessThanString(value string) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v < '%v'", c.String(), value))
}

func (c Column) LessThanInteger(value int64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v < %v", c.String(), value))
}

func (c Column) LessThanFloat(value float64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v < %v", c.String(), value))
}

func (c Column) LessThanDate(value time.Time) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("DATE(%v) < '%v'", c.String(), value.Format("2006-01-02")))
}

// LESS THAN OR EQUAL TO

func (c Column) LessThanOrEqualsPlaceholder() SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v <= ?", c.String()))
}

func (c Column) LessThanOrEqualsString(value string) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v <= '%v'", c.String(), value))
}

func (c Column) LessThanOrEqualsInteger(value int64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v <= %v", c.String(), value))
}

func (c Column) LessThanOrEqualsFloat(value float64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v <= %v", c.String(), value))
}

func (c Column) LessThanOrEqualsDate(value time.Time) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("DATE(%v) <= '%v'", c.String(), value.Format("2006-01-02")))
}

// IN ARRAY

func (c Column) InStringArray(values []string) SimpleRelation {
	if len(values) == 0 {
		return c.IsNull()
	}

	if len(values) == 1 {
		return c.EqualsString(values[0])
	}

	var transformedStrings []string
	for _, eachValue := range values {
		transformedStrings = append(transformedStrings, fmt.Sprintf("'%v'", eachValue))
	}
	return SimpleRelation(fmt.Sprintf("%v IN (%v)", c.String(), strings.Join(transformedStrings, ", ")))
}

func (c Column) InIntegerArray(values []int64) SimpleRelation {
	if len(values) == 0 {
		return c.IsNull()
	}

	if len(values) == 1 {
		return c.EqualsInteger(values[0])
	}

	var transformedStrings []string
	for _, eachValue := range values {
		transformedStrings = append(transformedStrings, fmt.Sprintf("%v", eachValue))
	}
	return SimpleRelation(fmt.Sprintf("%v IN (%v)", c.String(), strings.Join(transformedStrings, ", ")))
}

func (c Column) InFloatArray(values []float64) SimpleRelation {
	if len(values) == 0 {
		return c.IsNull()
	}

	if len(values) == 1 {
		return c.EqualsFloat(values[0])
	}

	var transformedStrings []string
	for _, eachValue := range values {
		transformedStrings = append(transformedStrings, fmt.Sprintf("%v", eachValue))
	}
	return SimpleRelation(fmt.Sprintf("%v IN (%v)", c.String(), strings.Join(transformedStrings, ", ")))
}

// Sorts
func (c Column) SortAscending() Order {
	return Order{Column: c, Direction: Ascending}
}

func (c Column) SortDescending() Order {
	return Order{Column: c, Direction: Descending}
}

// Subquery
func (c Column) In(subQuery *Query) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v IN (%v)", c.String(), subQuery.ToSQLForSubQuery()))
}

func (c Column) NotIn(subQuery *Query) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v NOT IN (%v)", c.String(), subQuery.ToSQLForSubQuery()))
}

func (c Column) GroupConcat() ColumnFunction {
	return ColumnFunction{Column: c, Function: "group_concat"}
}
