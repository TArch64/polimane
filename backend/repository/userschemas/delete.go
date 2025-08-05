package userschemas

import (
	"gorm.io/gorm"

	"polimane/backend/model"
)

func (c *Impl) DeleteTx(tx *gorm.DB, userID, schemaID model.ID) error {
	return tx.
		Where("user_id = ? AND schema_id = ?", userID, schemaID).
		Delete(&model.UserSchema{}).
		Error
}
