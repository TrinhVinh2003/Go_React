package router

import (
	"project/vnexpress/controller"

	"github.com/gofiber/fiber/v2"
)

// thiết lập thông tin route
func SetupRoutes(app *fiber.App) {
	app.Get("api/vnexpress", controller.ArticleList)
	app.Get("api/vnexpress/:id", controller.ArticleDetail)
	app.Post("api/vnexpress/", controller.ArticleCreate)
	app.Put("api/vnexpress/:id", controller.ArticleUpdate)
	app.Delete("api/vnexpress/:id", controller.ArticleDelete)

}
