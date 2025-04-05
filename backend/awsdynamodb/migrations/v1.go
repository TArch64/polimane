package migrations

import (
	"polimane/backend/argon"
	"polimane/backend/env"
	"polimane/backend/model"
)

func v1(ctx *migrationCtx) error {
	config := env.Env().DefaultUser

	passwordHash, err := argon.Hash(config.Password)
	if err != nil {
		return err
	}

	defaultUser := model.NewUser(&model.NewUserOptions{
		Username:     config.User,
		PasswordHash: passwordHash,
	})

	return ctx.Table.Put(defaultUser).Run(ctx)
}
