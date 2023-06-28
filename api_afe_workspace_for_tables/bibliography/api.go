package api

import (
	apiv1 "github.com/DarrenMannuela/svc-bibliography/API/v1"
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
	v1.Get("/bibliography/:num", apiv1.GetRandomBibliographyData)

	v1.Get("/bibliography-afe", apiv1.GetAfe)
	v1.Post("/bibliography-afe", apiv1.SetAfe)
	v1.Get("/bibliography-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/bibliography-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/bibliography-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/bibliography-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/bibliography-workspace", apiv1.GetWorkspace)
	v1.Post("/bibliography-workspace", apiv1.SetWorkspace)
	v1.Get("/bibliography-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/bibliography-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/bibliography-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/bibliography-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/bibliography", apiv1.GetBibliography)
	v1.Post("/bibliography", apiv1.SetBibliography)
	v1.Get("/bibliography/:id", apiv1.GetBibliographyById)
	v1.Put("/bibliography/:id", apiv1.UpdateBibliography)
	v1.Patch("/bibliography/:id", apiv1.PatchBibliography)
	v1.Delete("/bibliography/:id", apiv1.DeleteBibliography)

	return app
    }
    