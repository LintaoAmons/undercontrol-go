// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/LintaoAmons/undercontrol/ent/account"
	"github.com/LintaoAmons/undercontrol/ent/accounthistory"
)

// AccountHistory is the model entity for the AccountHistory schema.
type AccountHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount *string `json:"amount,omitempty"`
	// CurrencyCode holds the value of the "currency_code" field.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy *string `json:"created_by,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy *string `json:"updated_by,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountHistoryQuery when eager-loading is set.
	Edges                  AccountHistoryEdges `json:"edges"`
	account_accounthistory *int
	selectValues           sql.SelectValues
}

// AccountHistoryEdges holds the relations/edges for other nodes in the graph.
type AccountHistoryEdges struct {
	// Account holds the value of the account edge.
	Account *Account `json:"account,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AccountOrErr returns the Account value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountHistoryEdges) AccountOrErr() (*Account, error) {
	if e.loadedTypes[0] {
		if e.Account == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: account.Label}
		}
		return e.Account, nil
	}
	return nil, &NotLoadedError{edge: "account"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AccountHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case accounthistory.FieldID:
			values[i] = new(sql.NullInt64)
		case accounthistory.FieldName, accounthistory.FieldAmount, accounthistory.FieldCurrencyCode, accounthistory.FieldCreatedBy, accounthistory.FieldUpdatedBy:
			values[i] = new(sql.NullString)
		case accounthistory.FieldCreatedAt, accounthistory.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case accounthistory.ForeignKeys[0]: // account_accounthistory
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AccountHistory fields.
func (ah *AccountHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case accounthistory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ah.ID = int(value.Int64)
		case accounthistory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ah.Name = value.String
			}
		case accounthistory.FieldAmount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				ah.Amount = new(string)
				*ah.Amount = value.String
			}
		case accounthistory.FieldCurrencyCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field currency_code", values[i])
			} else if value.Valid {
				ah.CurrencyCode = new(string)
				*ah.CurrencyCode = value.String
			}
		case accounthistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ah.CreatedAt = new(time.Time)
				*ah.CreatedAt = value.Time
			}
		case accounthistory.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ah.CreatedBy = new(string)
				*ah.CreatedBy = value.String
			}
		case accounthistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ah.UpdatedAt = new(time.Time)
				*ah.UpdatedAt = value.Time
			}
		case accounthistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ah.UpdatedBy = new(string)
				*ah.UpdatedBy = value.String
			}
		case accounthistory.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field account_accounthistory", value)
			} else if value.Valid {
				ah.account_accounthistory = new(int)
				*ah.account_accounthistory = int(value.Int64)
			}
		default:
			ah.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AccountHistory.
// This includes values selected through modifiers, order, etc.
func (ah *AccountHistory) Value(name string) (ent.Value, error) {
	return ah.selectValues.Get(name)
}

// QueryAccount queries the "account" edge of the AccountHistory entity.
func (ah *AccountHistory) QueryAccount() *AccountQuery {
	return NewAccountHistoryClient(ah.config).QueryAccount(ah)
}

// Update returns a builder for updating this AccountHistory.
// Note that you need to call AccountHistory.Unwrap() before calling this method if this AccountHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (ah *AccountHistory) Update() *AccountHistoryUpdateOne {
	return NewAccountHistoryClient(ah.config).UpdateOne(ah)
}

// Unwrap unwraps the AccountHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ah *AccountHistory) Unwrap() *AccountHistory {
	_tx, ok := ah.config.driver.(*txDriver)
	if !ok {
		panic("ent: AccountHistory is not a transactional entity")
	}
	ah.config.driver = _tx.drv
	return ah
}

// String implements the fmt.Stringer.
func (ah *AccountHistory) String() string {
	var builder strings.Builder
	builder.WriteString("AccountHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ah.ID))
	builder.WriteString("name=")
	builder.WriteString(ah.Name)
	builder.WriteString(", ")
	if v := ah.Amount; v != nil {
		builder.WriteString("amount=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := ah.CurrencyCode; v != nil {
		builder.WriteString("currency_code=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := ah.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := ah.CreatedBy; v != nil {
		builder.WriteString("created_by=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := ah.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := ah.UpdatedBy; v != nil {
		builder.WriteString("updated_by=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// AccountHistories is a parsable slice of AccountHistory.
type AccountHistories []*AccountHistory
