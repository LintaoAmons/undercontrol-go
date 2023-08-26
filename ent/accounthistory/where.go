// Code generated by ent, DO NOT EDIT.

package accounthistory

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/LintaoAmons/undercontrol/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEQ(FieldName, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEQ(FieldAmount, v))
}

// CurrencyCode applies equality check predicate on the "currency_code" field. It's identical to CurrencyCodeEQ.
func CurrencyCode(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEQ(FieldCurrencyCode, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEQ(FieldUpdatedBy, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldContainsFold(FieldName, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldLTE(FieldAmount, v))
}

// AmountContains applies the Contains predicate on the "amount" field.
func AmountContains(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldContains(FieldAmount, v))
}

// AmountHasPrefix applies the HasPrefix predicate on the "amount" field.
func AmountHasPrefix(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldHasPrefix(FieldAmount, v))
}

// AmountHasSuffix applies the HasSuffix predicate on the "amount" field.
func AmountHasSuffix(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldHasSuffix(FieldAmount, v))
}

// AmountIsNil applies the IsNil predicate on the "amount" field.
func AmountIsNil() predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldIsNull(FieldAmount))
}

// AmountNotNil applies the NotNil predicate on the "amount" field.
func AmountNotNil() predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNotNull(FieldAmount))
}

// AmountEqualFold applies the EqualFold predicate on the "amount" field.
func AmountEqualFold(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEqualFold(FieldAmount, v))
}

// AmountContainsFold applies the ContainsFold predicate on the "amount" field.
func AmountContainsFold(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldContainsFold(FieldAmount, v))
}

// CurrencyCodeEQ applies the EQ predicate on the "currency_code" field.
func CurrencyCodeEQ(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEQ(FieldCurrencyCode, v))
}

// CurrencyCodeNEQ applies the NEQ predicate on the "currency_code" field.
func CurrencyCodeNEQ(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNEQ(FieldCurrencyCode, v))
}

// CurrencyCodeIn applies the In predicate on the "currency_code" field.
func CurrencyCodeIn(vs ...string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldIn(FieldCurrencyCode, vs...))
}

// CurrencyCodeNotIn applies the NotIn predicate on the "currency_code" field.
func CurrencyCodeNotIn(vs ...string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNotIn(FieldCurrencyCode, vs...))
}

// CurrencyCodeGT applies the GT predicate on the "currency_code" field.
func CurrencyCodeGT(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldGT(FieldCurrencyCode, v))
}

// CurrencyCodeGTE applies the GTE predicate on the "currency_code" field.
func CurrencyCodeGTE(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldGTE(FieldCurrencyCode, v))
}

// CurrencyCodeLT applies the LT predicate on the "currency_code" field.
func CurrencyCodeLT(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldLT(FieldCurrencyCode, v))
}

// CurrencyCodeLTE applies the LTE predicate on the "currency_code" field.
func CurrencyCodeLTE(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldLTE(FieldCurrencyCode, v))
}

// CurrencyCodeContains applies the Contains predicate on the "currency_code" field.
func CurrencyCodeContains(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldContains(FieldCurrencyCode, v))
}

// CurrencyCodeHasPrefix applies the HasPrefix predicate on the "currency_code" field.
func CurrencyCodeHasPrefix(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldHasPrefix(FieldCurrencyCode, v))
}

// CurrencyCodeHasSuffix applies the HasSuffix predicate on the "currency_code" field.
func CurrencyCodeHasSuffix(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldHasSuffix(FieldCurrencyCode, v))
}

// CurrencyCodeIsNil applies the IsNil predicate on the "currency_code" field.
func CurrencyCodeIsNil() predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldIsNull(FieldCurrencyCode))
}

// CurrencyCodeNotNil applies the NotNil predicate on the "currency_code" field.
func CurrencyCodeNotNil() predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNotNull(FieldCurrencyCode))
}

// CurrencyCodeEqualFold applies the EqualFold predicate on the "currency_code" field.
func CurrencyCodeEqualFold(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEqualFold(FieldCurrencyCode, v))
}

// CurrencyCodeContainsFold applies the ContainsFold predicate on the "currency_code" field.
func CurrencyCodeContainsFold(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldContainsFold(FieldCurrencyCode, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNotNull(FieldCreatedAt))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNotNull(FieldUpdatedAt))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.AccountHistory {
	return predicate.AccountHistory(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// HasAccount applies the HasEdge predicate on the "account" edge.
func HasAccount() predicate.AccountHistory {
	return predicate.AccountHistory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountWith applies the HasEdge predicate on the "account" edge with a given conditions (other predicates).
func HasAccountWith(preds ...predicate.Account) predicate.AccountHistory {
	return predicate.AccountHistory(func(s *sql.Selector) {
		step := newAccountStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AccountHistory) predicate.AccountHistory {
	return predicate.AccountHistory(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AccountHistory) predicate.AccountHistory {
	return predicate.AccountHistory(func(s *sql.Selector) {
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
func Not(p predicate.AccountHistory) predicate.AccountHistory {
	return predicate.AccountHistory(func(s *sql.Selector) {
		p(s.Not())
	})
}