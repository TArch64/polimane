package schemas

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"polimane/backend/model"
)

type CopyOptions struct {
	Ctx      context.Context
	User     *model.User
	SchemaID model.ID
}

var nameCopyCounter = regexp.MustCompile(`\((\d+)\)$`)

func makeCopyName(originalName string) string {
	counterMatch := nameCopyCounter.FindStringSubmatch(originalName)
	if len(counterMatch) == 0 {
		return originalName + " (1)"
	}
	counter, _ := strconv.Atoi(counterMatch[1])
	counterStr := strconv.Itoa(counter + 1)
	return strings.ReplaceAll(originalName, counterMatch[0], " ("+counterStr+")")
}

func (c *Client) Copy(options *CopyOptions) (*model.Schema, error) {
	original, err := c.ByID(&ByIDOptions{
		Ctx:      options.Ctx,
		SchemaID: options.SchemaID,
		User:     options.User,
	})

	if err != nil {
		return nil, err
	}

	return c.Create(&CreateOptions{
		Ctx:     options.Ctx,
		User:    options.User,
		Name:    makeCopyName(original.Name),
		Palette: original.Palette,
		Content: original.Content,
	})
}
