// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/m3-app/backend/ent/location"
	"github.com/m3-app/backend/ent/owner"
	"github.com/m3-app/backend/ent/user"
)

// OwnerCreate is the builder for creating a Owner entity.
type OwnerCreate struct {
	config
	mutation *OwnerMutation
	hooks    []Hook
}

// SetFirstName sets the "first_name" field.
func (oc *OwnerCreate) SetFirstName(s string) *OwnerCreate {
	oc.mutation.SetFirstName(s)
	return oc
}

// SetMiddleName sets the "middle_name" field.
func (oc *OwnerCreate) SetMiddleName(s string) *OwnerCreate {
	oc.mutation.SetMiddleName(s)
	return oc
}

// SetNillableMiddleName sets the "middle_name" field if the given value is not nil.
func (oc *OwnerCreate) SetNillableMiddleName(s *string) *OwnerCreate {
	if s != nil {
		oc.SetMiddleName(*s)
	}
	return oc
}

// SetLastName sets the "last_name" field.
func (oc *OwnerCreate) SetLastName(s string) *OwnerCreate {
	oc.mutation.SetLastName(s)
	return oc
}

// SetFullName sets the "full_name" field.
func (oc *OwnerCreate) SetFullName(s string) *OwnerCreate {
	oc.mutation.SetFullName(s)
	return oc
}

// SetIDURL sets the "id_url" field.
func (oc *OwnerCreate) SetIDURL(s string) *OwnerCreate {
	oc.mutation.SetIDURL(s)
	return oc
}

// SetVerified sets the "verified" field.
func (oc *OwnerCreate) SetVerified(b bool) *OwnerCreate {
	oc.mutation.SetVerified(b)
	return oc
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (oc *OwnerCreate) SetNillableVerified(b *bool) *OwnerCreate {
	if b != nil {
		oc.SetVerified(*b)
	}
	return oc
}

// SetID sets the "id" field.
func (oc *OwnerCreate) SetID(u uuid.UUID) *OwnerCreate {
	oc.mutation.SetID(u)
	return oc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (oc *OwnerCreate) SetNillableID(u *uuid.UUID) *OwnerCreate {
	if u != nil {
		oc.SetID(*u)
	}
	return oc
}

// AddLocationIDs adds the "locations" edge to the Location entity by IDs.
func (oc *OwnerCreate) AddLocationIDs(ids ...uuid.UUID) *OwnerCreate {
	oc.mutation.AddLocationIDs(ids...)
	return oc
}

// AddLocations adds the "locations" edges to the Location entity.
func (oc *OwnerCreate) AddLocations(l ...*Location) *OwnerCreate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return oc.AddLocationIDs(ids...)
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (oc *OwnerCreate) AddUserIDs(ids ...uuid.UUID) *OwnerCreate {
	oc.mutation.AddUserIDs(ids...)
	return oc
}

// AddUser adds the "user" edges to the User entity.
func (oc *OwnerCreate) AddUser(u ...*User) *OwnerCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return oc.AddUserIDs(ids...)
}

// Mutation returns the OwnerMutation object of the builder.
func (oc *OwnerCreate) Mutation() *OwnerMutation {
	return oc.mutation
}

// Save creates the Owner in the database.
func (oc *OwnerCreate) Save(ctx context.Context) (*Owner, error) {
	oc.defaults()
	return withHooks[*Owner, OwnerMutation](ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OwnerCreate) SaveX(ctx context.Context) *Owner {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OwnerCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OwnerCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OwnerCreate) defaults() {
	if _, ok := oc.mutation.Verified(); !ok {
		v := owner.DefaultVerified
		oc.mutation.SetVerified(v)
	}
	if _, ok := oc.mutation.ID(); !ok {
		v := owner.DefaultID()
		oc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OwnerCreate) check() error {
	if _, ok := oc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "first_name", err: errors.New(`ent: missing required field "Owner.first_name"`)}
	}
	if v, ok := oc.mutation.FirstName(); ok {
		if err := owner.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "Owner.first_name": %w`, err)}
		}
	}
	if _, ok := oc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "Owner.last_name"`)}
	}
	if v, ok := oc.mutation.LastName(); ok {
		if err := owner.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "Owner.last_name": %w`, err)}
		}
	}
	if _, ok := oc.mutation.FullName(); !ok {
		return &ValidationError{Name: "full_name", err: errors.New(`ent: missing required field "Owner.full_name"`)}
	}
	if v, ok := oc.mutation.FullName(); ok {
		if err := owner.FullNameValidator(v); err != nil {
			return &ValidationError{Name: "full_name", err: fmt.Errorf(`ent: validator failed for field "Owner.full_name": %w`, err)}
		}
	}
	if _, ok := oc.mutation.IDURL(); !ok {
		return &ValidationError{Name: "id_url", err: errors.New(`ent: missing required field "Owner.id_url"`)}
	}
	if v, ok := oc.mutation.IDURL(); ok {
		if err := owner.IDURLValidator(v); err != nil {
			return &ValidationError{Name: "id_url", err: fmt.Errorf(`ent: validator failed for field "Owner.id_url": %w`, err)}
		}
	}
	if _, ok := oc.mutation.Verified(); !ok {
		return &ValidationError{Name: "verified", err: errors.New(`ent: missing required field "Owner.verified"`)}
	}
	return nil
}

func (oc *OwnerCreate) sqlSave(ctx context.Context) (*Owner, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OwnerCreate) createSpec() (*Owner, *sqlgraph.CreateSpec) {
	var (
		_node = &Owner{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(owner.Table, sqlgraph.NewFieldSpec(owner.FieldID, field.TypeUUID))
	)
	if id, ok := oc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := oc.mutation.FirstName(); ok {
		_spec.SetField(owner.FieldFirstName, field.TypeString, value)
		_node.FirstName = value
	}
	if value, ok := oc.mutation.MiddleName(); ok {
		_spec.SetField(owner.FieldMiddleName, field.TypeString, value)
		_node.MiddleName = value
	}
	if value, ok := oc.mutation.LastName(); ok {
		_spec.SetField(owner.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := oc.mutation.FullName(); ok {
		_spec.SetField(owner.FieldFullName, field.TypeString, value)
		_node.FullName = value
	}
	if value, ok := oc.mutation.IDURL(); ok {
		_spec.SetField(owner.FieldIDURL, field.TypeString, value)
		_node.IDURL = value
	}
	if value, ok := oc.mutation.Verified(); ok {
		_spec.SetField(owner.FieldVerified, field.TypeBool, value)
		_node.Verified = value
	}
	if nodes := oc.mutation.LocationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   owner.LocationsTable,
			Columns: []string{owner.LocationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: location.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   owner.UserTable,
			Columns: []string{owner.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OwnerCreateBulk is the builder for creating many Owner entities in bulk.
type OwnerCreateBulk struct {
	config
	builders []*OwnerCreate
}

// Save creates the Owner entities in the database.
func (ocb *OwnerCreateBulk) Save(ctx context.Context) ([]*Owner, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Owner, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OwnerMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OwnerCreateBulk) SaveX(ctx context.Context) []*Owner {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OwnerCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OwnerCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
