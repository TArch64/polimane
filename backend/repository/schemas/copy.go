package repositoryschemas

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"polimane/backend/model"
)

type CopyOptions struct {
	User     *model.User
	SchemaID string
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

func Copy(ctx context.Context, options *CopyOptions) (*model.Schema, error) {
	original, err := ById(ctx, options.User, options.SchemaID)
	if err != nil {
		return nil, err
	}

	return Create(ctx, &CreateOptions{
		User:    options.User,
		Name:    makeCopyName(original.Name),
		Palette: original.Palette,
		Content: original.Content,
	})
}
