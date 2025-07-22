package userschemas

import (
	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/model/modelbase"
)

func (c *Client) DeleteTx(tx *gorm.DB, userID, schemaID modelbase.ID) error {
	return tx.
		Where("user_id = ? AND schema_id = ?", userID, schemaID).
		Delete(&model.UserSchema{}).
		Error
}
