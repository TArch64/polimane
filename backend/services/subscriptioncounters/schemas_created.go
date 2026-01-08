package subscriptioncounters

const countSchemasSQL = `
	SELECT user_schemas.user_id, COUNT(user_schemas.schema_id) AS count
	FROM user_schemas
	WHERE user_schemas.user_id IN @user_ids AND user_schemas.deleted_at IS NULL
	GROUP BY user_schemas.user_id
`

func newSchemasCreated(deps *perUserDeps) *PerUser {
	return newPerUser(&perUserOptions{
		Deps:     deps,
		Name:     "schemasCreated",
		CountSQL: countSchemasSQL,
	})
}
