package attachments

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UpdateAttachmentRequest struct {
	AttachmentName string `form:"attachment_name" json:"attachment_name"`
	FileName       string `form:"file_name" json:"file_name"`
	StoragePath    string `form:"storage_path" json:"storage_path"`
	Deadline       string `form:"deadline" json:"deadline"`
}

func (r *UpdateAttachmentRequest) Authorize(ctx http.Context) error {
	// The form request class also contains an Authorize method.
	// Within this method, you may determine if the authenticated user actually has the authority to update a given resource.
	// For example, you may determine if a user actually owns a blog comment they are attempting to update.
	return nil
}

func (r *UpdateAttachmentRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateAttachmentRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"attachment_name": "required|max_len:50",
		"file_name":       "required|max_len:70",
		"storage_path":    "required|max_len:255",
		"deadline":        "regex:^\\d{4}/\\d{2}/\\d{2}$",
		"reference_id":    "required",
	}
}

func (r *UpdateAttachmentRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"attachment_name.required": "Attachment name is required",
		"file_name.required":       "File name is required",
		"storage_path.required":    "Storage path is required",
		"reference_id.required":    "Reference ID is required",
		"deadline.regex":           "Deadline format must be yyyy/mm/dd (2025/12/01)",
	}
}

func (r *UpdateAttachmentRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateAttachmentRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
