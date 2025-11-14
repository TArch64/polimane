package schemas

import (
	"polimane/backend/model"
	"polimane/backend/repository"
)

type listFolder struct {
	ID   model.ID `json:"id"`
	Name string   `json:"name"`
}

func (c *Controller) queryFolders(ctx *listContext) (err error) {
	err = c.folders.ListOut(ctx, &ctx.res.Folders,
		repository.Select("id", "name"),
		repository.UserIDEq(ctx.user.ID),
	)

	if err != nil {
		return err
	}

	if ctx.res.Folders == nil {
		ctx.res.Folders = []*listFolder{}
	}

	return nil
}

func (c *Controller) countFolders(ctx *listContext) (err error) {
	ctx.foldersTotal, err = c.folders.Count(ctx,
		repository.UserIDEq(ctx.user.ID),
	)
	return err
}
