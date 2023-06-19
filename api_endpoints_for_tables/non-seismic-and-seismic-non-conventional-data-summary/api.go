package api

import (
	apiv1 "github.com/DarrenMannuela/svc-non-seismic-and-seismic-non-conventional-data-summary/API/v1"
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
	v1.Get("/non-seismic-and-seismic-non-conventional-data-summary/:num", apiv1.GetRandomNonseismicandseismicnonconventionaldatasummaryData)

	v1.Get("/non-seismic-and-seismic-non-conventional-data-summary-afe", apiv1.GetAfe)
	v1.Post("/non-seismic-and-seismic-non-conventional-data-summary-afe", apiv1.SetAfe)
	v1.Get("/non-seismic-and-seismic-non-conventional-data-summary-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/non-seismic-and-seismic-non-conventional-data-summary-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/non-seismic-and-seismic-non-conventional-data-summary-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/non-seismic-and-seismic-non-conventional-data-summary-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/non-seismic-and-seismic-non-conventional-data-summary-workspace", apiv1.GetWorkspace)
	v1.Post("/non-seismic-and-seismic-non-conventional-data-summary-workspace", apiv1.SetWorkspace)
	v1.Get("/non-seismic-and-seismic-non-conventional-data-summary-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/non-seismic-and-seismic-non-conventional-data-summary-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/non-seismic-and-seismic-non-conventional-data-summary-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/non-seismic-and-seismic-non-conventional-data-summary-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/non-seismic-and-seismic-non-conventional-data-summary", apiv1.GetNonseismicandseismicnonconventionaldatasummary)
	v1.Post("/non-seismic-and-seismic-non-conventional-data-summary", apiv1.SetNonseismicandseismicnonconventionaldatasummary)
	v1.Get("/non-seismic-and-seismic-non-conventional-data-summary/:id", apiv1.GetNonseismicandseismicnonconventionaldatasummaryById)
	v1.Put("/non-seismic-and-seismic-non-conventional-data-summary/:id", apiv1.UpdateNonseismicandseismicnonconventionaldatasummary)
	v1.Patch("/non-seismic-and-seismic-non-conventional-data-summary/:id", apiv1.PatchNonseismicandseismicnonconventionaldatasummary)
	v1.Delete("/non-seismic-and-seismic-non-conventional-data-summary/:id", apiv1.DeleteNonseismicandseismicnonconventionaldatasummary)

	return app
    }
    