package migrations

import (
	"polimane/backend/env"
	"polimane/backend/model"
	"polimane/backend/services/argon"
)

func v1(ctx *Ctx) error {
	config := env.Env().DefaultUser

	passwordHash, err := argon.HashString(config.Password)
	if err != nil {
		return err
	}

	defaultUser := model.NewUser(&model.NewUserOptions{
		Username:     config.User,
		PasswordHash: passwordHash,
	})

	return ctx.Table.Put(defaultUser).Run(ctx)
}
