// Code generated by ent, DO NOT EDIT.

package supplier

import (
	"time"
)

const (
	// Label holds the string label denoting the supplier type in the database.
	Label = "supplier"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAlamat holds the string denoting the alamat field in the database.
	FieldAlamat = "alamat"
	// FieldTelepon holds the string denoting the telepon field in the database.
	FieldTelepon = "telepon"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeProductmasuk holds the string denoting the productmasuk edge name in mutations.
	EdgeProductmasuk = "productmasuk"
	// Table holds the table name of the supplier in the database.
	Table = "suppliers"
	// ProductmasukTable is the table that holds the productmasuk relation/edge. The primary key declared below.
	ProductmasukTable = "product_masuk_supplier"
	// ProductmasukInverseTable is the table name for the ProductMasuk entity.
	// It exists in this package in order to avoid circular dependency with the "productmasuk" package.
	ProductmasukInverseTable = "product_masuks"
)

// Columns holds all SQL columns for supplier fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldAlamat,
	FieldTelepon,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// ProductmasukPrimaryKey and ProductmasukColumn2 are the table columns denoting the
	// primary key for the productmasuk relation (M2M).
	ProductmasukPrimaryKey = []string{"product_masuk_id", "supplier_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)