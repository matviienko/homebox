// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent/authroles"
	"github.com/hay-kot/homebox/backend/internal/data/ent/authtokens"
	"github.com/hay-kot/homebox/backend/internal/data/ent/predicate"
	"github.com/hay-kot/homebox/backend/internal/data/ent/user"
)

// AuthTokensUpdate is the builder for updating AuthTokens entities.
type AuthTokensUpdate struct {
	config
	hooks    []Hook
	mutation *AuthTokensMutation
}

// Where appends a list predicates to the AuthTokensUpdate builder.
func (atu *AuthTokensUpdate) Where(ps ...predicate.AuthTokens) *AuthTokensUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetUpdatedAt sets the "updated_at" field.
func (atu *AuthTokensUpdate) SetUpdatedAt(t time.Time) *AuthTokensUpdate {
	atu.mutation.SetUpdatedAt(t)
	return atu
}

// SetToken sets the "token" field.
func (atu *AuthTokensUpdate) SetToken(b []byte) *AuthTokensUpdate {
	atu.mutation.SetToken(b)
	return atu
}

// SetExpiresAt sets the "expires_at" field.
func (atu *AuthTokensUpdate) SetExpiresAt(t time.Time) *AuthTokensUpdate {
	atu.mutation.SetExpiresAt(t)
	return atu
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (atu *AuthTokensUpdate) SetNillableExpiresAt(t *time.Time) *AuthTokensUpdate {
	if t != nil {
		atu.SetExpiresAt(*t)
	}
	return atu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (atu *AuthTokensUpdate) SetUserID(id uuid.UUID) *AuthTokensUpdate {
	atu.mutation.SetUserID(id)
	return atu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (atu *AuthTokensUpdate) SetNillableUserID(id *uuid.UUID) *AuthTokensUpdate {
	if id != nil {
		atu = atu.SetUserID(*id)
	}
	return atu
}

// SetUser sets the "user" edge to the User entity.
func (atu *AuthTokensUpdate) SetUser(u *User) *AuthTokensUpdate {
	return atu.SetUserID(u.ID)
}

// SetRolesID sets the "roles" edge to the AuthRoles entity by ID.
func (atu *AuthTokensUpdate) SetRolesID(id int) *AuthTokensUpdate {
	atu.mutation.SetRolesID(id)
	return atu
}

// SetNillableRolesID sets the "roles" edge to the AuthRoles entity by ID if the given value is not nil.
func (atu *AuthTokensUpdate) SetNillableRolesID(id *int) *AuthTokensUpdate {
	if id != nil {
		atu = atu.SetRolesID(*id)
	}
	return atu
}

// SetRoles sets the "roles" edge to the AuthRoles entity.
func (atu *AuthTokensUpdate) SetRoles(a *AuthRoles) *AuthTokensUpdate {
	return atu.SetRolesID(a.ID)
}

// Mutation returns the AuthTokensMutation object of the builder.
func (atu *AuthTokensUpdate) Mutation() *AuthTokensMutation {
	return atu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (atu *AuthTokensUpdate) ClearUser() *AuthTokensUpdate {
	atu.mutation.ClearUser()
	return atu
}

// ClearRoles clears the "roles" edge to the AuthRoles entity.
func (atu *AuthTokensUpdate) ClearRoles() *AuthTokensUpdate {
	atu.mutation.ClearRoles()
	return atu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *AuthTokensUpdate) Save(ctx context.Context) (int, error) {
	atu.defaults()
	return withHooks[int, AuthTokensMutation](ctx, atu.sqlSave, atu.mutation, atu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atu *AuthTokensUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *AuthTokensUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *AuthTokensUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atu *AuthTokensUpdate) defaults() {
	if _, ok := atu.mutation.UpdatedAt(); !ok {
		v := authtokens.UpdateDefaultUpdatedAt()
		atu.mutation.SetUpdatedAt(v)
	}
}

func (atu *AuthTokensUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(authtokens.Table, authtokens.Columns, sqlgraph.NewFieldSpec(authtokens.FieldID, field.TypeUUID))
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atu.mutation.UpdatedAt(); ok {
		_spec.SetField(authtokens.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := atu.mutation.Token(); ok {
		_spec.SetField(authtokens.FieldToken, field.TypeBytes, value)
	}
	if value, ok := atu.mutation.ExpiresAt(); ok {
		_spec.SetField(authtokens.FieldExpiresAt, field.TypeTime, value)
	}
	if atu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authtokens.UserTable,
			Columns: []string{authtokens.UserColumn},
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
	if nodes := atu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authtokens.UserTable,
			Columns: []string{authtokens.UserColumn},
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
	if atu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   authtokens.RolesTable,
			Columns: []string{authtokens.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: authroles.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   authtokens.RolesTable,
			Columns: []string{authtokens.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: authroles.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authtokens.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	atu.mutation.done = true
	return n, nil
}

// AuthTokensUpdateOne is the builder for updating a single AuthTokens entity.
type AuthTokensUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AuthTokensMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (atuo *AuthTokensUpdateOne) SetUpdatedAt(t time.Time) *AuthTokensUpdateOne {
	atuo.mutation.SetUpdatedAt(t)
	return atuo
}

// SetToken sets the "token" field.
func (atuo *AuthTokensUpdateOne) SetToken(b []byte) *AuthTokensUpdateOne {
	atuo.mutation.SetToken(b)
	return atuo
}

// SetExpiresAt sets the "expires_at" field.
func (atuo *AuthTokensUpdateOne) SetExpiresAt(t time.Time) *AuthTokensUpdateOne {
	atuo.mutation.SetExpiresAt(t)
	return atuo
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (atuo *AuthTokensUpdateOne) SetNillableExpiresAt(t *time.Time) *AuthTokensUpdateOne {
	if t != nil {
		atuo.SetExpiresAt(*t)
	}
	return atuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (atuo *AuthTokensUpdateOne) SetUserID(id uuid.UUID) *AuthTokensUpdateOne {
	atuo.mutation.SetUserID(id)
	return atuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (atuo *AuthTokensUpdateOne) SetNillableUserID(id *uuid.UUID) *AuthTokensUpdateOne {
	if id != nil {
		atuo = atuo.SetUserID(*id)
	}
	return atuo
}

// SetUser sets the "user" edge to the User entity.
func (atuo *AuthTokensUpdateOne) SetUser(u *User) *AuthTokensUpdateOne {
	return atuo.SetUserID(u.ID)
}

// SetRolesID sets the "roles" edge to the AuthRoles entity by ID.
func (atuo *AuthTokensUpdateOne) SetRolesID(id int) *AuthTokensUpdateOne {
	atuo.mutation.SetRolesID(id)
	return atuo
}

// SetNillableRolesID sets the "roles" edge to the AuthRoles entity by ID if the given value is not nil.
func (atuo *AuthTokensUpdateOne) SetNillableRolesID(id *int) *AuthTokensUpdateOne {
	if id != nil {
		atuo = atuo.SetRolesID(*id)
	}
	return atuo
}

// SetRoles sets the "roles" edge to the AuthRoles entity.
func (atuo *AuthTokensUpdateOne) SetRoles(a *AuthRoles) *AuthTokensUpdateOne {
	return atuo.SetRolesID(a.ID)
}

// Mutation returns the AuthTokensMutation object of the builder.
func (atuo *AuthTokensUpdateOne) Mutation() *AuthTokensMutation {
	return atuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (atuo *AuthTokensUpdateOne) ClearUser() *AuthTokensUpdateOne {
	atuo.mutation.ClearUser()
	return atuo
}

// ClearRoles clears the "roles" edge to the AuthRoles entity.
func (atuo *AuthTokensUpdateOne) ClearRoles() *AuthTokensUpdateOne {
	atuo.mutation.ClearRoles()
	return atuo
}

// Where appends a list predicates to the AuthTokensUpdate builder.
func (atuo *AuthTokensUpdateOne) Where(ps ...predicate.AuthTokens) *AuthTokensUpdateOne {
	atuo.mutation.Where(ps...)
	return atuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *AuthTokensUpdateOne) Select(field string, fields ...string) *AuthTokensUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated AuthTokens entity.
func (atuo *AuthTokensUpdateOne) Save(ctx context.Context) (*AuthTokens, error) {
	atuo.defaults()
	return withHooks[*AuthTokens, AuthTokensMutation](ctx, atuo.sqlSave, atuo.mutation, atuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *AuthTokensUpdateOne) SaveX(ctx context.Context) *AuthTokens {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *AuthTokensUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *AuthTokensUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atuo *AuthTokensUpdateOne) defaults() {
	if _, ok := atuo.mutation.UpdatedAt(); !ok {
		v := authtokens.UpdateDefaultUpdatedAt()
		atuo.mutation.SetUpdatedAt(v)
	}
}

func (atuo *AuthTokensUpdateOne) sqlSave(ctx context.Context) (_node *AuthTokens, err error) {
	_spec := sqlgraph.NewUpdateSpec(authtokens.Table, authtokens.Columns, sqlgraph.NewFieldSpec(authtokens.FieldID, field.TypeUUID))
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AuthTokens.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authtokens.FieldID)
		for _, f := range fields {
			if !authtokens.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != authtokens.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atuo.mutation.UpdatedAt(); ok {
		_spec.SetField(authtokens.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := atuo.mutation.Token(); ok {
		_spec.SetField(authtokens.FieldToken, field.TypeBytes, value)
	}
	if value, ok := atuo.mutation.ExpiresAt(); ok {
		_spec.SetField(authtokens.FieldExpiresAt, field.TypeTime, value)
	}
	if atuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authtokens.UserTable,
			Columns: []string{authtokens.UserColumn},
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
	if nodes := atuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authtokens.UserTable,
			Columns: []string{authtokens.UserColumn},
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
	if atuo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   authtokens.RolesTable,
			Columns: []string{authtokens.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: authroles.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   authtokens.RolesTable,
			Columns: []string{authtokens.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: authroles.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AuthTokens{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authtokens.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	atuo.mutation.done = true
	return _node, nil
}
