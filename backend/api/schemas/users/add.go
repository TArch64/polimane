package users

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"github.com/workos/workos-go/v4/pkg/workos_errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
	repositoryschemasinvitations "polimane/backend/repository/schemainvitations"
	"polimane/backend/services/workos"
)

type AddUserBody struct {
	BulkOperationBody
	Email string `validate:"required,email,max=255" json:"email"`
}

type AddUserResponse struct {
	User       *ListUser       `json:"user"`
	Invitation *ListInvitation `json:"invitation"`
}

func (c *Controller) Add(ctx *fiber.Ctx) (err error) {
	var body AddUserBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	requestCtx := ctx.Context()
	currentUser := auth.GetSessionUser(ctx)
	err = c.userSchemas.FilterByAccess(requestCtx, currentUser, &body.IDs, model.AccessAdmin)
	if err != nil {
		return err
	}

	user, err := c.users.Get(
		requestCtx,
		repository.Select("id", "email", "first_name", "last_name"),
		repository.EmailEq(body.Email),
	)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = nil
		err = nil
	}
	if err != nil {
		return err
	}

	var response *AddUserResponse

	if user == nil {
		response, err = c.inviteUser(requestCtx, currentUser, body.IDs, body.Email)
	} else {
		response, err = c.addExistingUser(requestCtx, currentUser, body.IDs, user)
	}

	if err != nil {
		return err
	}

	return ctx.JSON(response)
}

func (c *Controller) inviteUser(
	ctx context.Context,
	currentUser *model.User,
	schemaIDs []model.ID,
	email string,
) (*AddUserResponse, error) {
	invitation, err := c.workosClient.UserManagement.SendInvitation(ctx, usermanagement.SendInvitationOpts{
		Email:         email,
		InviterUserID: currentUser.WorkosID,
		ExpiresInDays: 30,
	})

	var httpError workos_errors.HTTPError
	if errors.As(err, &httpError) && workos.GetErrorCode(&httpError) == workos.CodeEmailAlreadyInvited {
		var response usermanagement.ListInvitationsResponse
		response, err = c.workosClient.UserManagement.ListInvitations(ctx, usermanagement.ListInvitationsOpts{
			Email: email,
			Limit: 1,
		})

		if err != nil {
			return nil, err
		}

		invitation = response.Data[0]
	} else if err != nil {
		return nil, err
	}

	expiresAt, _ := time.Parse(time.RFC3339, invitation.ExpiresAt)

	err = c.schemaInvitations.CreateMany(ctx, &repositoryschemasinvitations.CreateManyOptions{
		Email:     email,
		SchemaIDs: schemaIDs,
		Access:    model.AccessRead,
		ExpiresAt: expiresAt,
	})

	if err != nil {
		return nil, err
	}

	return &AddUserResponse{
		Invitation: &ListInvitation{
			Email:  email,
			Access: model.AccessRead,
		},
	}, nil
}

func (c *Controller) addExistingUser(
	ctx context.Context,
	currentUser *model.User,
	schemaIDs []model.ID,
	user *model.User,
) (*AddUserResponse, error) {
	if currentUser.ID == user.ID {
		return nil, base.InvalidRequestErr
	}

	userSchemas := make([]model.UserSchema, len(schemaIDs))
	for idx, schemaID := range schemaIDs {
		userSchemas[idx] = model.UserSchema{
			UserID:   user.ID,
			SchemaID: schemaID,
			Access:   model.AccessRead,
		}
	}

	if err := c.userSchemas.InsertMany(ctx, &userSchemas, clause.OnConflict{DoNothing: true}); err != nil {
		return nil, err
	}

	return &AddUserResponse{
		User: NewUserListItem(user, model.AccessRead),
	}, nil
}
