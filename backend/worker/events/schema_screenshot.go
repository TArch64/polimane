package events

import "polimane/backend/model"

const EventSchemaScreenshot = "schema-screenshot"

type SchemaScreenshotBody struct {
	SchemaID model.ID
}
