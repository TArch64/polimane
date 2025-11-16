package schemas

import (
	"time"

	"polimane/backend/model"
	"polimane/backend/repository"
)

type listFolder struct {
	ID              model.ID   `json:"id"`
	Name            string     `json:"name"`
	BackgroundColor string     `json:"backgroundColor"`
	ScreenshotID    *model.ID  `json:"-"`
	ScreenshotedAt  *time.Time `json:"-"`

	// Computed
	ScreenshotPath *string `json:"screenshotPath"`
}

func (l *listFolder) AfterScan() error {
	if l.ScreenshotID != nil {
		l.ScreenshotPath = model.SchemaScreenshotPath(*l.ScreenshotID, l.ScreenshotedAt)
	}
	return nil
}

func (c *Controller) queryFolders(ctx *listContext) (err error) {
	err = c.folders.ListWithScreenshotOut(ctx, ctx.user.ID, &ctx.res.Folders)
	if err != nil {
		return err
	}

	if ctx.res.Folders == nil {
		ctx.res.Folders = []*listFolder{}
	}

	return nil
}

func (c *Controller) countFolders(ctx *listContext) (err error) {
	ctx.foldersTotal, err = c.folders.Count(ctx, c.foldersFilter(ctx)...)
	return err
}

func (c *Controller) foldersFilter(ctx *listContext) []repository.Scope {
	return []repository.Scope{repository.UserIDEq(ctx.user.ID)}
}
