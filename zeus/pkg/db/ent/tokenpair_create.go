// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Geapefurit/kline-back/zeus/pkg/db/ent/tokenpair"
)

// TokenPairCreate is the builder for creating a TokenPair entity.
type TokenPairCreate struct {
	config
	mutation *TokenPairMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (tpc *TokenPairCreate) SetCreatedAt(u uint32) *TokenPairCreate {
	tpc.mutation.SetCreatedAt(u)
	return tpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tpc *TokenPairCreate) SetNillableCreatedAt(u *uint32) *TokenPairCreate {
	if u != nil {
		tpc.SetCreatedAt(*u)
	}
	return tpc
}

// SetUpdatedAt sets the "updated_at" field.
func (tpc *TokenPairCreate) SetUpdatedAt(u uint32) *TokenPairCreate {
	tpc.mutation.SetUpdatedAt(u)
	return tpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tpc *TokenPairCreate) SetNillableUpdatedAt(u *uint32) *TokenPairCreate {
	if u != nil {
		tpc.SetUpdatedAt(*u)
	}
	return tpc
}

// SetDeletedAt sets the "deleted_at" field.
func (tpc *TokenPairCreate) SetDeletedAt(u uint32) *TokenPairCreate {
	tpc.mutation.SetDeletedAt(u)
	return tpc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tpc *TokenPairCreate) SetNillableDeletedAt(u *uint32) *TokenPairCreate {
	if u != nil {
		tpc.SetDeletedAt(*u)
	}
	return tpc
}

// SetTokenOneID sets the "token_one_id" field.
func (tpc *TokenPairCreate) SetTokenOneID(u uint32) *TokenPairCreate {
	tpc.mutation.SetTokenOneID(u)
	return tpc
}

// SetTokenTwoID sets the "token_two_id" field.
func (tpc *TokenPairCreate) SetTokenTwoID(u uint32) *TokenPairCreate {
	tpc.mutation.SetTokenTwoID(u)
	return tpc
}

// SetRemark sets the "remark" field.
func (tpc *TokenPairCreate) SetRemark(s string) *TokenPairCreate {
	tpc.mutation.SetRemark(s)
	return tpc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (tpc *TokenPairCreate) SetNillableRemark(s *string) *TokenPairCreate {
	if s != nil {
		tpc.SetRemark(*s)
	}
	return tpc
}

// SetID sets the "id" field.
func (tpc *TokenPairCreate) SetID(u uint32) *TokenPairCreate {
	tpc.mutation.SetID(u)
	return tpc
}

// Mutation returns the TokenPairMutation object of the builder.
func (tpc *TokenPairCreate) Mutation() *TokenPairMutation {
	return tpc.mutation
}

// Save creates the TokenPair in the database.
func (tpc *TokenPairCreate) Save(ctx context.Context) (*TokenPair, error) {
	var (
		err  error
		node *TokenPair
	)
	if err := tpc.defaults(); err != nil {
		return nil, err
	}
	if len(tpc.hooks) == 0 {
		if err = tpc.check(); err != nil {
			return nil, err
		}
		node, err = tpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TokenPairMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tpc.check(); err != nil {
				return nil, err
			}
			tpc.mutation = mutation
			if node, err = tpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tpc.hooks) - 1; i >= 0; i-- {
			if tpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tpc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tpc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TokenPair)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TokenPairMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tpc *TokenPairCreate) SaveX(ctx context.Context) *TokenPair {
	v, err := tpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tpc *TokenPairCreate) Exec(ctx context.Context) error {
	_, err := tpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpc *TokenPairCreate) ExecX(ctx context.Context) {
	if err := tpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tpc *TokenPairCreate) defaults() error {
	if _, ok := tpc.mutation.CreatedAt(); !ok {
		if tokenpair.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized tokenpair.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := tokenpair.DefaultCreatedAt()
		tpc.mutation.SetCreatedAt(v)
	}
	if _, ok := tpc.mutation.UpdatedAt(); !ok {
		if tokenpair.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized tokenpair.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := tokenpair.DefaultUpdatedAt()
		tpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tpc.mutation.DeletedAt(); !ok {
		if tokenpair.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized tokenpair.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := tokenpair.DefaultDeletedAt()
		tpc.mutation.SetDeletedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tpc *TokenPairCreate) check() error {
	if _, ok := tpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TokenPair.created_at"`)}
	}
	if _, ok := tpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "TokenPair.updated_at"`)}
	}
	if _, ok := tpc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "TokenPair.deleted_at"`)}
	}
	if _, ok := tpc.mutation.TokenOneID(); !ok {
		return &ValidationError{Name: "token_one_id", err: errors.New(`ent: missing required field "TokenPair.token_one_id"`)}
	}
	if _, ok := tpc.mutation.TokenTwoID(); !ok {
		return &ValidationError{Name: "token_two_id", err: errors.New(`ent: missing required field "TokenPair.token_two_id"`)}
	}
	return nil
}

func (tpc *TokenPairCreate) sqlSave(ctx context.Context) (*TokenPair, error) {
	_node, _spec := tpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	return _node, nil
}

func (tpc *TokenPairCreate) createSpec() (*TokenPair, *sqlgraph.CreateSpec) {
	var (
		_node = &TokenPair{config: tpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tokenpair.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: tokenpair.FieldID,
			},
		}
	)
	_spec.OnConflict = tpc.conflict
	if id, ok := tpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tpc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tpc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := tpc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := tpc.mutation.TokenOneID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldTokenOneID,
		})
		_node.TokenOneID = value
	}
	if value, ok := tpc.mutation.TokenTwoID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tokenpair.FieldTokenTwoID,
		})
		_node.TokenTwoID = value
	}
	if value, ok := tpc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tokenpair.FieldRemark,
		})
		_node.Remark = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TokenPair.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TokenPairUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (tpc *TokenPairCreate) OnConflict(opts ...sql.ConflictOption) *TokenPairUpsertOne {
	tpc.conflict = opts
	return &TokenPairUpsertOne{
		create: tpc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TokenPair.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tpc *TokenPairCreate) OnConflictColumns(columns ...string) *TokenPairUpsertOne {
	tpc.conflict = append(tpc.conflict, sql.ConflictColumns(columns...))
	return &TokenPairUpsertOne{
		create: tpc,
	}
}

type (
	// TokenPairUpsertOne is the builder for "upsert"-ing
	//  one TokenPair node.
	TokenPairUpsertOne struct {
		create *TokenPairCreate
	}

	// TokenPairUpsert is the "OnConflict" setter.
	TokenPairUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *TokenPairUpsert) SetCreatedAt(v uint32) *TokenPairUpsert {
	u.Set(tokenpair.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TokenPairUpsert) UpdateCreatedAt() *TokenPairUpsert {
	u.SetExcluded(tokenpair.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *TokenPairUpsert) AddCreatedAt(v uint32) *TokenPairUpsert {
	u.Add(tokenpair.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TokenPairUpsert) SetUpdatedAt(v uint32) *TokenPairUpsert {
	u.Set(tokenpair.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TokenPairUpsert) UpdateUpdatedAt() *TokenPairUpsert {
	u.SetExcluded(tokenpair.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *TokenPairUpsert) AddUpdatedAt(v uint32) *TokenPairUpsert {
	u.Add(tokenpair.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TokenPairUpsert) SetDeletedAt(v uint32) *TokenPairUpsert {
	u.Set(tokenpair.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TokenPairUpsert) UpdateDeletedAt() *TokenPairUpsert {
	u.SetExcluded(tokenpair.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *TokenPairUpsert) AddDeletedAt(v uint32) *TokenPairUpsert {
	u.Add(tokenpair.FieldDeletedAt, v)
	return u
}

// SetTokenOneID sets the "token_one_id" field.
func (u *TokenPairUpsert) SetTokenOneID(v uint32) *TokenPairUpsert {
	u.Set(tokenpair.FieldTokenOneID, v)
	return u
}

// UpdateTokenOneID sets the "token_one_id" field to the value that was provided on create.
func (u *TokenPairUpsert) UpdateTokenOneID() *TokenPairUpsert {
	u.SetExcluded(tokenpair.FieldTokenOneID)
	return u
}

// AddTokenOneID adds v to the "token_one_id" field.
func (u *TokenPairUpsert) AddTokenOneID(v uint32) *TokenPairUpsert {
	u.Add(tokenpair.FieldTokenOneID, v)
	return u
}

// SetTokenTwoID sets the "token_two_id" field.
func (u *TokenPairUpsert) SetTokenTwoID(v uint32) *TokenPairUpsert {
	u.Set(tokenpair.FieldTokenTwoID, v)
	return u
}

// UpdateTokenTwoID sets the "token_two_id" field to the value that was provided on create.
func (u *TokenPairUpsert) UpdateTokenTwoID() *TokenPairUpsert {
	u.SetExcluded(tokenpair.FieldTokenTwoID)
	return u
}

// AddTokenTwoID adds v to the "token_two_id" field.
func (u *TokenPairUpsert) AddTokenTwoID(v uint32) *TokenPairUpsert {
	u.Add(tokenpair.FieldTokenTwoID, v)
	return u
}

// SetRemark sets the "remark" field.
func (u *TokenPairUpsert) SetRemark(v string) *TokenPairUpsert {
	u.Set(tokenpair.FieldRemark, v)
	return u
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *TokenPairUpsert) UpdateRemark() *TokenPairUpsert {
	u.SetExcluded(tokenpair.FieldRemark)
	return u
}

// ClearRemark clears the value of the "remark" field.
func (u *TokenPairUpsert) ClearRemark() *TokenPairUpsert {
	u.SetNull(tokenpair.FieldRemark)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TokenPair.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(tokenpair.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TokenPairUpsertOne) UpdateNewValues() *TokenPairUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(tokenpair.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TokenPair.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TokenPairUpsertOne) Ignore() *TokenPairUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TokenPairUpsertOne) DoNothing() *TokenPairUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TokenPairCreate.OnConflict
// documentation for more info.
func (u *TokenPairUpsertOne) Update(set func(*TokenPairUpsert)) *TokenPairUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TokenPairUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TokenPairUpsertOne) SetCreatedAt(v uint32) *TokenPairUpsertOne {
	return u.Update(func(s *TokenPairUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *TokenPairUpsertOne) AddCreatedAt(v uint32) *TokenPairUpsertOne {
	return u.Update(func(s *TokenPairUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TokenPairUpsertOne) UpdateCreatedAt() *TokenPairUpsertOne {
	return u.Update(func(s *TokenPairUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TokenPairUpsertOne) SetUpdatedAt(v uint32) *TokenPairUpsertOne {
	return u.Update(func(s *TokenPairUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *TokenPairUpsertOne) AddUpdatedAt(v uint32) *TokenPairUpsertOne {
	return u.Update(func(s *TokenPairUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TokenPairUpsertOne) UpdateUpdatedAt() *TokenPairUpsertOne {
	return u.Update(func(s *TokenPairUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TokenPairUpsertOne) SetDeletedAt(v uint32) *TokenPairUpsertOne {
	return u.Update(func(s *TokenPairUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *TokenPairUpsertOne) AddDeletedAt(v uint32) *TokenPairUpsertOne {
	return u.Update(func(s *TokenPairUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TokenPairUpsertOne) UpdateDeletedAt() *TokenPairUpsertOne {
	return u.Update(func(s *TokenPairUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetTokenOneID sets the "token_one_id" field.
func (u *TokenPairUpsertOne) SetTokenOneID(v uint32) *TokenPairUpsertOne {
	return u.Update(func(s *TokenPairUpsert) {
		s.SetTokenOneID(v)
	})
}

// AddTokenOneID adds v to the "token_one_id" field.
func (u *TokenPairUpsertOne) AddTokenOneID(v uint32) *TokenPairUpsertOne {
	return u.Update(func(s *TokenPairUpsert) {
		s.AddTokenOneID(v)
	})
}

// UpdateTokenOneID sets the "token_one_id" field to the value that was provided on create.
func (u *TokenPairUpsertOne) UpdateTokenOneID() *TokenPairUpsertOne {
	return u.Update(func(s *TokenPairUpsert) {
		s.UpdateTokenOneID()
	})
}

// SetTokenTwoID sets the "token_two_id" field.
func (u *TokenPairUpsertOne) SetTokenTwoID(v uint32) *TokenPairUpsertOne {
	return u.Update(func(s *TokenPairUpsert) {
		s.SetTokenTwoID(v)
	})
}

// AddTokenTwoID adds v to the "token_two_id" field.
func (u *TokenPairUpsertOne) AddTokenTwoID(v uint32) *TokenPairUpsertOne {
	return u.Update(func(s *TokenPairUpsert) {
		s.AddTokenTwoID(v)
	})
}

// UpdateTokenTwoID sets the "token_two_id" field to the value that was provided on create.
func (u *TokenPairUpsertOne) UpdateTokenTwoID() *TokenPairUpsertOne {
	return u.Update(func(s *TokenPairUpsert) {
		s.UpdateTokenTwoID()
	})
}

// SetRemark sets the "remark" field.
func (u *TokenPairUpsertOne) SetRemark(v string) *TokenPairUpsertOne {
	return u.Update(func(s *TokenPairUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *TokenPairUpsertOne) UpdateRemark() *TokenPairUpsertOne {
	return u.Update(func(s *TokenPairUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *TokenPairUpsertOne) ClearRemark() *TokenPairUpsertOne {
	return u.Update(func(s *TokenPairUpsert) {
		s.ClearRemark()
	})
}

// Exec executes the query.
func (u *TokenPairUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TokenPairCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TokenPairUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TokenPairUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TokenPairUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TokenPairCreateBulk is the builder for creating many TokenPair entities in bulk.
type TokenPairCreateBulk struct {
	config
	builders []*TokenPairCreate
	conflict []sql.ConflictOption
}

// Save creates the TokenPair entities in the database.
func (tpcb *TokenPairCreateBulk) Save(ctx context.Context) ([]*TokenPair, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tpcb.builders))
	nodes := make([]*TokenPair, len(tpcb.builders))
	mutators := make([]Mutator, len(tpcb.builders))
	for i := range tpcb.builders {
		func(i int, root context.Context) {
			builder := tpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TokenPairMutation)
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
					_, err = mutators[i+1].Mutate(root, tpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tpcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, tpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tpcb *TokenPairCreateBulk) SaveX(ctx context.Context) []*TokenPair {
	v, err := tpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tpcb *TokenPairCreateBulk) Exec(ctx context.Context) error {
	_, err := tpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpcb *TokenPairCreateBulk) ExecX(ctx context.Context) {
	if err := tpcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TokenPair.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TokenPairUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (tpcb *TokenPairCreateBulk) OnConflict(opts ...sql.ConflictOption) *TokenPairUpsertBulk {
	tpcb.conflict = opts
	return &TokenPairUpsertBulk{
		create: tpcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TokenPair.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tpcb *TokenPairCreateBulk) OnConflictColumns(columns ...string) *TokenPairUpsertBulk {
	tpcb.conflict = append(tpcb.conflict, sql.ConflictColumns(columns...))
	return &TokenPairUpsertBulk{
		create: tpcb,
	}
}

// TokenPairUpsertBulk is the builder for "upsert"-ing
// a bulk of TokenPair nodes.
type TokenPairUpsertBulk struct {
	create *TokenPairCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TokenPair.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(tokenpair.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TokenPairUpsertBulk) UpdateNewValues() *TokenPairUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(tokenpair.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TokenPair.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TokenPairUpsertBulk) Ignore() *TokenPairUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TokenPairUpsertBulk) DoNothing() *TokenPairUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TokenPairCreateBulk.OnConflict
// documentation for more info.
func (u *TokenPairUpsertBulk) Update(set func(*TokenPairUpsert)) *TokenPairUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TokenPairUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TokenPairUpsertBulk) SetCreatedAt(v uint32) *TokenPairUpsertBulk {
	return u.Update(func(s *TokenPairUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *TokenPairUpsertBulk) AddCreatedAt(v uint32) *TokenPairUpsertBulk {
	return u.Update(func(s *TokenPairUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TokenPairUpsertBulk) UpdateCreatedAt() *TokenPairUpsertBulk {
	return u.Update(func(s *TokenPairUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TokenPairUpsertBulk) SetUpdatedAt(v uint32) *TokenPairUpsertBulk {
	return u.Update(func(s *TokenPairUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *TokenPairUpsertBulk) AddUpdatedAt(v uint32) *TokenPairUpsertBulk {
	return u.Update(func(s *TokenPairUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TokenPairUpsertBulk) UpdateUpdatedAt() *TokenPairUpsertBulk {
	return u.Update(func(s *TokenPairUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TokenPairUpsertBulk) SetDeletedAt(v uint32) *TokenPairUpsertBulk {
	return u.Update(func(s *TokenPairUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *TokenPairUpsertBulk) AddDeletedAt(v uint32) *TokenPairUpsertBulk {
	return u.Update(func(s *TokenPairUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TokenPairUpsertBulk) UpdateDeletedAt() *TokenPairUpsertBulk {
	return u.Update(func(s *TokenPairUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetTokenOneID sets the "token_one_id" field.
func (u *TokenPairUpsertBulk) SetTokenOneID(v uint32) *TokenPairUpsertBulk {
	return u.Update(func(s *TokenPairUpsert) {
		s.SetTokenOneID(v)
	})
}

// AddTokenOneID adds v to the "token_one_id" field.
func (u *TokenPairUpsertBulk) AddTokenOneID(v uint32) *TokenPairUpsertBulk {
	return u.Update(func(s *TokenPairUpsert) {
		s.AddTokenOneID(v)
	})
}

// UpdateTokenOneID sets the "token_one_id" field to the value that was provided on create.
func (u *TokenPairUpsertBulk) UpdateTokenOneID() *TokenPairUpsertBulk {
	return u.Update(func(s *TokenPairUpsert) {
		s.UpdateTokenOneID()
	})
}

// SetTokenTwoID sets the "token_two_id" field.
func (u *TokenPairUpsertBulk) SetTokenTwoID(v uint32) *TokenPairUpsertBulk {
	return u.Update(func(s *TokenPairUpsert) {
		s.SetTokenTwoID(v)
	})
}

// AddTokenTwoID adds v to the "token_two_id" field.
func (u *TokenPairUpsertBulk) AddTokenTwoID(v uint32) *TokenPairUpsertBulk {
	return u.Update(func(s *TokenPairUpsert) {
		s.AddTokenTwoID(v)
	})
}

// UpdateTokenTwoID sets the "token_two_id" field to the value that was provided on create.
func (u *TokenPairUpsertBulk) UpdateTokenTwoID() *TokenPairUpsertBulk {
	return u.Update(func(s *TokenPairUpsert) {
		s.UpdateTokenTwoID()
	})
}

// SetRemark sets the "remark" field.
func (u *TokenPairUpsertBulk) SetRemark(v string) *TokenPairUpsertBulk {
	return u.Update(func(s *TokenPairUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *TokenPairUpsertBulk) UpdateRemark() *TokenPairUpsertBulk {
	return u.Update(func(s *TokenPairUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *TokenPairUpsertBulk) ClearRemark() *TokenPairUpsertBulk {
	return u.Update(func(s *TokenPairUpsert) {
		s.ClearRemark()
	})
}

// Exec executes the query.
func (u *TokenPairUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TokenPairCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TokenPairCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TokenPairUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}