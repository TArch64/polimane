package handlerdeleteusers

import (
	"context"
	"log/slog"

	"polimane/backend/model"
	"polimane/backend/repository"
	repositoryusers "polimane/backend/repository/users"
	"polimane/backend/services/logpersistent"
)

type DeleterUser struct {
	affected         *model.User
	Users            *repositoryusers.Client
	PersistentLogger *logpersistent.Logger
}

func (d *DeleterUser) Collect(_ context.Context, user *model.User) error {
	d.affected = user
	return nil
}

func (d *DeleterUser) Delete(ctx context.Context) error {
	return d.Users.Delete(ctx,
		repository.HardDelete,
		repository.IDEq(d.affected.ID),
	)
}

func (d *DeleterUser) LogResults(ctx context.Context) {
	d.PersistentLogger.InfoContext(ctx, "deleted user record",
		slog.String("operation", "user_deletion"),
		slog.String("id", d.affected.ID.String()),
		slog.String("email", d.affected.Email),
		slog.Time("delete_initiated_at", d.affected.DeletedAt.Time),
	)
}

func (d *DeleterUser) Cleanup() {
	d.affected = nil
}
