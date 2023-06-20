package api

import (
	apiv1 "github.com/DarrenMannuela/svc-field-information/API/v1"
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
	v1.Get("/field-information/:num", apiv1.GetRandomFieldinformationData)

	v1.Get("/field-information-afe", apiv1.GetAfe)
	v1.Post("/field-information-afe", apiv1.SetAfe)
	v1.Get("/field-information-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/field-information-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/field-information-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/field-information-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/field-information-workspace", apiv1.GetWorkspace)
	v1.Post("/field-information-workspace", apiv1.SetWorkspace)
	v1.Get("/field-information-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/field-information-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/field-information-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/field-information-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/field-information", apiv1.GetFieldinformation)
	v1.Post("/field-information", apiv1.SetFieldinformation)
	v1.Get("/field-information/:id", apiv1.GetFieldinformationById)
	v1.Put("/field-information/:id", apiv1.UpdateFieldinformation)
	v1.Patch("/field-information/:id", apiv1.PatchFieldinformation)
	v1.Delete("/field-information/:id", apiv1.DeleteFieldinformation)

	return app
    }
    