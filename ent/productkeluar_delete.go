// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/renaldyhidayatt/inventorygoent/ent/predicate"
	"github.com/renaldyhidayatt/inventorygoent/ent/productkeluar"
)

// ProductKeluarDelete is the builder for deleting a ProductKeluar entity.
type ProductKeluarDelete struct {
	config
	hooks    []Hook
	mutation *ProductKeluarMutation
}

// Where appends a list predicates to the ProductKeluarDelete builder.
func (pkd *ProductKeluarDelete) Where(ps ...predicate.ProductKeluar) *ProductKeluarDelete {
	pkd.mutation.Where(ps...)
	return pkd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pkd *ProductKeluarDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ProductKeluarMutation](ctx, pkd.sqlExec, pkd.mutation, pkd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pkd *ProductKeluarDelete) ExecX(ctx context.Context) int {
	n, err := pkd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pkd *ProductKeluarDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: productkeluar.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: productkeluar.FieldID,
			},
		},
	}
	if ps := pkd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pkd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pkd.mutation.done = true
	return affected, err
}

// ProductKeluarDeleteOne is the builder for deleting a single ProductKeluar entity.
type ProductKeluarDeleteOne struct {
	pkd *ProductKeluarDelete
}

// Where appends a list predicates to the ProductKeluarDelete builder.
func (pkdo *ProductKeluarDeleteOne) Where(ps ...predicate.ProductKeluar) *ProductKeluarDeleteOne {
	pkdo.pkd.mutation.Where(ps...)
	return pkdo
}

// Exec executes the deletion query.
func (pkdo *ProductKeluarDeleteOne) Exec(ctx context.Context) error {
	n, err := pkdo.pkd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{productkeluar.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pkdo *ProductKeluarDeleteOne) ExecX(ctx context.Context) {
	if err := pkdo.Exec(ctx); err != nil {
		panic(err)
	}
}
