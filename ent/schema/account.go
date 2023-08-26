package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique(),
		field.String("amount").
			Optional().
			Nillable(),
		field.String("currency_code").
			Optional().
			Nillable(),
		field.Time("created_at").
			Optional().
			Nillable(),
		field.String("created_by").
			Optional().
			Nillable(),
		field.Time("updated_at").
			Optional().
			Nillable(),
		field.String("updated_by").
			Optional().
			Nillable(),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounthistory", AccountHistory.Type),
	}
}
