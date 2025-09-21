package routes

import (
	"goravel/app/http/controllers"

	"github.com/goravel/framework/contracts/route"
)

func AttachmentServiceRoutes(router route.Router) {
	attachmentController := controllers.NewAttachmentController()

	router.Prefix("/general-svc").Group(func(r route.Router) {
		r.Post("/attachments", attachmentController.Create)
		r.Put("/attachments/:id", attachmentController.Update)
	})
}
