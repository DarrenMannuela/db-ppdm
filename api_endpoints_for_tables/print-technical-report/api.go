package api

import (
	apiv1 "github.com/DarrenMannuela/svc-print-technical-report/API/v1"
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
	v1.Get("/print-technical-report/:num", apiv1.GetRandomPrinttechnicalreportData)

	v1.Get("/print-technical-report-afe", apiv1.GetAfe)
	v1.Post("/print-technical-report-afe", apiv1.SetAfe)
	v1.Get("/print-technical-report-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/print-technical-report-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/print-technical-report-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/print-technical-report-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/print-technical-report-workspace", apiv1.GetWorkspace)
	v1.Post("/print-technical-report-workspace", apiv1.SetWorkspace)
	v1.Get("/print-technical-report-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/print-technical-report-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/print-technical-report-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/print-technical-report-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/print-technical-report", apiv1.GetPrinttechnicalreport)
	v1.Post("/print-technical-report", apiv1.SetPrinttechnicalreport)
	v1.Get("/print-technical-report/:id", apiv1.GetPrinttechnicalreportById)
	v1.Put("/print-technical-report/:id", apiv1.UpdatePrinttechnicalreport)
	v1.Patch("/print-technical-report/:id", apiv1.PatchPrinttechnicalreport)
	v1.Delete("/print-technical-report/:id", apiv1.DeletePrinttechnicalreport)

	return app
    }
    