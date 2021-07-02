package query

import "strings"

func NewInsert(table string, columns []string) string {
	query := "insert into " + table + " "
	query += "(" + strings.Join(columns, ",") + ") values "
	query += "(" + strings.Join(strings.Split(strings.Repeat("?", len(columns)), ""), ",") + ")"
	return query
}

func NewUpdate(table string, columns []string, key string) string {
	query := "update " + table + " "
	query += "set " + strings.Join(columns, "=?,") + "=? "
	query += "where " + key + " = ?"
	return query
}
