package presilo

type SchemaType int

const (
	SCHEMATYPE_OBJECT SchemaType = iota
	SCHEMATYPE_ARRAY
	SCHEMATYPE_STRING
	SCHEMATYPE_NUMBER
	SCHEMATYPE_INTEGER
	SCHEMATYPE_BOOLEAN
	SCHEMATYPE_UNRESOLVED
)
