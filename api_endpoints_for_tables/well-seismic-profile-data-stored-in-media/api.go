package api

import (
	apiv1 "github.com/DarrenMannuela/svc-well-seismic-profile-data-stored-in-media/API/v1"
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
	v1.Get("/well-seismic-profile-data-stored-in-media/:num", apiv1.GetRandomWellseismicprofiledatastoredinmediaData)

	v1.Get("/well-seismic-profile-data-stored-in-media-afe", apiv1.GetAfe)
	v1.Post("/well-seismic-profile-data-stored-in-media-afe", apiv1.SetAfe)
	v1.Get("/well-seismic-profile-data-stored-in-media-afe/:afe", apiv1.GetAfeByNum)
	v1.Put("/well-seismic-profile-data-stored-in-media-afe/:afe", apiv1.UpdateAfe)
	v1.Patch("/well-seismic-profile-data-stored-in-media-afe/:afe", apiv1.PatchAfe)
	v1.Delete("/well-seismic-profile-data-stored-in-media-afe/:afe", apiv1.DeleteAfe)


	v1.Get("/well-seismic-profile-data-stored-in-media-workspace", apiv1.GetWorkspace)
	v1.Post("/well-seismic-profile-data-stored-in-media-workspace", apiv1.SetWorkspace)
	v1.Get("/well-seismic-profile-data-stored-in-media-workspace/:afe", apiv1.GetWorkspaceByAfe)
	v1.Put("/well-seismic-profile-data-stored-in-media-workspace/:id", apiv1.UpdateWorkspace)
	v1.Patch("/well-seismic-profile-data-stored-in-media-workspace/:id", apiv1.PatchWorkspace)
	v1.Delete("/well-seismic-profile-data-stored-in-media-workspace/:id", apiv1.DeleteWorkspace)


	v1.Get("/well-seismic-profile-data-stored-in-media", apiv1.GetWellseismicprofiledatastoredinmedia)
	v1.Post("/well-seismic-profile-data-stored-in-media", apiv1.SetWellseismicprofiledatastoredinmedia)
	v1.Get("/well-seismic-profile-data-stored-in-media/:id", apiv1.GetWellseismicprofiledatastoredinmediaById)
	v1.Put("/well-seismic-profile-data-stored-in-media/:id", apiv1.UpdateWellseismicprofiledatastoredinmedia)
	v1.Patch("/well-seismic-profile-data-stored-in-media/:id", apiv1.PatchWellseismicprofiledatastoredinmedia)
	v1.Delete("/well-seismic-profile-data-stored-in-media/:id", apiv1.DeleteWellseismicprofiledatastoredinmedia)

	return app
    }
    