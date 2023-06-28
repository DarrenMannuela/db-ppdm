package api

import (
	apiv1 "github.com/DarrenMannuela/svc-basin/API/v1"
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
	v1.Get("/basin/:num", apiv1.GetRandomBasinData)

	v1.Get("/basin-afe", apiv1.GetAfe)
	v1.Post("/basin-afe", apiv1.SetAfe)
	v1.Get("/basin-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/basin-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/basin-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/basin-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/basin-workspace", apiv1.GetWorkspace)
	v1.Post("/basin-workspace", apiv1.SetWorkspace)
	v1.Get("/basin-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/basin-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/basin-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/basin-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/basin", apiv1.GetBasin)
	v1.Post("/basin", apiv1.SetBasin)
	v1.Get("/basin/:id", apiv1.GetBasinById)
	v1.Put("/basin/:id", apiv1.UpdateBasin)
	v1.Patch("/basin/:id", apiv1.PatchBasin)
	v1.Delete("/basin/:id", apiv1.DeleteBasin)

	return app
    }
    