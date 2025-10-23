package userschemas

import (
	"gorm.io/gorm"

	"polimane/backend/model"
)

func (c *Client) CreateTx(tx *gorm.DB, userID, schemaID model.ID, access model.AccessLevel) error {
	userSchema := &model.UserSchema{
		UserID:   userID,
		SchemaID: schemaID,
		Access:   access,
	}

	return tx.Create(userSchema).Error
}
