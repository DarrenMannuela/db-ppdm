package api

import (
	apiv1 "github.com/DarrenMannuela/svc-well-core-sample/API/v1"
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
	v1.Get("/well-core-sample/:num", apiv1.GetRandomWellcoresampleData)

	v1.Get("/well-core-sample-afe", apiv1.GetAfe)
	v1.Post("/well-core-sample-afe", apiv1.SetAfe)
	v1.Get("/well-core-sample-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/well-core-sample-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/well-core-sample-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/well-core-sample-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/well-core-sample-workspace", apiv1.GetWorkspace)
	v1.Post("/well-core-sample-workspace", apiv1.SetWorkspace)
	v1.Get("/well-core-sample-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/well-core-sample-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/well-core-sample-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/well-core-sample-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/well-core-sample", apiv1.GetWellcoresample)
	v1.Post("/well-core-sample", apiv1.SetWellcoresample)
	v1.Get("/well-core-sample/:id", apiv1.GetWellcoresampleById)
	v1.Put("/well-core-sample/:id", apiv1.UpdateWellcoresample)
	v1.Patch("/well-core-sample/:id", apiv1.PatchWellcoresample)
	v1.Delete("/well-core-sample/:id", apiv1.DeleteWellcoresample)

	return app
    }
    