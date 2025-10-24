package users

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"github.com/workos/workos-go/v4/pkg/workos_errors"
	"gorm.io/gorm"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemasinvitations "polimane/backend/repository/schemainvitations"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/services/workos"
)

type addUserBody struct {
	Email string `validate:"required,email,max=255" json:"email"`
}

type addResponse struct {
	User       *listUser       `json:"user"`
	Invitation *listInvitation `json:"invitation"`
}

func (c *Controller) apiAdd(ctx *fiber.Ctx) error {
	schemaID, err := base.GetParamID(ctx, schemaIDParam)
	if err != nil {
		return err
	}

	var body addUserBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	requestCtx := ctx.Context()
	user, err := c.users.GeyByEmail(requestCtx, body.Email)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = nil
		err = nil
	}
	if err != nil {
		return err
	}

	var response *addResponse
	currentUser := auth.GetSessionUser(ctx)

	err = c.userSchemas.HasAccess(requestCtx, currentUser.ID, schemaID, model.AccessAdmin)
	if err != nil {
		return err
	}

	if user == nil {
		response, err = c.inviteUser(requestCtx, currentUser, schemaID, body.Email)
	} else {
		response, err = c.addExistingUser(requestCtx, currentUser, schemaID, user)
	}

	if err != nil {
		return err
	}

	return ctx.JSON(response)
}

func (c *Controller) inviteUser(
	ctx context.Context,
	currentUser *model.User,
	schemaID model.ID,
	email string,
) (*addResponse, error) {
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

	err = c.schemaInvitations.Create(ctx, &repositoryschemasinvitations.CreateOptions{
		Email:     email,
		SchemaID:  schemaID,
		Access:    model.AccessRead,
		ExpiresAt: expiresAt,
	})

	if err != nil {
		return nil, err
	}

	return &addResponse{
		Invitation: &listInvitation{
			Email:  email,
			Access: model.AccessRead,
		},
	}, nil
}

func (c *Controller) addExistingUser(
	ctx context.Context,
	currentUser *model.User,
	schemaID model.ID,
	user *model.User,
) (*addResponse, error) {
	if currentUser.ID == user.ID {
		return nil, base.InvalidRequestErr
	}

	userSchema, err := c.userSchemas.Create(ctx, &repositoryuserschemas.CreateOptions{
		UserID:   user.ID,
		SchemaID: schemaID,
		Access:   model.AccessRead,
	})

	if err != nil {
		return nil, err
	}

	return &addResponse{
		User: newUserListItem(user, userSchema),
	}, nil
}
