package routes

import (
	c "go-project/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func Routes(app *fiber.App) {
	api := app.Group("/api")
	UserProfile := api.Group("/UserProfile")

	UserProfile.Get("", c.GetProfileUser)
}

func InetRoutes(app *fiber.App) {

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"testgo": "23012023", //5.0
		},
	}))

	api := app.Group("/api")
	UserProfile := api.Group("/UserProfile")
	UserProfile.Post("/", c.AddUserProfile)
	UserProfile.Delete("/:id", c.RemoveUserFile)
	UserProfile.Put("/:id", c.UpdateUserProfile)
	UserProfile.Get("/group", c.GetUserProfileGroup)
	UserProfile.Get("/:value", c.SearchValue)
}
