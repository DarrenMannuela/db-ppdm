package api

import (
	apiv1 "github.com/DarrenMannuela/svc-well-seismic-profile-digital/API/v1"
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
	v1.Get("/well-seismic-profile-digital/:num", apiv1.GetRandomWellseismicprofiledigitalData)

	v1.Get("/well-seismic-profile-digital-afe", apiv1.GetAfe)
	v1.Post("/well-seismic-profile-digital-afe", apiv1.SetAfe)
	v1.Get("/well-seismic-profile-digital-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/well-seismic-profile-digital-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/well-seismic-profile-digital-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/well-seismic-profile-digital-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/well-seismic-profile-digital-workspace", apiv1.GetWorkspace)
	v1.Post("/well-seismic-profile-digital-workspace", apiv1.SetWorkspace)
	v1.Get("/well-seismic-profile-digital-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/well-seismic-profile-digital-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/well-seismic-profile-digital-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/well-seismic-profile-digital-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/well-seismic-profile-digital", apiv1.GetWellseismicprofiledigital)
	v1.Post("/well-seismic-profile-digital", apiv1.SetWellseismicprofiledigital)
	v1.Get("/well-seismic-profile-digital/:id", apiv1.GetWellseismicprofiledigitalById)
	v1.Put("/well-seismic-profile-digital/:id", apiv1.UpdateWellseismicprofiledigital)
	v1.Patch("/well-seismic-profile-digital/:id", apiv1.PatchWellseismicprofiledigital)
	v1.Delete("/well-seismic-profile-digital/:id", apiv1.DeleteWellseismicprofiledigital)

	return app
    }
    