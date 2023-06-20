package api

import (
	apiv1 "github.com/DarrenMannuela/svc-well-logs/API/v1"
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
	v1.Get("/well-logs/:num", apiv1.GetRandomWelllogsData)

	v1.Get("/well-logs-afe", apiv1.GetAfe)
	v1.Post("/well-logs-afe", apiv1.SetAfe)
	v1.Get("/well-logs-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/well-logs-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/well-logs-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/well-logs-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/well-logs-workspace", apiv1.GetWorkspace)
	v1.Post("/well-logs-workspace", apiv1.SetWorkspace)
	v1.Get("/well-logs-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/well-logs-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/well-logs-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/well-logs-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/well-logs", apiv1.GetWelllogs)
	v1.Post("/well-logs", apiv1.SetWelllogs)
	v1.Get("/well-logs/:id", apiv1.GetWelllogsById)
	v1.Put("/well-logs/:id", apiv1.UpdateWelllogs)
	v1.Patch("/well-logs/:id", apiv1.PatchWelllogs)
	v1.Delete("/well-logs/:id", apiv1.DeleteWelllogs)

	return app
    }
    