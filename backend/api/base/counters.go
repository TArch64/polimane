package base

import (
	"encoding/base64"
	"encoding/json"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/model"
	"polimane/backend/services/xor"
)

const (
	userCountersHeader  = "X-UC"
	schemaCounterHeader = "X-SC"
)

type countersPayload struct {
	EntityID model.ID    `json:"entityId"`
	Counters interface{} `json:"counters"`
}

func SetResponseUserCounters(ctx *fiber.Ctx, subscription *model.UserSubscription) {
	setResponseCounter(ctx, userCountersHeader, subscription.UserID, &countersPayload{
		EntityID: subscription.UserID,
		Counters: subscription.Counters,
	})
}

func SetResponseSchemaCounters(ctx *fiber.Ctx, userSchema *model.UserSchema) {
	setResponseCounter(ctx, schemaCounterHeader, userSchema.UserID, &countersPayload{
		EntityID: userSchema.SchemaID,
		Counters: userSchema.Counters,
	})
}

func setResponseCounter(ctx *fiber.Ctx, header string, userID model.ID, payload *countersPayload) {
	bytes, _ := json.Marshal(payload)
	encrypted := xor.Encrypt(bytes, userID.Bytes[:])
	encoded := base64.StdEncoding.EncodeToString(encrypted)
	ctx.Set(header, encoded)
}
