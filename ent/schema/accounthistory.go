package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AccountHistory holds the schema definition for the AccountHistory entity.
type AccountHistory struct {
	ent.Schema
}

// Fields of the AccountHistory.
func (AccountHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("name"),
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

// Edges of the AccountHistory.
func (AccountHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Ref("accounthistory").
			Unique(),
	}
}
