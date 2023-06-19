package api

import (
	apiv1 "github.com/DarrenMannuela/svc-print-maps-and-technical-drawing/API/v1"
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
	v1.Get("/print-maps-and-technical-drawing/:num", apiv1.GetRandomPrintmapsandtechnicaldrawingData)

	v1.Get("/print-maps-and-technical-drawing-afe", apiv1.GetAfe)
	v1.Post("/print-maps-and-technical-drawing-afe", apiv1.SetAfe)
	v1.Get("/print-maps-and-technical-drawing-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/print-maps-and-technical-drawing-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/print-maps-and-technical-drawing-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/print-maps-and-technical-drawing-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/print-maps-and-technical-drawing-workspace", apiv1.GetWorkspace)
	v1.Post("/print-maps-and-technical-drawing-workspace", apiv1.SetWorkspace)
	v1.Get("/print-maps-and-technical-drawing-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/print-maps-and-technical-drawing-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/print-maps-and-technical-drawing-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/print-maps-and-technical-drawing-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/print-maps-and-technical-drawing", apiv1.GetPrintmapsandtechnicaldrawing)
	v1.Post("/print-maps-and-technical-drawing", apiv1.SetPrintmapsandtechnicaldrawing)
	v1.Get("/print-maps-and-technical-drawing/:id", apiv1.GetPrintmapsandtechnicaldrawingById)
	v1.Put("/print-maps-and-technical-drawing/:id", apiv1.UpdatePrintmapsandtechnicaldrawing)
	v1.Patch("/print-maps-and-technical-drawing/:id", apiv1.PatchPrintmapsandtechnicaldrawing)
	v1.Delete("/print-maps-and-technical-drawing/:id", apiv1.DeletePrintmapsandtechnicaldrawing)

	return app
    }
    