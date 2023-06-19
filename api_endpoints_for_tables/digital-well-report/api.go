package api

import (
	apiv1 "github.com/DarrenMannuela/svc-digital-well-report/API/v1"
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
	v1.Get("/digital-well-report/:num", apiv1.GetRandomDigitalwellreportData)

	v1.Get("/digital-well-report-afe", apiv1.GetAfe)
	v1.Post("/digital-well-report-afe", apiv1.SetAfe)
	v1.Get("/digital-well-report-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/digital-well-report-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/digital-well-report-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/digital-well-report-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/digital-well-report-workspace", apiv1.GetWorkspace)
	v1.Post("/digital-well-report-workspace", apiv1.SetWorkspace)
	v1.Get("/digital-well-report-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/digital-well-report-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/digital-well-report-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/digital-well-report-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/digital-well-report", apiv1.GetDigitalwellreport)
	v1.Post("/digital-well-report", apiv1.SetDigitalwellreport)
	v1.Get("/digital-well-report/:id", apiv1.GetDigitalwellreportById)
	v1.Put("/digital-well-report/:id", apiv1.UpdateDigitalwellreport)
	v1.Patch("/digital-well-report/:id", apiv1.PatchDigitalwellreport)
	v1.Delete("/digital-well-report/:id", apiv1.DeleteDigitalwellreport)

	return app
    }
    