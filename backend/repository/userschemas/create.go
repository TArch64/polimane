package userschemas

import (
	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/model/modelbase"
)

func (c *Client) CreateTx(tx *gorm.DB, userID, schemaID modelbase.ID) error {
	userSchema := &model.UserSchema{
		UserID:   userID,
		SchemaID: schemaID,
	}

	return tx.Create(userSchema).Error
}
