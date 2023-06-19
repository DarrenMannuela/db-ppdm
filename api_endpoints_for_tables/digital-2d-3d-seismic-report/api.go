package api

import (
	apiv1 "github.com/DarrenMannuela/svc-digital-2d-3d-seismic-report/API/v1"
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
	v1.Get("/digital-2d-3d-seismic-report/:num", apiv1.GetRandomDigital2d3dseismicreportData)

	v1.Get("/digital-2d-3d-seismic-report-afe", apiv1.GetAfe)
	v1.Post("/digital-2d-3d-seismic-report-afe", apiv1.SetAfe)
	v1.Get("/digital-2d-3d-seismic-report-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/digital-2d-3d-seismic-report-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/digital-2d-3d-seismic-report-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/digital-2d-3d-seismic-report-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/digital-2d-3d-seismic-report-workspace", apiv1.GetWorkspace)
	v1.Post("/digital-2d-3d-seismic-report-workspace", apiv1.SetWorkspace)
	v1.Get("/digital-2d-3d-seismic-report-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/digital-2d-3d-seismic-report-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/digital-2d-3d-seismic-report-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/digital-2d-3d-seismic-report-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/digital-2d-3d-seismic-report", apiv1.GetDigital2d3dseismicreport)
	v1.Post("/digital-2d-3d-seismic-report", apiv1.SetDigital2d3dseismicreport)
	v1.Get("/digital-2d-3d-seismic-report/:id", apiv1.GetDigital2d3dseismicreportById)
	v1.Put("/digital-2d-3d-seismic-report/:id", apiv1.UpdateDigital2d3dseismicreport)
	v1.Patch("/digital-2d-3d-seismic-report/:id", apiv1.PatchDigital2d3dseismicreport)
	v1.Delete("/digital-2d-3d-seismic-report/:id", apiv1.DeleteDigital2d3dseismicreport)

	return app
    }
    