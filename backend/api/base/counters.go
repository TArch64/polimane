package base

import (
	"encoding/base64"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/model"
	"polimane/backend/services/xor"
)

const (
	userCountersHeader = "X-User-Counters"
)

func SetResponseUserCounters(ctx *fiber.Ctx, subscription *model.UserSubscription) {
	counters, _ := subscription.Counters.MarshalJSON()
	setResponseCounter(ctx, userCountersHeader, subscription.UserID, counters)
}

func setResponseCounter(ctx *fiber.Ctx, header string, userID model.ID, counters []byte) {
	encrypted := xor.Encrypt(counters, userID.Bytes[:])
	encoded := base64.StdEncoding.EncodeToString(encrypted)
	ctx.Set(header, encoded)
}
