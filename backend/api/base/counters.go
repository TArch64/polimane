package base

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/model"
)

const (
	userCountersHeader = "X-User-Counters"
)

func SetResponseUserCounters(ctx *fiber.Ctx, subscription *model.UserSubscription) {
	counters, _ := subscription.Counters.MarshalJSON()
	setResponseCounter(ctx, userCountersHeader, counters)
}

func setResponseCounter(ctx *fiber.Ctx, header string, counters []byte) {
	//encoded := base64.StdEncoding.EncodeToString(counters)
	ctx.Set(header, string(counters))
}
