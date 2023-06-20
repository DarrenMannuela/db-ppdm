package api

import (
	apiv1 "github.com/DarrenMannuela/svc-digital-well-log/API/v1"
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
	v1.Get("/digital-well-log/:num", apiv1.GetRandomDigitalwelllogData)

	v1.Get("/digital-well-log-afe", apiv1.GetAfe)
	v1.Post("/digital-well-log-afe", apiv1.SetAfe)
	v1.Get("/digital-well-log-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/digital-well-log-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/digital-well-log-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/digital-well-log-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/digital-well-log-workspace", apiv1.GetWorkspace)
	v1.Post("/digital-well-log-workspace", apiv1.SetWorkspace)
	v1.Get("/digital-well-log-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/digital-well-log-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/digital-well-log-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/digital-well-log-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/digital-well-log", apiv1.GetDigitalwelllog)
	v1.Post("/digital-well-log", apiv1.SetDigitalwelllog)
	v1.Get("/digital-well-log/:id", apiv1.GetDigitalwelllogById)
	v1.Put("/digital-well-log/:id", apiv1.UpdateDigitalwelllog)
	v1.Patch("/digital-well-log/:id", apiv1.PatchDigitalwelllog)
	v1.Delete("/digital-well-log/:id", apiv1.DeleteDigitalwelllog)

	return app
    }
    