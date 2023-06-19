package api

import (
	apiv1 "github.com/DarrenMannuela/svc-well-data/API/v1"
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
	v1.Get("/well-data/:num", apiv1.GetRandomWelldataData)

	v1.Get("/well-data-afe", apiv1.GetAfe)
	v1.Post("/well-data-afe", apiv1.SetAfe)
	v1.Get("/well-data-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/well-data-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/well-data-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/well-data-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/well-data-workspace", apiv1.GetWorkspace)
	v1.Post("/well-data-workspace", apiv1.SetWorkspace)
	v1.Get("/well-data-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/well-data-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/well-data-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/well-data-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/well-data", apiv1.GetWelldata)
	v1.Post("/well-data", apiv1.SetWelldata)
	v1.Get("/well-data/:id", apiv1.GetWelldataById)
	v1.Put("/well-data/:id", apiv1.UpdateWelldata)
	v1.Patch("/well-data/:id", apiv1.PatchWelldata)
	v1.Delete("/well-data/:id", apiv1.DeleteWelldata)

	return app
    }
    