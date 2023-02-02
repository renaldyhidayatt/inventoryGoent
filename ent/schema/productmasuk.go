package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ProductMasuk holds the schema definition for the ProductMasuk entity.
type ProductMasuk struct {
	ent.Schema
}

// Fields of the ProductMasuk.
func (ProductMasuk) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("name"),
		field.String("qty"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").UpdateDefault(time.Now),
	}
}

// Edges of the ProductMasuk.
func (ProductMasuk) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).Ref("productmasuk"),
		edge.To("supplier", Supplier.Type),
	}
}
