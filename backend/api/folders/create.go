package folders

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
)

type createBody struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
}

func (c *Controller) apiCreate(ctx *fiber.Ctx) (err error) {
	var body createBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	folder := &model.Folder{
		Name:   body.Name,
		UserID: auth.GetSessionUser(ctx).ID,
	}

	err = c.db.Transaction(func(tx *gorm.DB) error {
		if err = c.folders.CreateTx(ctx.Context(), tx, folder); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return ctx.JSON(folder)
}
