package logpersistent

import (
	"log/slog"
)

type OperationName string

var (
	OperationSchemaDeletion            = operation("schema_deletion")
	OperationUserDeletion              = operation("user_deletion")
	OperationCleanupExpiredInvitations = operation("cleanup_expired_invitations")
)

func operation(name string) any {
	return slog.String("operation", name)
}
