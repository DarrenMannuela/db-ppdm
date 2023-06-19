package api

import (
	apiv1 "github.com/DarrenMannuela/svc-working-area/API/v1"
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
	v1.Get("/working-area/:num", apiv1.GetRandomWorkingareaData)

	v1.Get("/working-area-afe", apiv1.GetAfe)
	v1.Post("/working-area-afe", apiv1.SetAfe)
	v1.Get("/working-area-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/working-area-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/working-area-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/working-area-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/working-area-workspace", apiv1.GetWorkspace)
	v1.Post("/working-area-workspace", apiv1.SetWorkspace)
	v1.Get("/working-area-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/working-area-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/working-area-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/working-area-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/working-area", apiv1.GetWorkingarea)
	v1.Post("/working-area", apiv1.SetWorkingarea)
	v1.Get("/working-area/:id", apiv1.GetWorkingareaById)
	v1.Put("/working-area/:id", apiv1.UpdateWorkingarea)
	v1.Patch("/working-area/:id", apiv1.PatchWorkingarea)
	v1.Delete("/working-area/:id", apiv1.DeleteWorkingarea)

	return app
    }
    