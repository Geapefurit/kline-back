package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/Geapefurit/kline-back/zeus/pkg/db/mixin"
)

type TokenPair struct {
	ent.Schema
}

func (TokenPair) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (TokenPair) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id"),
		field.Uint32("token_one_id"),
		field.Uint32("token_two_id"),
		field.String("remark").Optional(),
	}
}

func (TokenPair) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("token_one_id", "token_two_id").Unique(),
	}
}
