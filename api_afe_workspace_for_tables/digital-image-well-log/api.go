package api

import (
	apiv1 "github.com/DarrenMannuela/svc-digital-image-well-log/API/v1"
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
	v1.Get("/digital-image-well-log/:num", apiv1.GetRandomDigitalimagewelllogData)

	v1.Get("/digital-image-well-log-afe", apiv1.GetAfe)
	v1.Post("/digital-image-well-log-afe", apiv1.SetAfe)
	v1.Get("/digital-image-well-log-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/digital-image-well-log-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/digital-image-well-log-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/digital-image-well-log-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/digital-image-well-log-workspace", apiv1.GetWorkspace)
	v1.Post("/digital-image-well-log-workspace", apiv1.SetWorkspace)
	v1.Get("/digital-image-well-log-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/digital-image-well-log-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/digital-image-well-log-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/digital-image-well-log-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/digital-image-well-log", apiv1.GetDigitalimagewelllog)
	v1.Post("/digital-image-well-log", apiv1.SetDigitalimagewelllog)
	v1.Get("/digital-image-well-log/:id", apiv1.GetDigitalimagewelllogById)
	v1.Put("/digital-image-well-log/:id", apiv1.UpdateDigitalimagewelllog)
	v1.Patch("/digital-image-well-log/:id", apiv1.PatchDigitalimagewelllog)
	v1.Delete("/digital-image-well-log/:id", apiv1.DeleteDigitalimagewelllog)

	return app
    }
    