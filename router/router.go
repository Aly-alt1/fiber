package router

import (
	"fiber3-app/backend/handler"
	"fiber3-app/backend/middleware"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/etag"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/session"
)

func SetupRoutes(app *fiber.App, s *session.Store) {

	api := app.Group("/api")

	// Auth
	authHandler := handler.NewAuthHandler(s)
	api.Post("/login", authHandler.Login)
	api.Post("/logout", authHandler.Logout)
	api.Post("/update-passwd", middleware.Validate, authHandler.UpdatePasswd)
	api.Get("/current-user-acc", middleware.Validate, authHandler.CurrentUserAcc)
	// Template rendering
	api.Get("/email-alert", handler.EmailAlert)
	// Upload file
	api.Post("/upload", handler.Upload)
	api.Post("/upload2", handler.Upload2)

	// Profile
	profile := api.Group("/profiles")
	profileHandler := handler.NewProfileHandler()
	profile.Post("/", middleware.Validate, profileHandler.Create)
	profile.Get("/", middleware.Validate, profileHandler.GetAll)
	profile.Get("/:id", middleware.Validate, profileHandler.GetByID)
	profile.Patch("/:id", middleware.Validate, profileHandler.Update)

	// Mice
	mice := api.Group("/mice")
	miceHandler := handler.NewMouseHandler()
	mice.Post("/", middleware.Validate, miceHandler.Create)
	//mice.Get("/", middleware.Validate, miceHandler.GetAll)
	mice.Get("/", miceHandler.GetAll)
	mice.Get("/:id", middleware.Validate, miceHandler.GetByID)
	mice.Put("/:id", middleware.Validate, miceHandler.Update)

	// Middlewares for performance: cache de assets de Vite (assets/*.js, *.css).
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(logger.New())

}
