package database

import (
	"gorm.io/gorm/schema"
	"strings"
)

// FieldCamelCaseNamer is a naming convention for GORM that use camelCase instead of snake_case.
type FieldCamelCaseNamer struct {
	Default schema.NamingStrategy
}

func (n FieldCamelCaseNamer) TableName(table string) string {
	return n.Default.TableName(table)
}

func (n FieldCamelCaseNamer) SchemaName(table string) string {
	return n.Default.SchemaName(table)
}

func (n FieldCamelCaseNamer) ColumnName(table string, column string) string {
	if strings.ToLower(column) == "id" {
		return "id"
	}
	return strings.ToLower(column[0:1]) + column[1:]
}

func (n FieldCamelCaseNamer) JoinTableName(joinTable string) string {
	return n.Default.JoinTableName(joinTable)
}

func (n FieldCamelCaseNamer) RelationshipFKName(rel schema.Relationship) string {
	return n.Default.RelationshipFKName(rel)
}

func (n FieldCamelCaseNamer) CheckerName(table string, column string) string {
	return n.Default.CheckerName(table, column)
}

func (n FieldCamelCaseNamer) IndexName(table string, column string) string {
	return n.Default.IndexName(table, column)
}
