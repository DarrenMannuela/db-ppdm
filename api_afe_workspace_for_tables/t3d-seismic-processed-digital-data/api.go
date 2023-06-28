package api

import (
	apiv1 "github.com/DarrenMannuela/svc-t3d-seismic-processed-digital-data/API/v1"
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
	v1.Get("/3d-seismic-processed-digital-data/:num", apiv1.GetRandom3dseismicprocesseddigitaldataData)

	v1.Get("/3d-seismic-processed-digital-data-afe", apiv1.GetAfe)
	v1.Post("/3d-seismic-processed-digital-data-afe", apiv1.SetAfe)
	v1.Get("/3d-seismic-processed-digital-data-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/3d-seismic-processed-digital-data-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/3d-seismic-processed-digital-data-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/3d-seismic-processed-digital-data-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/3d-seismic-processed-digital-data-workspace", apiv1.GetWorkspace)
	v1.Post("/3d-seismic-processed-digital-data-workspace", apiv1.SetWorkspace)
	v1.Get("/3d-seismic-processed-digital-data-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/3d-seismic-processed-digital-data-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/3d-seismic-processed-digital-data-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/3d-seismic-processed-digital-data-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/3d-seismic-processed-digital-data", apiv1.Get3dseismicprocesseddigitaldata)
	v1.Post("/3d-seismic-processed-digital-data", apiv1.Set3dseismicprocesseddigitaldata)
	v1.Get("/3d-seismic-processed-digital-data/:id", apiv1.Get3dseismicprocesseddigitaldataById)
	v1.Put("/3d-seismic-processed-digital-data/:id", apiv1.Update3dseismicprocesseddigitaldata)
	v1.Patch("/3d-seismic-processed-digital-data/:id", apiv1.Patch3dseismicprocesseddigitaldata)
	v1.Delete("/3d-seismic-processed-digital-data/:id", apiv1.Delete3dseismicprocesseddigitaldata)

	return app
    }
    