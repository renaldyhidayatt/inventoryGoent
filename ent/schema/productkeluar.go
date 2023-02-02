package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ProductKeluar holds the schema definition for the ProductKeluar entity.
type ProductKeluar struct {
	ent.Schema
}

// Fields of the ProductKeluar.
func (ProductKeluar) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("qty"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").UpdateDefault(time.Now),
	}
}

// Edges of the ProductKeluar.
func (ProductKeluar) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("products", Product.Type).Ref("productkeluar"),
		edge.From("category", Category.Type).Ref("productkeluar"),
	}
}
