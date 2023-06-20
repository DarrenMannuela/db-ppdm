package api

import (
	apiv1 "github.com/DarrenMannuela/svc-seismic-interpretation-data/API/v1"
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
	v1.Get("/seismic-interpretation-data/:num", apiv1.GetRandomSeismicinterpretationdataData)

	v1.Get("/seismic-interpretation-data-afe", apiv1.GetAfe)
	v1.Post("/seismic-interpretation-data-afe", apiv1.SetAfe)
	v1.Get("/seismic-interpretation-data-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/seismic-interpretation-data-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/seismic-interpretation-data-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/seismic-interpretation-data-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/seismic-interpretation-data-workspace", apiv1.GetWorkspace)
	v1.Post("/seismic-interpretation-data-workspace", apiv1.SetWorkspace)
	v1.Get("/seismic-interpretation-data-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/seismic-interpretation-data-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/seismic-interpretation-data-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/seismic-interpretation-data-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/seismic-interpretation-data", apiv1.GetSeismicinterpretationdata)
	v1.Post("/seismic-interpretation-data", apiv1.SetSeismicinterpretationdata)
	v1.Get("/seismic-interpretation-data/:id", apiv1.GetSeismicinterpretationdataById)
	v1.Put("/seismic-interpretation-data/:id", apiv1.UpdateSeismicinterpretationdata)
	v1.Patch("/seismic-interpretation-data/:id", apiv1.PatchSeismicinterpretationdata)
	v1.Delete("/seismic-interpretation-data/:id", apiv1.DeleteSeismicinterpretationdata)

	return app
    }
    