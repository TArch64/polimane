package migrations

import (
	"maps"
	"slices"

	"github.com/guregu/dynamo/v2"

	"polimane/backend/model"
)

func v2(ctx *Ctx) error {
	migration := &v2Migration{ctx: ctx}
	return migration.migrate()
}

type v2OldSchema struct {
	*model.Base
	Name    string              `json:"name" dynamo:"Name"`
	Palette []string            `json:"palette" dynamo:"Palette"`
	Content model.SchemaContent `json:"content" dynamo:"Content"`
}

type v2Migration struct {
	ctx     *Ctx
	tx      *dynamo.WriteTx
	schemas []*v2OldSchema
}

func (m *v2Migration) migrate() (err error) {
	if err = m.fetch(); err != nil {
		return err
	}

	if len(m.schemas) == 0 {
		return nil
	}

	m.tx = m.ctx.DB.WriteTx()

	userSchemas := make(map[model.ID][]model.PrimaryKey)
	for _, old := range m.schemas {
		if userSchemas[old.PK] == nil {
			userSchemas[old.PK] = make([]model.PrimaryKey, 0)
		}
	}

	userMap, err := m.fetchUsers(slices.Collect(maps.Keys(userSchemas)))
	if err != nil {
		return err
	}

	for _, old := range m.schemas {
		user := userMap[old.PK]
		migrated := m.migrateSchema(user, old)
		userSchemas[old.PK] = append(userSchemas[old.PK], migrated.PrimaryKey())
	}

	for userID, schemaKeys := range userSchemas {
		m.addSchemasToUser(userMap[userID], schemaKeys)
	}

	return m.tx.Run(m.ctx)
}

func (m *v2Migration) fetch() error {
	return m.ctx.Table.Scan().
		Filter("begins_with(PK, ?) AND begins_with(SK, ?)", "USER#", "SCHEMA#").
		All(m.ctx, &m.schemas)
}

func (m *v2Migration) migrateSchema(user *model.User, old *v2OldSchema) *model.Schema {
	schema := &model.Schema{
		Base: &model.Base{
			PK: model.ID(old.SK),
			SK: model.SKSchema,
		},
		UserIDs: []model.PrimaryKey{user.PrimaryKey()},
		Name:    old.Name,
		Palette: old.Palette,
		Content: old.Content,
	}

	put := m.ctx.Table.Put(schema)

	del := m.ctx.Table.
		Delete("PK", old.PK).
		Range("SK", old.SK)

	m.tx.Put(put)
	m.tx.Delete(del)
	return schema
}

func (m *v2Migration) fetchUsers(IDs []model.ID) (map[model.ID]*model.User, error) {
	var err error
	result := make(map[model.ID]*model.User, len(IDs))

	for _, id := range IDs {
		var user model.User

		err = m.ctx.Table.
			Get("PK", id).
			Range("SK", dynamo.BeginsWith, model.SKUserPrefix).
			One(m.ctx, &user)

		if err != nil {
			return nil, err
		}

		result[id] = &user
	}

	return result, nil
}

func (m *v2Migration) addSchemasToUser(user *model.User, schemaKeys []model.PrimaryKey) {
	update := m.ctx.Table.
		Update("PK", user.PK).
		Range("SK", user.SK).
		Set("SchemaIDs", schemaKeys).
		If(model.IfKeyExists)

	m.tx.Update(update)
}
