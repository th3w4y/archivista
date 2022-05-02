// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/testifysec/archivist/ent/dssesignature"
)

// DsseSignatureCreate is the builder for creating a DsseSignature entity.
type DsseSignatureCreate struct {
	config
	mutation *DsseSignatureMutation
	hooks    []Hook
}

// Mutation returns the DsseSignatureMutation object of the builder.
func (dsc *DsseSignatureCreate) Mutation() *DsseSignatureMutation {
	return dsc.mutation
}

// Save creates the DsseSignature in the database.
func (dsc *DsseSignatureCreate) Save(ctx context.Context) (*DsseSignature, error) {
	var (
		err  error
		node *DsseSignature
	)
	if len(dsc.hooks) == 0 {
		if err = dsc.check(); err != nil {
			return nil, err
		}
		node, err = dsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DsseSignatureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dsc.check(); err != nil {
				return nil, err
			}
			dsc.mutation = mutation
			if node, err = dsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dsc.hooks) - 1; i >= 0; i-- {
			if dsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dsc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dsc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dsc *DsseSignatureCreate) SaveX(ctx context.Context) *DsseSignature {
	v, err := dsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dsc *DsseSignatureCreate) Exec(ctx context.Context) error {
	_, err := dsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dsc *DsseSignatureCreate) ExecX(ctx context.Context) {
	if err := dsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dsc *DsseSignatureCreate) check() error {
	return nil
}

func (dsc *DsseSignatureCreate) sqlSave(ctx context.Context) (*DsseSignature, error) {
	_node, _spec := dsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dsc *DsseSignatureCreate) createSpec() (*DsseSignature, *sqlgraph.CreateSpec) {
	var (
		_node = &DsseSignature{config: dsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dssesignature.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dssesignature.FieldID,
			},
		}
	)
	return _node, _spec
}

// DsseSignatureCreateBulk is the builder for creating many DsseSignature entities in bulk.
type DsseSignatureCreateBulk struct {
	config
	builders []*DsseSignatureCreate
}

// Save creates the DsseSignature entities in the database.
func (dscb *DsseSignatureCreateBulk) Save(ctx context.Context) ([]*DsseSignature, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dscb.builders))
	nodes := make([]*DsseSignature, len(dscb.builders))
	mutators := make([]Mutator, len(dscb.builders))
	for i := range dscb.builders {
		func(i int, root context.Context) {
			builder := dscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DsseSignatureMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dscb *DsseSignatureCreateBulk) SaveX(ctx context.Context) []*DsseSignature {
	v, err := dscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dscb *DsseSignatureCreateBulk) Exec(ctx context.Context) error {
	_, err := dscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dscb *DsseSignatureCreateBulk) ExecX(ctx context.Context) {
	if err := dscb.Exec(ctx); err != nil {
		panic(err)
	}
}