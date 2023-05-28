package abstraction

type Entity interface {
	FieldMap() ([]string, []interface{})
	TableName() string
}
