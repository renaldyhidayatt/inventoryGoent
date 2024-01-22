package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Supplier holds the schema definition for the Supplier entity.
type Supplier struct {
	ent.Schema
}

// Fields of the Supplier.
func (Supplier) Fields() []ent.Field {
	return []ent.Field{

		field.String("name"),
		field.String("alamat"),
		field.String("telepon"),

		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").UpdateDefault(time.Now),
	}
}

// Edges of the Supplier.
func (Supplier) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("productmasuk", ProductMasuk.Type).Ref("supplier"),
	}
}
