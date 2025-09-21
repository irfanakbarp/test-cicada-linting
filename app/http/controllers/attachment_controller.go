package controllers

import (
	"goravel/app/http/requests/attachments"
	"goravel/app/http/responses"
	"goravel/app/models"
	"goravel/app/traits"

	"github.com/google/uuid"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/validation"
)

// AttachmentController manages attachment operations,
// including creating, updating, retrieving attachments, etc.

type AttachmentController struct {
	response traits.ResponseAPI
}

func NewAttachmentController() *AttachmentController {
	return &AttachmentController{
		response: traits.ResponseAPI{},
	}
}

func (c *AttachmentController) Create(ctx http.Context) http.Response {
	return nil
}

func (c *AttachmentController) Update(ctx http.Context) http.Response {
	idStr := ctx.Request().Route("id")
	if idStr == "" {
		return c.response.Error(ctx, http.StatusBadRequest, "Invalid attachment ID", "ID parameter is required")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.response.Error(ctx, http.StatusBadRequest, "Invalid attachment ID format", err.Error())
	}

	request := &attachments.UpdateAttachmentRequest{}
	// Validate attachment request using custom messages and rules
	validator, err := facades.Validation().Make(
		ctx.Request().All(),
		request.Rules(ctx),
		validation.Messages(request.Messages(ctx)),
	)

	if err != nil {
		return c.response.Error(ctx, 500, "Internal validation error", err.Error())
	}

	if validator.Fails() {
		return c.response.Error(ctx, 422, "Validation failed", validator.Errors().All())
	}

	var attachment models.Attachment
	if err := facades.Orm().Query().Where("id", id).First(&attachment); err != nil {
		return c.response.Error(ctx, http.StatusNotFound, "Attachment not found", err.Error())
	}

	// Take user_id from auth-svc for created_by, updated_by, and deleted_by
	userID := ctx.Request().Input("user_id")

	attachment.AttachmentName = ctx.Request().Input("attachment_name")
	attachment.FileName = ctx.Request().Input("file_name")
	attachment.StoragePath = ctx.Request().Input("storage_path")
	attachment.Deadline = ctx.Request().Input("deadline")
	attachment.ReferenceID = ctx.Request().Input("reference_id")
	attachment.UpdatedBy = &userID

	if err := facades.Orm().Query().Save(&attachment); err != nil {
		return c.response.Error(ctx, http.StatusInternalServerError, "Failed to update attachment", err.Error())
	}

	data := responses.NewAttachmentResponse(attachment)
	return c.response.Success(ctx, data, "Attachment updated successfully")
}
