// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/emoss08/trenova/ent/shipmentcommodity"
)

// ShipmentCommodityDelete is the builder for deleting a ShipmentCommodity entity.
type ShipmentCommodityDelete struct {
	config
	hooks    []Hook
	mutation *ShipmentCommodityMutation
}

// Where appends a list predicates to the ShipmentCommodityDelete builder.
func (scd *ShipmentCommodityDelete) Where(ps ...predicate.ShipmentCommodity) *ShipmentCommodityDelete {
	scd.mutation.Where(ps...)
	return scd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (scd *ShipmentCommodityDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, scd.sqlExec, scd.mutation, scd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (scd *ShipmentCommodityDelete) ExecX(ctx context.Context) int {
	n, err := scd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (scd *ShipmentCommodityDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(shipmentcommodity.Table, sqlgraph.NewFieldSpec(shipmentcommodity.FieldID, field.TypeUUID))
	if ps := scd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, scd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	scd.mutation.done = true
	return affected, err
}

// ShipmentCommodityDeleteOne is the builder for deleting a single ShipmentCommodity entity.
type ShipmentCommodityDeleteOne struct {
	scd *ShipmentCommodityDelete
}

// Where appends a list predicates to the ShipmentCommodityDelete builder.
func (scdo *ShipmentCommodityDeleteOne) Where(ps ...predicate.ShipmentCommodity) *ShipmentCommodityDeleteOne {
	scdo.scd.mutation.Where(ps...)
	return scdo
}

// Exec executes the deletion query.
func (scdo *ShipmentCommodityDeleteOne) Exec(ctx context.Context) error {
	n, err := scdo.scd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{shipmentcommodity.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (scdo *ShipmentCommodityDeleteOne) ExecX(ctx context.Context) {
	if err := scdo.Exec(ctx); err != nil {
		panic(err)
	}
}
