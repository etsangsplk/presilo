package presilo

type TypeSchema interface {
	GetSchemaType() SchemaType
	GetTitle() string
	GetDescription() string
	SetTitle(string)
	GetID() string
	SetID(string)
	GetNullable() bool
	SetNullable(bool)
	HasConstraints() bool
}

/*
  Represents the schema of one field in a json document.
*/
type Schema struct {
	Title       string `json:"title"`
	ID          string `json:"id"`
	Description string `json:"description"`
	Nullable    bool
	typeCode    SchemaType
}

func (this *Schema) GetSchemaType() SchemaType {
	return this.typeCode
}

func (this *Schema) GetTitle() string {
	return this.Title
}

func (this *Schema) GetDescription() string {
	return this.Description
}

func (this *Schema) SetTitle(title string) {
	this.Title = title
}

func (this *Schema) GetID() string {
	return this.ID
}

func (this *Schema) SetID(id string) {
	this.ID = id
}

func (this *Schema) GetNullable() bool {
	return this.Nullable
}

func (this *Schema) SetNullable(nullable bool) {
	this.Nullable = nullable
}
