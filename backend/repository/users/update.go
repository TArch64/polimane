package repositoryusers

import (
	"github.com/guregu/dynamo/v2"

	"polimane/backend/model"
	awsdynamodb "polimane/backend/services/dynamodb"
)

func UpdateTx(key model.PrimaryKey, updates model.Updates) *dynamo.Update {
	update := awsdynamodb.Table().
		Update("PK", key.PK()).
		Range("SK", key.SK())

	return updates.Apply(update)
}
