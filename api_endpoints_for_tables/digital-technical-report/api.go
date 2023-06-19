package api

import (
	apiv1 "github.com/DarrenMannuela/svc-digital-technical-report/API/v1"
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
	v1.Get("/digital-technical-report/:num", apiv1.GetRandomDigitaltechnicalreportData)

	v1.Get("/digital-technical-report-afe", apiv1.GetAfe)
	v1.Post("/digital-technical-report-afe", apiv1.SetAfe)
	v1.Get("/digital-technical-report-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/digital-technical-report-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/digital-technical-report-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/digital-technical-report-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/digital-technical-report-workspace", apiv1.GetWorkspace)
	v1.Post("/digital-technical-report-workspace", apiv1.SetWorkspace)
	v1.Get("/digital-technical-report-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/digital-technical-report-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/digital-technical-report-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/digital-technical-report-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/digital-technical-report", apiv1.GetDigitaltechnicalreport)
	v1.Post("/digital-technical-report", apiv1.SetDigitaltechnicalreport)
	v1.Get("/digital-technical-report/:id", apiv1.GetDigitaltechnicalreportById)
	v1.Put("/digital-technical-report/:id", apiv1.UpdateDigitaltechnicalreport)
	v1.Patch("/digital-technical-report/:id", apiv1.PatchDigitaltechnicalreport)
	v1.Delete("/digital-technical-report/:id", apiv1.DeleteDigitaltechnicalreport)

	return app
    }
    