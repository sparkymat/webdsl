package sql

import (
	"fmt"
	"strings"
	"time"
)

type SelectedColumn struct {
	Table  table
	Column string
	alias  string
}

func (c SelectedColumn) As(alias string) SelectedColumn {
	c.alias = alias
	return c
}

func (c SelectedColumn) String() string {
	return fmt.Sprintf("%v.%v", c.Table.Reference(), c.Column)
}

func (c SelectedColumn) SelectionString() string {
	if c.alias == "" {
		return c.String()
	}

	return fmt.Sprintf("%v as %v", c.String(), c.alias)
}

func (c SelectedColumn) IsNull() SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v IS NULL", c.String()))
}

func (c SelectedColumn) IsNotNull() SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v IS NULL", c.String()))
}

func (c SelectedColumn) Like(pattern string) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v LIKE '%v'", c.String(), pattern))
}

// EQUALS

func (c SelectedColumn) EqualsPlaceholder() SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v = ?", c.String()))
}

func (c SelectedColumn) EqualsString(value string) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v = '%v'", c.String(), value))
}

func (c SelectedColumn) EqualsInteger(value int64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v = %v", c.String(), value))
}

func (c SelectedColumn) EqualsFloat(value float64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v = %v", c.String(), value))
}

func (c SelectedColumn) EqualsDate(value time.Time) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("DATE(%v) = %v", c.String(), value.Format("2006-01-02")))
}

// NOT EQUALS

func (c SelectedColumn) NotEqualsPlaceholder() SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v != ?", c.String()))
}

func (c SelectedColumn) NotEqualsString(value string) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v != '%v'", c.String(), value))
}

func (c SelectedColumn) NotEqualsInteger(value int64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v != %v", c.String(), value))
}

func (c SelectedColumn) NotEqualsFloat(value float64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v != %v", c.String(), value))
}

// GREATER THAN

func (c SelectedColumn) GreaterThanPlaceholder() SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v > ?", c.String()))
}

func (c SelectedColumn) GreaterThanString(value string) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v > '%v'", c.String(), value))
}

func (c SelectedColumn) GreaterThanInteger(value int64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v > %v", c.String(), value))
}

func (c SelectedColumn) GreaterThanFloat(value float64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v > %v", c.String(), value))
}

func (c SelectedColumn) GreaterThanDate(value time.Time) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("DATE(%v) > '%v'", c.String(), value.Format("2006-01-02")))
}

// GREATER THAN OR EQUAL TO

func (c SelectedColumn) GreaterThanOrEqualsPlaceholder() SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v >= ?", c.String()))
}

func (c SelectedColumn) GreaterThanOrEqualsString(value string) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v >= '%v'", c.String(), value))
}

func (c SelectedColumn) GreaterThanOrEqualsInteger(value int64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v >= %v", c.String(), value))
}

func (c SelectedColumn) GreaterThanOrEqualsFloat(value float64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v >= %v", c.String(), value))
}

func (c SelectedColumn) GreaterThanOrEqualsDate(value time.Time) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("DATE(%v) >= '%v'", c.String(), value.Format("2006-01-02")))
}

// LESS THAN

func (c SelectedColumn) LessThanPlaceholder() SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v < ?", c.String()))
}

func (c SelectedColumn) LessThanString(value string) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v < '%v'", c.String(), value))
}

func (c SelectedColumn) LessThanInteger(value int64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v < %v", c.String(), value))
}

func (c SelectedColumn) LessThanFloat(value float64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v < %v", c.String(), value))
}

func (c SelectedColumn) LessThanDate(value time.Time) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("DATE(%v) < '%v'", c.String(), value.Format("2006-01-02")))
}

// LESS THAN OR EQUAL TO

func (c SelectedColumn) LessThanOrEqualsPlaceholder() SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v <= ?", c.String()))
}

func (c SelectedColumn) LessThanOrEqualsString(value string) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v <= '%v'", c.String(), value))
}

func (c SelectedColumn) LessThanOrEqualsInteger(value int64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v <= %v", c.String(), value))
}

func (c SelectedColumn) LessThanOrEqualsFloat(value float64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v <= %v", c.String(), value))
}

func (c SelectedColumn) LessThanOrEqualsDate(value time.Time) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("DATE(%v) <= '%v'", c.String(), value.Format("2006-01-02")))
}

// IN ARRAY

func (c SelectedColumn) InStringArray(values []string) SimpleRelation {
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

func (c SelectedColumn) InIntegerArray(values []int64) SimpleRelation {
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

func (c SelectedColumn) InFloatArray(values []float64) SimpleRelation {
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
func (c SelectedColumn) SortAscending() Order {
	return Order{Column: c, Direction: Ascending}
}

func (c SelectedColumn) SortDescending() Order {
	return Order{Column: c, Direction: Descending}
}
