//go:build dev

package migrations

import "context"

func isTableLocked(_ context.Context) (bool, error) {
	return false, nil
}

func setTableLock(_ context.Context, _ bool) {}
