package schemas

import "github.com/gofiber/fiber/v2"

func Group(group fiber.Router) {
	group = group.Group("schemas")
	group.Get("", apiList)
	group.Post("", apiCreate)

	group = group.Group(":schemaId")
	group.Get("", apiById)
	group.Delete("", apiDelete)
}
