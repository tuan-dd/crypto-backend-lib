package orm

import (
	"github.com/tuan-dd/common-lib/utils"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(utils.UTCDate).
			Immutable(),
		field.Time("updated_at").UpdateDefault(utils.UTCDate).Optional().Nillable(),
	}
}
