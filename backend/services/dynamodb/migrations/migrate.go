package migrations

import "log"

type Migration func(ctx *Ctx) error

var migrations = []Migration{
	v0,
	v1,
}

func Migrate(ctx *Ctx) error {
	version, err := getTableVersion(ctx)
	if err != nil {
		return err
	}

	if version.IsLatest(len(migrations)) {
		log.Printf("[DynamoDB] Table is already at version %d\n", version.Version)
		return nil
	}

	isLocked, err := isTableLocked(ctx)
	if err != nil {
		return err
	}

	if isLocked {
		return nil
	}

	setTableLock(ctx, true)
	defer setTableLock(ctx, false)

	log.Printf("[DynamoDB] Current version: %d\n", version.Version)

	startVersion := version.NextVersion()

	for i := startVersion; i < len(migrations); i++ {
		log.Printf("[DynamoDB] Running migration %d\n", i)
		if err = migrations[i](ctx); err != nil {
			return err
		}

		log.Printf("[DynamoDB] Migration %d complete\n", i)
		_ = setTableVersion(ctx, version, i)
	}

	return nil
}
