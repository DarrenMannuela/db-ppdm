package api

import (
	apiv1 "github.com/DarrenMannuela/pt_gtn_print_well_report/API/v1"
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
	v1.Get("/business-associate", apiv1.GetBusinessAssociate)
	v1.Post("/business-associate", apiv1.SetBusinessAssociate)
	v1.Put("/business-associate/:id", apiv1.UpdateBusinessAssociate)
	v1.Patch("/business-associate/:id", apiv1.PatchBusinessAssociate)
	v1.Delete("/business-associate/:id", apiv1.DeleteBusinessAssociate)


	v1.Get("/area", apiv1.GetArea)
	v1.Post("/area", apiv1.SetArea)
	v1.Put("/area/:id", apiv1.UpdateArea)
	v1.Patch("/area/:id", apiv1.PatchArea)
	v1.Delete("/area/:id", apiv1.DeleteArea)


	v1.Get("/field", apiv1.GetField)
	v1.Post("/field", apiv1.SetField)
	v1.Put("/field/:id", apiv1.UpdateField)
	v1.Patch("/field/:id", apiv1.PatchField)
	v1.Delete("/field/:id", apiv1.DeleteField)


	v1.Get("/well", apiv1.GetWell)
	v1.Post("/well", apiv1.SetWell)
	v1.Put("/well/:id", apiv1.UpdateWell)
	v1.Patch("/well/:id", apiv1.PatchWell)
	v1.Delete("/well/:id", apiv1.DeleteWell)


	v1.Get("/rm-information-item", apiv1.GetRmInformationItem)
	v1.Post("/rm-information-item", apiv1.SetRmInformationItem)
	v1.Put("/rm-information-item/:id", apiv1.UpdateRmInformationItem)
	v1.Patch("/rm-information-item/:id", apiv1.PatchRmInformationItem)
	v1.Delete("/rm-information-item/:id", apiv1.DeleteRmInformationItem)


	v1.Get("/rm-creator", apiv1.GetRmCreator)
	v1.Post("/rm-creator", apiv1.SetRmCreator)
	v1.Put("/rm-creator/:id", apiv1.UpdateRmCreator)
	v1.Patch("/rm-creator/:id", apiv1.PatchRmCreator)
	v1.Delete("/rm-creator/:id", apiv1.DeleteRmCreator)


	v1.Get("/r-media-type", apiv1.GetRMediaType)
	v1.Post("/r-media-type", apiv1.SetRMediaType)
	v1.Put("/r-media-type/:id", apiv1.UpdateRMediaType)
	v1.Patch("/r-media-type/:id", apiv1.PatchRMediaType)
	v1.Delete("/r-media-type/:id", apiv1.DeleteRMediaType)


	v1.Get("/r-document-type", apiv1.GetRDocumentType)
	v1.Post("/r-document-type", apiv1.SetRDocumentType)
	v1.Put("/r-document-type/:id", apiv1.UpdateRDocumentType)
	v1.Patch("/r-document-type/:id", apiv1.PatchRDocumentType)
	v1.Delete("/r-document-type/:id", apiv1.DeleteRDocumentType)


	v1.Get("/r-item-category", apiv1.GetRItemCategory)
	v1.Post("/r-item-category", apiv1.SetRItemCategory)
	v1.Put("/r-item-category/:id", apiv1.UpdateRItemCategory)
	v1.Patch("/r-item-category/:id", apiv1.PatchRItemCategory)
	v1.Delete("/r-item-category/:id", apiv1.DeleteRItemCategory)


	v1.Get("/r-item-sub-category", apiv1.GetRItemSubCategory)
	v1.Post("/r-item-sub-category", apiv1.SetRItemSubCategory)
	v1.Put("/r-item-sub-category/:id", apiv1.UpdateRItemSubCategory)
	v1.Patch("/r-item-sub-category/:id", apiv1.PatchRItemSubCategory)
	v1.Delete("/r-item-sub-category/:id", apiv1.DeleteRItemSubCategory)


	v1.Get("/rm-physical-item", apiv1.GetRmPhysicalItem)
	v1.Post("/rm-physical-item", apiv1.SetRmPhysicalItem)
	v1.Put("/rm-physical-item/:id", apiv1.UpdateRmPhysicalItem)
	v1.Patch("/rm-physical-item/:id", apiv1.PatchRmPhysicalItem)
	v1.Delete("/rm-physical-item/:id", apiv1.DeleteRmPhysicalItem)


	v1.Get("/anl-accuracy", apiv1.GetAnlAccuracy)
	v1.Post("/anl-accuracy", apiv1.SetAnlAccuracy)
	v1.Put("/anl-accuracy/:id", apiv1.UpdateAnlAccuracy)
	v1.Patch("/anl-accuracy/:id", apiv1.PatchAnlAccuracy)
	v1.Delete("/anl-accuracy/:id", apiv1.DeleteAnlAccuracy)


	v1.Get("/rm-data-store", apiv1.GetRmDataStore)
	v1.Post("/rm-data-store", apiv1.SetRmDataStore)
	v1.Put("/rm-data-store/:id", apiv1.UpdateRmDataStore)
	v1.Patch("/rm-data-store/:id", apiv1.PatchRmDataStore)
	v1.Delete("/rm-data-store/:id", apiv1.DeleteRmDataStore)


	v1.Get("/r-source", apiv1.GetRSource)
	v1.Post("/r-source", apiv1.SetRSource)
	v1.Put("/r-source/:id", apiv1.UpdateRSource)
	v1.Patch("/r-source/:id", apiv1.PatchRSource)
	v1.Delete("/r-source/:id", apiv1.DeleteRSource)


	v1.Get("/ppdm-quality-control", apiv1.GetPpdmQualityControl)
	v1.Post("/ppdm-quality-control", apiv1.SetPpdmQualityControl)
	v1.Put("/ppdm-quality-control/:id", apiv1.UpdatePpdmQualityControl)
	v1.Patch("/ppdm-quality-control/:id", apiv1.PatchPpdmQualityControl)
	v1.Delete("/ppdm-quality-control/:id", apiv1.DeletePpdmQualityControl)

	return app
}
