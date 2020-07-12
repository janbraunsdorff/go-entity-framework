package item

type Table struct {
	Columns     []Field
	Foreign     []ForeignKey
	Synthetic   []Table
	Synthetics  []Table
	Constraints []string
	Index       []string
}

type Field struct {
	Name      string
	NotNull   bool
	IsUnique  bool
	IsPrimary bool
	DataType  string
	Check     string
	DefValue  string
}

type ForeignKey struct {
	FieldName []string
	ForeignTableName string
	ForeignFieldName []string
}
