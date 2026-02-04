package schemas

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

type CopyOptions struct {
	User     *model.User
	SchemaID model.ID
}

var nameCopyCounter = regexp.MustCompile(`\((\d+)\)$`)

type copyingSource struct {
	*model.Schema
	FolderID *model.ID
}

func (c *Client) CopyTx(ctx context.Context, tx *gorm.DB, options *CopyOptions) (*model.Schema, error) {
	var source copyingSource
	err := c.GetOut(ctx, &source,
		repository.Select("schemas.*", "folder_id"),
		repository.IDEq(options.SchemaID),
		IncludeUserSchemaScope(options.User.ID),
	)
	if err != nil {
		return nil, err
	}

	copyName, err := c.findLastCopiedByName(ctx, options.User.ID, source.Name)
	if err != nil {
		return nil, err
	}

	return c.CreateTx(ctx, tx, &CreateOptions{
		User:            options.User,
		Name:            copyName,
		Layout:          source.Layout,
		BackgroundColor: source.BackgroundColor,
		Palette:         source.Palette.Data(),
		Size:            source.Size.Data(),
		Beads:           source.Beads.Data(),
		FolderID:        source.FolderID,
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

func (c *Client) buildCopyNamePattern(sourceName string) string {
	counterMatch := nameCopyCounter.FindStringSubmatch(sourceName)
	if len(counterMatch) == 0 {
		return sourceName + " (%)"
	}
	return strings.Replace(sourceName, counterMatch[0], "(%)", 1)
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
