package api

import (
	apiv1 "github.com/DarrenMannuela/svc-outcrop-sample/API/v1"
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
	v1.Get("/outcrop-sample/:num", apiv1.GetRandomOutcropsampleData)

	v1.Get("/outcrop-sample-afe", apiv1.GetAfe)
	v1.Post("/outcrop-sample-afe", apiv1.SetAfe)
	v1.Get("/outcrop-sample-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/outcrop-sample-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/outcrop-sample-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/outcrop-sample-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/outcrop-sample-workspace", apiv1.GetWorkspace)
	v1.Post("/outcrop-sample-workspace", apiv1.SetWorkspace)
	v1.Get("/outcrop-sample-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/outcrop-sample-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/outcrop-sample-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/outcrop-sample-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/outcrop-sample", apiv1.GetOutcropsample)
	v1.Post("/outcrop-sample", apiv1.SetOutcropsample)
	v1.Get("/outcrop-sample/:id", apiv1.GetOutcropsampleById)
	v1.Put("/outcrop-sample/:id", apiv1.UpdateOutcropsample)
	v1.Patch("/outcrop-sample/:id", apiv1.PatchOutcropsample)
	v1.Delete("/outcrop-sample/:id", apiv1.DeleteOutcropsample)

	return app
    }
    