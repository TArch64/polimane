package migrations

import (
	"polimane/backend/env"
	"polimane/backend/model"
)

func v1(ctx *migrationCtx) error {
	defaultUser := model.User{
		Base: &model.Base{
			ID: model.NewID(model.PKUser),
			SK: model.NewKey(model.SKUser, env.Env().DefaultUser.User),
		},
		PasswordHash: env.Env().DefaultUser.Password,
	}

	return ctx.Table.Put(defaultUser).Run(ctx)
}
