package sql

import "fmt"

type SelectedColumn struct {
	Table  Table
	Column string
}

func (c SelectedColumn) String() string {
	return fmt.Sprintf("%v.%v", c.Table, c.Column)
}

func (c SelectedColumn) IsNull() SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v.%v IS NULL", c.Table, c.Column))
}

func (c SelectedColumn) IsNotNull() SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v.%v IS NULL", c.Table, c.Column))
}

func (c SelectedColumn) Like(pattern string) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v.%v LIKE '%v'", c.Table, c.Column, pattern))
}

func (c SelectedColumn) EqualsString(value string) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v.%v = '%v'", c.Table, c.Column, value))
}

func (c SelectedColumn) EqualsInteger(value int64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v.%v = %v", c.Table, c.Column, value))
}

func (c SelectedColumn) EqualsFloat(value float64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v.%v = %v", c.Table, c.Column, value))
}

func (c SelectedColumn) NotEqualsString(value string) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v.%v != '%v'", c.Table, c.Column, value))
}

func (c SelectedColumn) NotEqualsInteger(value int64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v.%v != %v", c.Table, c.Column, value))
}

func (c SelectedColumn) NotEqualsFloat(value float64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v.%v != %v", c.Table, c.Column, value))
}

func (c SelectedColumn) GreaterThanInteger(value int64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v.%v > %v", c.Table, c.Column, value))
}

func (c SelectedColumn) GreaterThanOrEqualToInteger(value int64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v.%v >= %v", c.Table, c.Column, value))
}

func (c SelectedColumn) LessThanInteger(value int64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v.%v < %v", c.Table, c.Column, value))
}

func (c SelectedColumn) LessThanOrEqualToInteger(value int64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v.%v <= %v", c.Table, c.Column, value))
}

func (c SelectedColumn) GreaterThanFloat(value float64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v.%v > %v", c.Table, c.Column, value))
}

func (c SelectedColumn) GreaterThanOrEqualToFloat(value float64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v.%v >= %v", c.Table, c.Column, value))
}

func (c SelectedColumn) LessThanFloat(value float64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v.%v < %v", c.Table, c.Column, value))
}

func (c SelectedColumn) LessThanOrEqualToFloat(value float64) SimpleRelation {
	return SimpleRelation(fmt.Sprintf("%v.%v <= %v", c.Table, c.Column, value))
}
