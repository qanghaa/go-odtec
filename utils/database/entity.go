package database

type Entity interface {
	FieldMap() ([]string, []interface{})
	TableName() string
}
