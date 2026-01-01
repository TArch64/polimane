package handlerdeleteusers

import (
	"context"
	"log/slog"
	"time"

	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/model"
	"polimane/backend/services/logpersistent"
	"polimane/backend/services/workos"
)

type DeleterWorkos struct {
	affected         *usermanagement.User
	Workos           *workos.Client
	PersistentLogger *logpersistent.Logger
}

func (d *DeleterWorkos) Collect(ctx context.Context, user *model.User) error {
	affected, err := d.Workos.UserManagement.GetUser(ctx, usermanagement.GetUserOpts{
		User: user.WorkosID,
	})
	if err != nil {
		return err
	}
	d.affected = &affected
	return nil
}

func (d *DeleterWorkos) Delete(ctx context.Context) error {
	return d.Workos.UserManagement.DeleteUser(ctx, usermanagement.DeleteUserOpts{
		User: d.affected.ID,
	})
}

func (d *DeleterWorkos) LogResults(ctx context.Context) {
	lastSignedInAt, _ := time.Parse(time.RFC3339, d.affected.LastSignInAt)

	d.PersistentLogger.InfoContext(ctx, "deleted Workos user",
		logpersistent.Operation("user_deletion"),
		slog.String("id", d.affected.ID),
		slog.String("email", d.affected.Email),
		slog.String("first_name", d.affected.FirstName),
		slog.String("last_name", d.affected.LastName),
		slog.Time("last_signed_in_at", lastSignedInAt),
	)
}

func (d *DeleterWorkos) Cleanup() {
	d.affected = nil
}
