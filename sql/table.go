package sql

import "fmt"

type table struct {
	name  string
	alias string
}

func Table(name string) table {
	return table{name: name}
}

func (t table) AllColumns() Column {
	return Column{Table: t, Column: "*"}
}

func (t table) Column(columnName string) Column {
	return Column{Table: t, Column: columnName}
}

func (t table) Name() string {
	if t.alias != "" {
		return fmt.Sprintf("%v %v", t.name, t.alias)
	}

	return t.name
}

func (t table) Reference() string {
	if t.alias != "" {
		return t.alias
	}

	return t.name
}

func (t table) Alias(alias string) table {
	t.alias = alias

	return t
}
