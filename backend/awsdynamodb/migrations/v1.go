package migrations

import (
	"polimane/backend/argon"
	"polimane/backend/env"
	"polimane/backend/model"
)

func v1(ctx *migrationCtx) error {
	passwordHash, err := argon.Hash(env.Env().DefaultUser.Password)
	if err != nil {
		return err
	}

	defaultUser := model.User{
		Base: &model.Base{
			ID: model.NewID(model.PKUser),
			SK: model.NewKey(model.SKUser, env.Env().DefaultUser.User),
		},
		PasswordHash: passwordHash,
	}

	return ctx.Table.Put(defaultUser).Run(ctx)
}
