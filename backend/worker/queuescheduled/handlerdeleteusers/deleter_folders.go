package handlerdeleteusers

import (
	"context"
	"log/slog"

	"polimane/backend/model"
	"polimane/backend/repository"
	repositoryfolders "polimane/backend/repository/folders"
	"polimane/backend/services/logpersistent"
)

type affectedFolder struct {
	*model.Identifiable
	*model.Timestamps
	Name string
}

func (a *affectedFolder) getAttrs() []any {
	return []any{
		logpersistent.Operation("user_deletion"),
		slog.String("id", a.ID.String()),
		slog.String("name", a.Name),
		slog.Time("created_at", a.CreatedAt),
	}
}

type DeleterFolders struct {
	affected         []*affectedFolder
	Folders          *repositoryfolders.Client
	PersistentLogger *logpersistent.Logger
}

func (d *DeleterFolders) Collect(ctx context.Context, user *model.User) error {
	return d.Folders.ListOut(ctx, &d.affected,
		repository.UserIDEq(user.ID),
	)
}

func (d *DeleterFolders) Delete(ctx context.Context) error {
	if len(d.affected) == 0 {
		return nil
	}

	var ids []model.ID
	for _, folder := range d.affected {
		ids = append(ids, folder.ID)
	}

	return d.Folders.Delete(ctx,
		repository.IDsIn(ids),
	)
}

func (d *DeleterFolders) LogResults(ctx context.Context) {
	if len(d.affected) == 0 {
		d.PersistentLogger.InfoContext(ctx, "no user Folders to delete for user",
			logpersistent.Operation("user_deletion"),
		)
	} else {
		for _, resource := range d.affected {
			d.PersistentLogger.InfoContext(ctx, "deleted folder due to user deletion",
				logpersistent.Operation("user_deletion"),
				slog.String("id", resource.ID.String()),
				slog.String("name", resource.Name),
				slog.Time("created_at", resource.CreatedAt),
			)
		}
	}
}

func (d *DeleterFolders) Cleanup() {
	d.affected = nil
}
