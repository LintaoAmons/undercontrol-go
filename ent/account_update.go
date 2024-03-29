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
	"github.com/LintaoAmons/undercontrol/ent/account"
	"github.com/LintaoAmons/undercontrol/ent/accounthistory"
	"github.com/LintaoAmons/undercontrol/ent/predicate"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *AccountUpdate) SetName(s string) *AccountUpdate {
	au.mutation.SetName(s)
	return au
}

// SetAmount sets the "amount" field.
func (au *AccountUpdate) SetAmount(s string) *AccountUpdate {
	au.mutation.SetAmount(s)
	return au
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (au *AccountUpdate) SetNillableAmount(s *string) *AccountUpdate {
	if s != nil {
		au.SetAmount(*s)
	}
	return au
}

// ClearAmount clears the value of the "amount" field.
func (au *AccountUpdate) ClearAmount() *AccountUpdate {
	au.mutation.ClearAmount()
	return au
}

// SetCurrencyCode sets the "currency_code" field.
func (au *AccountUpdate) SetCurrencyCode(s string) *AccountUpdate {
	au.mutation.SetCurrencyCode(s)
	return au
}

// SetNillableCurrencyCode sets the "currency_code" field if the given value is not nil.
func (au *AccountUpdate) SetNillableCurrencyCode(s *string) *AccountUpdate {
	if s != nil {
		au.SetCurrencyCode(*s)
	}
	return au
}

// ClearCurrencyCode clears the value of the "currency_code" field.
func (au *AccountUpdate) ClearCurrencyCode() *AccountUpdate {
	au.mutation.ClearCurrencyCode()
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AccountUpdate) SetCreatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AccountUpdate) SetNillableCreatedAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// ClearCreatedAt clears the value of the "created_at" field.
func (au *AccountUpdate) ClearCreatedAt() *AccountUpdate {
	au.mutation.ClearCreatedAt()
	return au
}

// SetCreatedBy sets the "created_by" field.
func (au *AccountUpdate) SetCreatedBy(s string) *AccountUpdate {
	au.mutation.SetCreatedBy(s)
	return au
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (au *AccountUpdate) SetNillableCreatedBy(s *string) *AccountUpdate {
	if s != nil {
		au.SetCreatedBy(*s)
	}
	return au
}

// ClearCreatedBy clears the value of the "created_by" field.
func (au *AccountUpdate) ClearCreatedBy() *AccountUpdate {
	au.mutation.ClearCreatedBy()
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AccountUpdate) SetUpdatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (au *AccountUpdate) SetNillableUpdatedAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetUpdatedAt(*t)
	}
	return au
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (au *AccountUpdate) ClearUpdatedAt() *AccountUpdate {
	au.mutation.ClearUpdatedAt()
	return au
}

// SetUpdatedBy sets the "updated_by" field.
func (au *AccountUpdate) SetUpdatedBy(s string) *AccountUpdate {
	au.mutation.SetUpdatedBy(s)
	return au
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (au *AccountUpdate) SetNillableUpdatedBy(s *string) *AccountUpdate {
	if s != nil {
		au.SetUpdatedBy(*s)
	}
	return au
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (au *AccountUpdate) ClearUpdatedBy() *AccountUpdate {
	au.mutation.ClearUpdatedBy()
	return au
}

// AddAccounthistoryIDs adds the "accounthistory" edge to the AccountHistory entity by IDs.
func (au *AccountUpdate) AddAccounthistoryIDs(ids ...int) *AccountUpdate {
	au.mutation.AddAccounthistoryIDs(ids...)
	return au
}

// AddAccounthistory adds the "accounthistory" edges to the AccountHistory entity.
func (au *AccountUpdate) AddAccounthistory(a ...*AccountHistory) *AccountUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddAccounthistoryIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// ClearAccounthistory clears all "accounthistory" edges to the AccountHistory entity.
func (au *AccountUpdate) ClearAccounthistory() *AccountUpdate {
	au.mutation.ClearAccounthistory()
	return au
}

// RemoveAccounthistoryIDs removes the "accounthistory" edge to AccountHistory entities by IDs.
func (au *AccountUpdate) RemoveAccounthistoryIDs(ids ...int) *AccountUpdate {
	au.mutation.RemoveAccounthistoryIDs(ids...)
	return au
}

// RemoveAccounthistory removes "accounthistory" edges to AccountHistory entities.
func (au *AccountUpdate) RemoveAccounthistory(a ...*AccountHistory) *AccountUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveAccounthistoryIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(account.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.Amount(); ok {
		_spec.SetField(account.FieldAmount, field.TypeString, value)
	}
	if au.mutation.AmountCleared() {
		_spec.ClearField(account.FieldAmount, field.TypeString)
	}
	if value, ok := au.mutation.CurrencyCode(); ok {
		_spec.SetField(account.FieldCurrencyCode, field.TypeString, value)
	}
	if au.mutation.CurrencyCodeCleared() {
		_spec.ClearField(account.FieldCurrencyCode, field.TypeString)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(account.FieldCreatedAt, field.TypeTime, value)
	}
	if au.mutation.CreatedAtCleared() {
		_spec.ClearField(account.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := au.mutation.CreatedBy(); ok {
		_spec.SetField(account.FieldCreatedBy, field.TypeString, value)
	}
	if au.mutation.CreatedByCleared() {
		_spec.ClearField(account.FieldCreatedBy, field.TypeString)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeTime, value)
	}
	if au.mutation.UpdatedAtCleared() {
		_spec.ClearField(account.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := au.mutation.UpdatedBy(); ok {
		_spec.SetField(account.FieldUpdatedBy, field.TypeString, value)
	}
	if au.mutation.UpdatedByCleared() {
		_spec.ClearField(account.FieldUpdatedBy, field.TypeString)
	}
	if au.mutation.AccounthistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.AccounthistoryTable,
			Columns: []string{account.AccounthistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accounthistory.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedAccounthistoryIDs(); len(nodes) > 0 && !au.mutation.AccounthistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.AccounthistoryTable,
			Columns: []string{account.AccounthistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accounthistory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AccounthistoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.AccounthistoryTable,
			Columns: []string{account.AccounthistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accounthistory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountMutation
}

// SetName sets the "name" field.
func (auo *AccountUpdateOne) SetName(s string) *AccountUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetAmount sets the "amount" field.
func (auo *AccountUpdateOne) SetAmount(s string) *AccountUpdateOne {
	auo.mutation.SetAmount(s)
	return auo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableAmount(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetAmount(*s)
	}
	return auo
}

// ClearAmount clears the value of the "amount" field.
func (auo *AccountUpdateOne) ClearAmount() *AccountUpdateOne {
	auo.mutation.ClearAmount()
	return auo
}

// SetCurrencyCode sets the "currency_code" field.
func (auo *AccountUpdateOne) SetCurrencyCode(s string) *AccountUpdateOne {
	auo.mutation.SetCurrencyCode(s)
	return auo
}

// SetNillableCurrencyCode sets the "currency_code" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableCurrencyCode(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetCurrencyCode(*s)
	}
	return auo
}

// ClearCurrencyCode clears the value of the "currency_code" field.
func (auo *AccountUpdateOne) ClearCurrencyCode() *AccountUpdateOne {
	auo.mutation.ClearCurrencyCode()
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AccountUpdateOne) SetCreatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableCreatedAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (auo *AccountUpdateOne) ClearCreatedAt() *AccountUpdateOne {
	auo.mutation.ClearCreatedAt()
	return auo
}

// SetCreatedBy sets the "created_by" field.
func (auo *AccountUpdateOne) SetCreatedBy(s string) *AccountUpdateOne {
	auo.mutation.SetCreatedBy(s)
	return auo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableCreatedBy(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetCreatedBy(*s)
	}
	return auo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (auo *AccountUpdateOne) ClearCreatedBy() *AccountUpdateOne {
	auo.mutation.ClearCreatedBy()
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AccountUpdateOne) SetUpdatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableUpdatedAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetUpdatedAt(*t)
	}
	return auo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (auo *AccountUpdateOne) ClearUpdatedAt() *AccountUpdateOne {
	auo.mutation.ClearUpdatedAt()
	return auo
}

// SetUpdatedBy sets the "updated_by" field.
func (auo *AccountUpdateOne) SetUpdatedBy(s string) *AccountUpdateOne {
	auo.mutation.SetUpdatedBy(s)
	return auo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableUpdatedBy(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetUpdatedBy(*s)
	}
	return auo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (auo *AccountUpdateOne) ClearUpdatedBy() *AccountUpdateOne {
	auo.mutation.ClearUpdatedBy()
	return auo
}

// AddAccounthistoryIDs adds the "accounthistory" edge to the AccountHistory entity by IDs.
func (auo *AccountUpdateOne) AddAccounthistoryIDs(ids ...int) *AccountUpdateOne {
	auo.mutation.AddAccounthistoryIDs(ids...)
	return auo
}

// AddAccounthistory adds the "accounthistory" edges to the AccountHistory entity.
func (auo *AccountUpdateOne) AddAccounthistory(a ...*AccountHistory) *AccountUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddAccounthistoryIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// ClearAccounthistory clears all "accounthistory" edges to the AccountHistory entity.
func (auo *AccountUpdateOne) ClearAccounthistory() *AccountUpdateOne {
	auo.mutation.ClearAccounthistory()
	return auo
}

// RemoveAccounthistoryIDs removes the "accounthistory" edge to AccountHistory entities by IDs.
func (auo *AccountUpdateOne) RemoveAccounthistoryIDs(ids ...int) *AccountUpdateOne {
	auo.mutation.RemoveAccounthistoryIDs(ids...)
	return auo
}

// RemoveAccounthistory removes "accounthistory" edges to AccountHistory entities.
func (auo *AccountUpdateOne) RemoveAccounthistory(a ...*AccountHistory) *AccountUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveAccounthistoryIDs(ids...)
}

// Where appends a list predicates to the AccountUpdate builder.
func (auo *AccountUpdateOne) Where(ps ...predicate.Account) *AccountUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(account.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.Amount(); ok {
		_spec.SetField(account.FieldAmount, field.TypeString, value)
	}
	if auo.mutation.AmountCleared() {
		_spec.ClearField(account.FieldAmount, field.TypeString)
	}
	if value, ok := auo.mutation.CurrencyCode(); ok {
		_spec.SetField(account.FieldCurrencyCode, field.TypeString, value)
	}
	if auo.mutation.CurrencyCodeCleared() {
		_spec.ClearField(account.FieldCurrencyCode, field.TypeString)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(account.FieldCreatedAt, field.TypeTime, value)
	}
	if auo.mutation.CreatedAtCleared() {
		_spec.ClearField(account.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := auo.mutation.CreatedBy(); ok {
		_spec.SetField(account.FieldCreatedBy, field.TypeString, value)
	}
	if auo.mutation.CreatedByCleared() {
		_spec.ClearField(account.FieldCreatedBy, field.TypeString)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeTime, value)
	}
	if auo.mutation.UpdatedAtCleared() {
		_spec.ClearField(account.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := auo.mutation.UpdatedBy(); ok {
		_spec.SetField(account.FieldUpdatedBy, field.TypeString, value)
	}
	if auo.mutation.UpdatedByCleared() {
		_spec.ClearField(account.FieldUpdatedBy, field.TypeString)
	}
	if auo.mutation.AccounthistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.AccounthistoryTable,
			Columns: []string{account.AccounthistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accounthistory.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedAccounthistoryIDs(); len(nodes) > 0 && !auo.mutation.AccounthistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.AccounthistoryTable,
			Columns: []string{account.AccounthistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accounthistory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AccounthistoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.AccounthistoryTable,
			Columns: []string{account.AccounthistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accounthistory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
