package api

import (
	apiv1 "github.com/DarrenMannuela/svc-well-samples/API/v1"
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
	v1.Get("/well-samples/:num", apiv1.GetRandomWellsamplesData)

	v1.Get("/well-samples-afe", apiv1.GetAfe)
	v1.Post("/well-samples-afe", apiv1.SetAfe)
	v1.Get("/well-samples-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/well-samples-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/well-samples-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/well-samples-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/well-samples-workspace", apiv1.GetWorkspace)
	v1.Post("/well-samples-workspace", apiv1.SetWorkspace)
	v1.Get("/well-samples-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/well-samples-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/well-samples-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/well-samples-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/well-samples", apiv1.GetWellsamples)
	v1.Post("/well-samples", apiv1.SetWellsamples)
	v1.Get("/well-samples/:id", apiv1.GetWellsamplesById)
	v1.Put("/well-samples/:id", apiv1.UpdateWellsamples)
	v1.Patch("/well-samples/:id", apiv1.PatchWellsamples)
	v1.Delete("/well-samples/:id", apiv1.DeleteWellsamples)

	return app
    }
    