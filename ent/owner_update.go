// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/m3-app/backend/ent/location"
	"github.com/m3-app/backend/ent/owner"
	"github.com/m3-app/backend/ent/predicate"
	"github.com/m3-app/backend/ent/user"
)

// OwnerUpdate is the builder for updating Owner entities.
type OwnerUpdate struct {
	config
	hooks    []Hook
	mutation *OwnerMutation
}

// Where appends a list predicates to the OwnerUpdate builder.
func (ou *OwnerUpdate) Where(ps ...predicate.Owner) *OwnerUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetFirstName sets the "first_name" field.
func (ou *OwnerUpdate) SetFirstName(s string) *OwnerUpdate {
	ou.mutation.SetFirstName(s)
	return ou
}

// SetMiddleName sets the "middle_name" field.
func (ou *OwnerUpdate) SetMiddleName(s string) *OwnerUpdate {
	ou.mutation.SetMiddleName(s)
	return ou
}

// SetNillableMiddleName sets the "middle_name" field if the given value is not nil.
func (ou *OwnerUpdate) SetNillableMiddleName(s *string) *OwnerUpdate {
	if s != nil {
		ou.SetMiddleName(*s)
	}
	return ou
}

// ClearMiddleName clears the value of the "middle_name" field.
func (ou *OwnerUpdate) ClearMiddleName() *OwnerUpdate {
	ou.mutation.ClearMiddleName()
	return ou
}

// SetLastName sets the "last_name" field.
func (ou *OwnerUpdate) SetLastName(s string) *OwnerUpdate {
	ou.mutation.SetLastName(s)
	return ou
}

// SetFullName sets the "full_name" field.
func (ou *OwnerUpdate) SetFullName(s string) *OwnerUpdate {
	ou.mutation.SetFullName(s)
	return ou
}

// SetIDURL sets the "id_url" field.
func (ou *OwnerUpdate) SetIDURL(s string) *OwnerUpdate {
	ou.mutation.SetIDURL(s)
	return ou
}

// SetVerified sets the "verified" field.
func (ou *OwnerUpdate) SetVerified(b bool) *OwnerUpdate {
	ou.mutation.SetVerified(b)
	return ou
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (ou *OwnerUpdate) SetNillableVerified(b *bool) *OwnerUpdate {
	if b != nil {
		ou.SetVerified(*b)
	}
	return ou
}

// AddLocationIDs adds the "locations" edge to the Location entity by IDs.
func (ou *OwnerUpdate) AddLocationIDs(ids ...uuid.UUID) *OwnerUpdate {
	ou.mutation.AddLocationIDs(ids...)
	return ou
}

// AddLocations adds the "locations" edges to the Location entity.
func (ou *OwnerUpdate) AddLocations(l ...*Location) *OwnerUpdate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ou.AddLocationIDs(ids...)
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (ou *OwnerUpdate) AddUserIDs(ids ...uuid.UUID) *OwnerUpdate {
	ou.mutation.AddUserIDs(ids...)
	return ou
}

// AddUser adds the "user" edges to the User entity.
func (ou *OwnerUpdate) AddUser(u ...*User) *OwnerUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ou.AddUserIDs(ids...)
}

// Mutation returns the OwnerMutation object of the builder.
func (ou *OwnerUpdate) Mutation() *OwnerMutation {
	return ou.mutation
}

// ClearLocations clears all "locations" edges to the Location entity.
func (ou *OwnerUpdate) ClearLocations() *OwnerUpdate {
	ou.mutation.ClearLocations()
	return ou
}

// RemoveLocationIDs removes the "locations" edge to Location entities by IDs.
func (ou *OwnerUpdate) RemoveLocationIDs(ids ...uuid.UUID) *OwnerUpdate {
	ou.mutation.RemoveLocationIDs(ids...)
	return ou
}

// RemoveLocations removes "locations" edges to Location entities.
func (ou *OwnerUpdate) RemoveLocations(l ...*Location) *OwnerUpdate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ou.RemoveLocationIDs(ids...)
}

// ClearUser clears all "user" edges to the User entity.
func (ou *OwnerUpdate) ClearUser() *OwnerUpdate {
	ou.mutation.ClearUser()
	return ou
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (ou *OwnerUpdate) RemoveUserIDs(ids ...uuid.UUID) *OwnerUpdate {
	ou.mutation.RemoveUserIDs(ids...)
	return ou
}

// RemoveUser removes "user" edges to User entities.
func (ou *OwnerUpdate) RemoveUser(u ...*User) *OwnerUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ou.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OwnerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, OwnerMutation](ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OwnerUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OwnerUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OwnerUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OwnerUpdate) check() error {
	if v, ok := ou.mutation.FirstName(); ok {
		if err := owner.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "Owner.first_name": %w`, err)}
		}
	}
	if v, ok := ou.mutation.LastName(); ok {
		if err := owner.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "Owner.last_name": %w`, err)}
		}
	}
	if v, ok := ou.mutation.FullName(); ok {
		if err := owner.FullNameValidator(v); err != nil {
			return &ValidationError{Name: "full_name", err: fmt.Errorf(`ent: validator failed for field "Owner.full_name": %w`, err)}
		}
	}
	if v, ok := ou.mutation.IDURL(); ok {
		if err := owner.IDURLValidator(v); err != nil {
			return &ValidationError{Name: "id_url", err: fmt.Errorf(`ent: validator failed for field "Owner.id_url": %w`, err)}
		}
	}
	return nil
}

func (ou *OwnerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(owner.Table, owner.Columns, sqlgraph.NewFieldSpec(owner.FieldID, field.TypeUUID))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.FirstName(); ok {
		_spec.SetField(owner.FieldFirstName, field.TypeString, value)
	}
	if value, ok := ou.mutation.MiddleName(); ok {
		_spec.SetField(owner.FieldMiddleName, field.TypeString, value)
	}
	if ou.mutation.MiddleNameCleared() {
		_spec.ClearField(owner.FieldMiddleName, field.TypeString)
	}
	if value, ok := ou.mutation.LastName(); ok {
		_spec.SetField(owner.FieldLastName, field.TypeString, value)
	}
	if value, ok := ou.mutation.FullName(); ok {
		_spec.SetField(owner.FieldFullName, field.TypeString, value)
	}
	if value, ok := ou.mutation.IDURL(); ok {
		_spec.SetField(owner.FieldIDURL, field.TypeString, value)
	}
	if value, ok := ou.mutation.Verified(); ok {
		_spec.SetField(owner.FieldVerified, field.TypeBool, value)
	}
	if ou.mutation.LocationsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedLocationsIDs(); len(nodes) > 0 && !ou.mutation.LocationsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.LocationsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedUserIDs(); len(nodes) > 0 && !ou.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{owner.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OwnerUpdateOne is the builder for updating a single Owner entity.
type OwnerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OwnerMutation
}

// SetFirstName sets the "first_name" field.
func (ouo *OwnerUpdateOne) SetFirstName(s string) *OwnerUpdateOne {
	ouo.mutation.SetFirstName(s)
	return ouo
}

// SetMiddleName sets the "middle_name" field.
func (ouo *OwnerUpdateOne) SetMiddleName(s string) *OwnerUpdateOne {
	ouo.mutation.SetMiddleName(s)
	return ouo
}

// SetNillableMiddleName sets the "middle_name" field if the given value is not nil.
func (ouo *OwnerUpdateOne) SetNillableMiddleName(s *string) *OwnerUpdateOne {
	if s != nil {
		ouo.SetMiddleName(*s)
	}
	return ouo
}

// ClearMiddleName clears the value of the "middle_name" field.
func (ouo *OwnerUpdateOne) ClearMiddleName() *OwnerUpdateOne {
	ouo.mutation.ClearMiddleName()
	return ouo
}

// SetLastName sets the "last_name" field.
func (ouo *OwnerUpdateOne) SetLastName(s string) *OwnerUpdateOne {
	ouo.mutation.SetLastName(s)
	return ouo
}

// SetFullName sets the "full_name" field.
func (ouo *OwnerUpdateOne) SetFullName(s string) *OwnerUpdateOne {
	ouo.mutation.SetFullName(s)
	return ouo
}

// SetIDURL sets the "id_url" field.
func (ouo *OwnerUpdateOne) SetIDURL(s string) *OwnerUpdateOne {
	ouo.mutation.SetIDURL(s)
	return ouo
}

// SetVerified sets the "verified" field.
func (ouo *OwnerUpdateOne) SetVerified(b bool) *OwnerUpdateOne {
	ouo.mutation.SetVerified(b)
	return ouo
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (ouo *OwnerUpdateOne) SetNillableVerified(b *bool) *OwnerUpdateOne {
	if b != nil {
		ouo.SetVerified(*b)
	}
	return ouo
}

// AddLocationIDs adds the "locations" edge to the Location entity by IDs.
func (ouo *OwnerUpdateOne) AddLocationIDs(ids ...uuid.UUID) *OwnerUpdateOne {
	ouo.mutation.AddLocationIDs(ids...)
	return ouo
}

// AddLocations adds the "locations" edges to the Location entity.
func (ouo *OwnerUpdateOne) AddLocations(l ...*Location) *OwnerUpdateOne {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ouo.AddLocationIDs(ids...)
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (ouo *OwnerUpdateOne) AddUserIDs(ids ...uuid.UUID) *OwnerUpdateOne {
	ouo.mutation.AddUserIDs(ids...)
	return ouo
}

// AddUser adds the "user" edges to the User entity.
func (ouo *OwnerUpdateOne) AddUser(u ...*User) *OwnerUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ouo.AddUserIDs(ids...)
}

// Mutation returns the OwnerMutation object of the builder.
func (ouo *OwnerUpdateOne) Mutation() *OwnerMutation {
	return ouo.mutation
}

// ClearLocations clears all "locations" edges to the Location entity.
func (ouo *OwnerUpdateOne) ClearLocations() *OwnerUpdateOne {
	ouo.mutation.ClearLocations()
	return ouo
}

// RemoveLocationIDs removes the "locations" edge to Location entities by IDs.
func (ouo *OwnerUpdateOne) RemoveLocationIDs(ids ...uuid.UUID) *OwnerUpdateOne {
	ouo.mutation.RemoveLocationIDs(ids...)
	return ouo
}

// RemoveLocations removes "locations" edges to Location entities.
func (ouo *OwnerUpdateOne) RemoveLocations(l ...*Location) *OwnerUpdateOne {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ouo.RemoveLocationIDs(ids...)
}

// ClearUser clears all "user" edges to the User entity.
func (ouo *OwnerUpdateOne) ClearUser() *OwnerUpdateOne {
	ouo.mutation.ClearUser()
	return ouo
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (ouo *OwnerUpdateOne) RemoveUserIDs(ids ...uuid.UUID) *OwnerUpdateOne {
	ouo.mutation.RemoveUserIDs(ids...)
	return ouo
}

// RemoveUser removes "user" edges to User entities.
func (ouo *OwnerUpdateOne) RemoveUser(u ...*User) *OwnerUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ouo.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the OwnerUpdate builder.
func (ouo *OwnerUpdateOne) Where(ps ...predicate.Owner) *OwnerUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OwnerUpdateOne) Select(field string, fields ...string) *OwnerUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Owner entity.
func (ouo *OwnerUpdateOne) Save(ctx context.Context) (*Owner, error) {
	return withHooks[*Owner, OwnerMutation](ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OwnerUpdateOne) SaveX(ctx context.Context) *Owner {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OwnerUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OwnerUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OwnerUpdateOne) check() error {
	if v, ok := ouo.mutation.FirstName(); ok {
		if err := owner.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "Owner.first_name": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.LastName(); ok {
		if err := owner.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "Owner.last_name": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.FullName(); ok {
		if err := owner.FullNameValidator(v); err != nil {
			return &ValidationError{Name: "full_name", err: fmt.Errorf(`ent: validator failed for field "Owner.full_name": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.IDURL(); ok {
		if err := owner.IDURLValidator(v); err != nil {
			return &ValidationError{Name: "id_url", err: fmt.Errorf(`ent: validator failed for field "Owner.id_url": %w`, err)}
		}
	}
	return nil
}

func (ouo *OwnerUpdateOne) sqlSave(ctx context.Context) (_node *Owner, err error) {
	if err := ouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(owner.Table, owner.Columns, sqlgraph.NewFieldSpec(owner.FieldID, field.TypeUUID))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Owner.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, owner.FieldID)
		for _, f := range fields {
			if !owner.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != owner.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.FirstName(); ok {
		_spec.SetField(owner.FieldFirstName, field.TypeString, value)
	}
	if value, ok := ouo.mutation.MiddleName(); ok {
		_spec.SetField(owner.FieldMiddleName, field.TypeString, value)
	}
	if ouo.mutation.MiddleNameCleared() {
		_spec.ClearField(owner.FieldMiddleName, field.TypeString)
	}
	if value, ok := ouo.mutation.LastName(); ok {
		_spec.SetField(owner.FieldLastName, field.TypeString, value)
	}
	if value, ok := ouo.mutation.FullName(); ok {
		_spec.SetField(owner.FieldFullName, field.TypeString, value)
	}
	if value, ok := ouo.mutation.IDURL(); ok {
		_spec.SetField(owner.FieldIDURL, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Verified(); ok {
		_spec.SetField(owner.FieldVerified, field.TypeBool, value)
	}
	if ouo.mutation.LocationsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedLocationsIDs(); len(nodes) > 0 && !ouo.mutation.LocationsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.LocationsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedUserIDs(); len(nodes) > 0 && !ouo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Owner{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{owner.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
