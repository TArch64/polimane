package schemas

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"polimane/backend/model"
	"polimane/backend/repository"
)

type CopyOptions struct {
	User     *model.User
	SchemaID model.ID
}

var nameCopyCounter = regexp.MustCompile(`\((\d+)\)$`)

func (c *Client) Copy(ctx context.Context, options *CopyOptions) (*model.Schema, error) {
	original, err := c.Get(ctx,
		repository.IDEq(options.SchemaID),
		IncludeUserSchemaScope(options.User.ID),
	)
	if err != nil {
		return nil, err
	}

	copyName, err := c.findLastCopiedByName(ctx, options.User.ID, original.Name)
	if err != nil {
		return nil, err
	}

	return c.Create(ctx, &CreateOptions{
		User:            options.User,
		Name:            copyName,
		BackgroundColor: original.BackgroundColor,
		Palette:         original.Palette.Data(),
		Size:            original.Size.Data(),
		Beads:           original.Beads.Data(),
	})
}

func (c *Client) findLastCopiedByName(ctx context.Context, userID model.ID, name string) (string, error) {
	namePattern := c.buildCopyNamePattern(name)
	var names []string

	err := c.GetOut(ctx, &names,
		repository.Select("name"),
		IncludeUserSchemaScope(userID),
		repository.Where("name LIKE ?", namePattern),
		repository.Group("name"),
		repository.Order("MAX(schemas.created_at) DESC"),
		repository.Limit(5),
	)

	if err != nil {
		return "", err
	}

	return c.buildCopyName(namePattern, names), nil
}

func (c *Client) buildCopyNamePattern(originalName string) string {
	counterMatch := nameCopyCounter.FindStringSubmatch(originalName)
	if len(counterMatch) == 0 {
		return originalName + " (%)"
	}
	return strings.Replace(originalName, counterMatch[0], "(%)", 1)
}

func (c *Client) buildCopyName(pattern string, names []string) string {
	maxCounter := 0
	for _, name := range names {
		counterMatch := nameCopyCounter.FindStringSubmatch(name)
		counter, _ := strconv.Atoi(counterMatch[1])
		if counter > maxCounter {
			maxCounter = counter
		}
	}

	return strings.Replace(pattern, "%", strconv.Itoa(maxCounter+1), 1)
}
