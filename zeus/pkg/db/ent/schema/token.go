package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/Geapefurit/kline-back/zeus/pkg/db/mixin"
)

type Token struct {
	ent.Schema
}

func (Token) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id"),
		field.String("address"),
		field.String("site"),
		field.String("icon"),
		field.String("name"),
	}
}

func (Token) Indexes() []ent.Index {
	return []ent.Index{}
}
