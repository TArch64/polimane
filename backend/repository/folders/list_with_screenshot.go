package folders

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

const listWithScreenshotsSQL = `
WITH schema_screenshots AS (
	SELECT DISTINCT ON (user_schemas.folder_id)
		user_schemas.folder_id,
		schemas.id as screenshot_id,
		schemas.screenshoted_at,
		schemas.background_color
	FROM user_schemas
		JOIN schemas ON user_schemas.schema_id = schemas.id
			AND user_schemas.user_id = '72f666c6-e170-4a32-91c5-54052b3a1756'
			AND screenshoted_at IS NOT NULL
			AND folder_id IS NOT NULL
	ORDER BY user_schemas.folder_id, schemas.created_at
)
SELECT id, name, screenshoted_at, screenshot_id, background_color
FROM folders
	LEFT JOIN schema_screenshots ON folders.id = schema_screenshots.folder_id
WHERE folders.user_id = ?
ORDER BY folders.created_at DESC`

func (c *Client) ListWithScreenshotOut(ctx context.Context, userID model.ID, out interface{}) (err error) {
	err = gorm.
		G[model.Folder](c.DB).
		Raw(listWithScreenshotsSQL, userID).
		Scan(ctx, out)

	if err != nil {
		return err
	}

	return repository.DoAfterScan(out)
}
