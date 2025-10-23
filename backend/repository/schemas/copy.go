package schemas

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"polimane/backend/model"
)

type CopyOptions struct {
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

func (i *Impl) Copy(ctx context.Context, options *CopyOptions) (*model.Schema, error) {
	original, err := i.ByID(ctx, &ByIDOptions{
		SchemaID: options.SchemaID,
		User:     options.User,
	})

	if err != nil {
		return nil, err
	}

	return i.Create(ctx, &CreateOptions{
		User:            options.User,
		Name:            makeCopyName(original.Name),
		BackgroundColor: original.BackgroundColor,
		Palette:         original.Palette.Data(),
		Size:            original.Size.Data(),
		Beads:           original.Beads.Data(),
	})
}
