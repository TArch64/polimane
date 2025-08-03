package userschemas

import (
	"gorm.io/gorm"

	"polimane/backend/model"
)

func (c *Client) CreateTx(tx *gorm.DB, userID, schemaID model.ID) error {
	userSchema := &model.UserSchema{
		UserID:   userID,
		SchemaID: schemaID,
	}

	return tx.Create(userSchema).Error
}
