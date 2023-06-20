package api

import (
	apiv1 "github.com/DarrenMannuela/svc-2d-seismic-summary/API/v1"
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
	v1.Get("/2d-seismic-summary/:num", apiv1.GetRandom2dseismicsummaryData)

	v1.Get("/2d-seismic-summary-afe", apiv1.GetAfe)
	v1.Post("/2d-seismic-summary-afe", apiv1.SetAfe)
	v1.Get("/2d-seismic-summary-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/2d-seismic-summary-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/2d-seismic-summary-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/2d-seismic-summary-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/2d-seismic-summary-workspace", apiv1.GetWorkspace)
	v1.Post("/2d-seismic-summary-workspace", apiv1.SetWorkspace)
	v1.Get("/2d-seismic-summary-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/2d-seismic-summary-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/2d-seismic-summary-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/2d-seismic-summary-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/2d-seismic-summary", apiv1.Get2dseismicsummary)
	v1.Post("/2d-seismic-summary", apiv1.Set2dseismicsummary)
	v1.Get("/2d-seismic-summary/:id", apiv1.Get2dseismicsummaryById)
	v1.Put("/2d-seismic-summary/:id", apiv1.Update2dseismicsummary)
	v1.Patch("/2d-seismic-summary/:id", apiv1.Patch2dseismicsummary)
	v1.Delete("/2d-seismic-summary/:id", apiv1.Delete2dseismicsummary)

	return app
    }
    