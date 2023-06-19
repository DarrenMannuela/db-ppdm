package api

import (
	apiv1 "github.com/DarrenMannuela/svc-digital-maps-and-technical-drawing/API/v1"
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
	v1.Get("/digital-maps-and-technical-drawing/:num", apiv1.GetRandomDigitalmapsandtechnicaldrawingData)

	v1.Get("/digital-maps-and-technical-drawing-afe", apiv1.GetAfe)
	v1.Post("/digital-maps-and-technical-drawing-afe", apiv1.SetAfe)
	v1.Get("/digital-maps-and-technical-drawing-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/digital-maps-and-technical-drawing-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/digital-maps-and-technical-drawing-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/digital-maps-and-technical-drawing-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/digital-maps-and-technical-drawing-workspace", apiv1.GetWorkspace)
	v1.Post("/digital-maps-and-technical-drawing-workspace", apiv1.SetWorkspace)
	v1.Get("/digital-maps-and-technical-drawing-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/digital-maps-and-technical-drawing-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/digital-maps-and-technical-drawing-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/digital-maps-and-technical-drawing-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/digital-maps-and-technical-drawing", apiv1.GetDigitalmapsandtechnicaldrawing)
	v1.Post("/digital-maps-and-technical-drawing", apiv1.SetDigitalmapsandtechnicaldrawing)
	v1.Get("/digital-maps-and-technical-drawing/:id", apiv1.GetDigitalmapsandtechnicaldrawingById)
	v1.Put("/digital-maps-and-technical-drawing/:id", apiv1.UpdateDigitalmapsandtechnicaldrawing)
	v1.Patch("/digital-maps-and-technical-drawing/:id", apiv1.PatchDigitalmapsandtechnicaldrawing)
	v1.Delete("/digital-maps-and-technical-drawing/:id", apiv1.DeleteDigitalmapsandtechnicaldrawing)

	return app
    }
    