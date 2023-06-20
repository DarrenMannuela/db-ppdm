package api

import (
	apiv1 "github.com/DarrenMannuela/svc-digital-project-file/API/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func middleware(c *fiber.Ctx) error {
	return c.Next()
}

func handler(c *fiber.Ctx) error {
	c.Status(fiber.StatusOK)
	return c.SendString(c.Path())
}

func Api() *fiber.App {

	// Webserver
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	// Define Root API path
	api := app.Group("/api", middleware)

	// Group v1 API path
	v1 := api.Group("/v1", middleware)
	v1.Get("/digital-project-file/:num", apiv1.GetRandomDigitalprojectfileData)

	v1.Get("/digital-project-file-afe", apiv1.GetAfe)
	v1.Post("/digital-project-file-afe", apiv1.SetAfe)
	v1.Get("/digital-project-file-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/digital-project-file-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/digital-project-file-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/digital-project-file-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/digital-project-file-workspace", apiv1.GetWorkspace)
	v1.Post("/digital-project-file-workspace", apiv1.SetWorkspace)
	v1.Get("/digital-project-file-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/digital-project-file-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/digital-project-file-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/digital-project-file-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/digital-project-file", apiv1.GetDigitalprojectfile)
	v1.Post("/digital-project-file", apiv1.SetDigitalprojectfile)
	v1.Get("/digital-project-file/:id", apiv1.GetDigitalprojectfileById)
	v1.Put("/digital-project-file/:id", apiv1.UpdateDigitalprojectfile)
	v1.Patch("/digital-project-file/:id", apiv1.PatchDigitalprojectfile)
	v1.Delete("/digital-project-file/:id", apiv1.DeleteDigitalprojectfile)

	return app
    }
    