// Code generated by ent, DO NOT EDIT.

package customer

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/renaldyhidayatt/inventorygoent/pkg/database/postgres/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldName, v))
}

// Alamat applies equality check predicate on the "alamat" field. It's identical to AlamatEQ.
func Alamat(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldAlamat, v))
}

// Telepon applies equality check predicate on the "telepon" field. It's identical to TeleponEQ.
func Telepon(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldTelepon, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldName, v))
}

// AlamatEQ applies the EQ predicate on the "alamat" field.
func AlamatEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldAlamat, v))
}

// AlamatNEQ applies the NEQ predicate on the "alamat" field.
func AlamatNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldAlamat, v))
}

// AlamatIn applies the In predicate on the "alamat" field.
func AlamatIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldAlamat, vs...))
}

// AlamatNotIn applies the NotIn predicate on the "alamat" field.
func AlamatNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldAlamat, vs...))
}

// AlamatGT applies the GT predicate on the "alamat" field.
func AlamatGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldAlamat, v))
}

// AlamatGTE applies the GTE predicate on the "alamat" field.
func AlamatGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldAlamat, v))
}

// AlamatLT applies the LT predicate on the "alamat" field.
func AlamatLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldAlamat, v))
}

// AlamatLTE applies the LTE predicate on the "alamat" field.
func AlamatLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldAlamat, v))
}

// AlamatContains applies the Contains predicate on the "alamat" field.
func AlamatContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldAlamat, v))
}

// AlamatHasPrefix applies the HasPrefix predicate on the "alamat" field.
func AlamatHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldAlamat, v))
}

// AlamatHasSuffix applies the HasSuffix predicate on the "alamat" field.
func AlamatHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldAlamat, v))
}

// AlamatEqualFold applies the EqualFold predicate on the "alamat" field.
func AlamatEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldAlamat, v))
}

// AlamatContainsFold applies the ContainsFold predicate on the "alamat" field.
func AlamatContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldAlamat, v))
}

// TeleponEQ applies the EQ predicate on the "telepon" field.
func TeleponEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldTelepon, v))
}

// TeleponNEQ applies the NEQ predicate on the "telepon" field.
func TeleponNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldTelepon, v))
}

// TeleponIn applies the In predicate on the "telepon" field.
func TeleponIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldTelepon, vs...))
}

// TeleponNotIn applies the NotIn predicate on the "telepon" field.
func TeleponNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldTelepon, vs...))
}

// TeleponGT applies the GT predicate on the "telepon" field.
func TeleponGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldTelepon, v))
}

// TeleponGTE applies the GTE predicate on the "telepon" field.
func TeleponGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldTelepon, v))
}

// TeleponLT applies the LT predicate on the "telepon" field.
func TeleponLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldTelepon, v))
}

// TeleponLTE applies the LTE predicate on the "telepon" field.
func TeleponLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldTelepon, v))
}

// TeleponContains applies the Contains predicate on the "telepon" field.
func TeleponContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldTelepon, v))
}

// TeleponHasPrefix applies the HasPrefix predicate on the "telepon" field.
func TeleponHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldTelepon, v))
}

// TeleponHasSuffix applies the HasSuffix predicate on the "telepon" field.
func TeleponHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldTelepon, v))
}

// TeleponEqualFold applies the EqualFold predicate on the "telepon" field.
func TeleponEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldTelepon, v))
}

// TeleponContainsFold applies the ContainsFold predicate on the "telepon" field.
func TeleponContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldTelepon, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Customer) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Customer) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Customer) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		p(s.Not())
	})
}
