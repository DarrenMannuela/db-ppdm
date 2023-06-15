package api

import (
	apiv1 "github.com/DarrenMannuela/pt_gtn_bibliography/API/v1"
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
	v1.Get("/area", apiv1.GetArea)
	v1.Post("/area", apiv1.SetArea)
	v1.Put("/area/:id", apiv1.UpdateArea)
	v1.Patch("/area/:id", apiv1.PatchArea)
	v1.Delete("/area/:id", apiv1.DeleteArea)


	v1.Get("/source-document", apiv1.GetSourceDocument)
	v1.Post("/source-document", apiv1.SetSourceDocument)
	v1.Put("/source-document/:id", apiv1.UpdateSourceDocument)
	v1.Patch("/source-document/:id", apiv1.PatchSourceDocument)
	v1.Delete("/source-document/:id", apiv1.DeleteSourceDocument)


	v1.Get("/r-document-type", apiv1.GetRDocumentType)
	v1.Post("/r-document-type", apiv1.SetRDocumentType)
	v1.Put("/r-document-type/:id", apiv1.UpdateRDocumentType)
	v1.Patch("/r-document-type/:id", apiv1.PatchRDocumentType)
	v1.Delete("/r-document-type/:id", apiv1.DeleteRDocumentType)


	v1.Get("/r-language", apiv1.GetRLanguage)
	v1.Post("/r-language", apiv1.SetRLanguage)
	v1.Put("/r-language/:id", apiv1.UpdateRLanguage)
	v1.Patch("/r-language/:id", apiv1.PatchRLanguage)
	v1.Delete("/r-language/:id", apiv1.DeleteRLanguage)


	v1.Get("/r-publication-name", apiv1.GetRPublicationName)
	v1.Post("/r-publication-name", apiv1.SetRPublicationName)
	v1.Put("/r-publication-name/:id", apiv1.UpdateRPublicationName)
	v1.Patch("/r-publication-name/:id", apiv1.PatchRPublicationName)
	v1.Delete("/r-publication-name/:id", apiv1.DeleteRPublicationName)


	v1.Get("/r-source", apiv1.GetRSource)
	v1.Post("/r-source", apiv1.SetRSource)
	v1.Put("/r-source/:id", apiv1.UpdateRSource)
	v1.Patch("/r-source/:id", apiv1.PatchRSource)
	v1.Delete("/r-source/:id", apiv1.DeleteRSource)


	v1.Get("/r-ppdm-row-quality", apiv1.GetRPpdmRowQuality)
	v1.Post("/r-ppdm-row-quality", apiv1.SetRPpdmRowQuality)
	v1.Put("/r-ppdm-row-quality/:id", apiv1.UpdateRPpdmRowQuality)
	v1.Patch("/r-ppdm-row-quality/:id", apiv1.PatchRPpdmRowQuality)
	v1.Delete("/r-ppdm-row-quality/:id", apiv1.DeleteRPpdmRowQuality)


	v1.Get("/source-doc-author", apiv1.GetSourceDocAuthor)
	v1.Post("/source-doc-author", apiv1.SetSourceDocAuthor)
	v1.Put("/source-doc-author/:id", apiv1.UpdateSourceDocAuthor)
	v1.Patch("/source-doc-author/:id", apiv1.PatchSourceDocAuthor)
	v1.Delete("/source-doc-author/:id", apiv1.DeleteSourceDocAuthor)


	v1.Get("/business-associate", apiv1.GetBusinessAssociate)
	v1.Post("/business-associate", apiv1.SetBusinessAssociate)
	v1.Put("/business-associate/:id", apiv1.UpdateBusinessAssociate)
	v1.Patch("/business-associate/:id", apiv1.PatchBusinessAssociate)
	v1.Delete("/business-associate/:id", apiv1.DeleteBusinessAssociate)


	v1.Get("/r-author-type", apiv1.GetRAuthorType)
	v1.Post("/r-author-type", apiv1.SetRAuthorType)
	v1.Put("/r-author-type/:id", apiv1.UpdateRAuthorType)
	v1.Patch("/r-author-type/:id", apiv1.PatchRAuthorType)
	v1.Delete("/r-author-type/:id", apiv1.DeleteRAuthorType)


	v1.Get("/r-ba-category", apiv1.GetRBaCategory)
	v1.Post("/r-ba-category", apiv1.SetRBaCategory)
	v1.Put("/r-ba-category/:id", apiv1.UpdateRBaCategory)
	v1.Patch("/r-ba-category/:id", apiv1.PatchRBaCategory)
	v1.Delete("/r-ba-category/:id", apiv1.DeleteRBaCategory)


	v1.Get("/r-ba-type", apiv1.GetRBaType)
	v1.Post("/r-ba-type", apiv1.SetRBaType)
	v1.Put("/r-ba-type/:id", apiv1.UpdateRBaType)
	v1.Patch("/r-ba-type/:id", apiv1.PatchRBaType)
	v1.Delete("/r-ba-type/:id", apiv1.DeleteRBaType)


	v1.Get("/r-ba-status", apiv1.GetRBaStatus)
	v1.Post("/r-ba-status", apiv1.SetRBaStatus)
	v1.Put("/r-ba-status/:id", apiv1.UpdateRBaStatus)
	v1.Patch("/r-ba-status/:id", apiv1.PatchRBaStatus)
	v1.Delete("/r-ba-status/:id", apiv1.DeleteRBaStatus)


	v1.Get("/rm-data-store", apiv1.GetRmDataStore)
	v1.Post("/rm-data-store", apiv1.SetRmDataStore)
	v1.Put("/rm-data-store/:id", apiv1.UpdateRmDataStore)
	v1.Patch("/rm-data-store/:id", apiv1.PatchRmDataStore)
	v1.Delete("/rm-data-store/:id", apiv1.DeleteRmDataStore)


	v1.Get("/ba-address", apiv1.GetBaAddress)
	v1.Post("/ba-address", apiv1.SetBaAddress)
	v1.Put("/ba-address/:id", apiv1.UpdateBaAddress)
	v1.Patch("/ba-address/:id", apiv1.PatchBaAddress)
	v1.Delete("/ba-address/:id", apiv1.DeleteBaAddress)


	v1.Get("/rm-data-store-hier", apiv1.GetRmDataStoreHier)
	v1.Post("/rm-data-store-hier", apiv1.SetRmDataStoreHier)
	v1.Put("/rm-data-store-hier/:id", apiv1.UpdateRmDataStoreHier)
	v1.Patch("/rm-data-store-hier/:id", apiv1.PatchRmDataStoreHier)
	v1.Delete("/rm-data-store-hier/:id", apiv1.DeleteRmDataStoreHier)


	v1.Get("/rm-data-store-hier-level", apiv1.GetRmDataStoreHierLevel)
	v1.Post("/rm-data-store-hier-level", apiv1.SetRmDataStoreHierLevel)
	v1.Put("/rm-data-store-hier-level/:id", apiv1.UpdateRmDataStoreHierLevel)
	v1.Patch("/rm-data-store-hier-level/:id", apiv1.PatchRmDataStoreHierLevel)
	v1.Delete("/rm-data-store-hier-level/:id", apiv1.DeleteRmDataStoreHierLevel)


	v1.Get("/r-store-status", apiv1.GetRStoreStatus)
	v1.Post("/r-store-status", apiv1.SetRStoreStatus)
	v1.Put("/r-store-status/:id", apiv1.UpdateRStoreStatus)
	v1.Patch("/r-store-status/:id", apiv1.PatchRStoreStatus)
	v1.Delete("/r-store-status/:id", apiv1.DeleteRStoreStatus)


	v1.Get("/r-address-type", apiv1.GetRAddressType)
	v1.Post("/r-address-type", apiv1.SetRAddressType)
	v1.Put("/r-address-type/:id", apiv1.UpdateRAddressType)
	v1.Patch("/r-address-type/:id", apiv1.PatchRAddressType)
	v1.Delete("/r-address-type/:id", apiv1.DeleteRAddressType)


	v1.Get("/r-service-quality", apiv1.GetRServiceQuality)
	v1.Post("/r-service-quality", apiv1.SetRServiceQuality)
	v1.Put("/r-service-quality/:id", apiv1.UpdateRServiceQuality)
	v1.Patch("/r-service-quality/:id", apiv1.PatchRServiceQuality)
	v1.Delete("/r-service-quality/:id", apiv1.DeleteRServiceQuality)


	v1.Get("/r-data-store-type", apiv1.GetRDataStoreType)
	v1.Post("/r-data-store-type", apiv1.SetRDataStoreType)
	v1.Put("/r-data-store-type/:id", apiv1.UpdateRDataStoreType)
	v1.Patch("/r-data-store-type/:id", apiv1.PatchRDataStoreType)
	v1.Delete("/r-data-store-type/:id", apiv1.DeleteRDataStoreType)


	v1.Get("/cs-geodetic-datum", apiv1.GetCsGeodeticDatum)
	v1.Post("/cs-geodetic-datum", apiv1.SetCsGeodeticDatum)
	v1.Put("/cs-geodetic-datum/:id", apiv1.UpdateCsGeodeticDatum)
	v1.Patch("/cs-geodetic-datum/:id", apiv1.PatchCsGeodeticDatum)
	v1.Delete("/cs-geodetic-datum/:id", apiv1.DeleteCsGeodeticDatum)


	v1.Get("/cs-ellipsoid", apiv1.GetCsEllipsoid)
	v1.Post("/cs-ellipsoid", apiv1.SetCsEllipsoid)
	v1.Put("/cs-ellipsoid/:id", apiv1.UpdateCsEllipsoid)
	v1.Patch("/cs-ellipsoid/:id", apiv1.PatchCsEllipsoid)
	v1.Delete("/cs-ellipsoid/:id", apiv1.DeleteCsEllipsoid)


	v1.Get("/cs-principal-meridian", apiv1.GetCsPrincipalMeridian)
	v1.Post("/cs-principal-meridian", apiv1.SetCsPrincipalMeridian)
	v1.Put("/cs-principal-meridian/:id", apiv1.UpdateCsPrincipalMeridian)
	v1.Patch("/cs-principal-meridian/:id", apiv1.PatchCsPrincipalMeridian)
	v1.Delete("/cs-principal-meridian/:id", apiv1.DeleteCsPrincipalMeridian)


	v1.Get("/cs-prime-meridian", apiv1.GetCsPrimeMeridian)
	v1.Post("/cs-prime-meridian", apiv1.SetCsPrimeMeridian)
	v1.Put("/cs-prime-meridian/:id", apiv1.UpdateCsPrimeMeridian)
	v1.Patch("/cs-prime-meridian/:id", apiv1.PatchCsPrimeMeridian)
	v1.Delete("/cs-prime-meridian/:id", apiv1.DeleteCsPrimeMeridian)


	v1.Get("/cs-coordinate-system", apiv1.GetCsCoordinateSystem)
	v1.Post("/cs-coordinate-system", apiv1.SetCsCoordinateSystem)
	v1.Put("/cs-coordinate-system/:id", apiv1.UpdateCsCoordinateSystem)
	v1.Patch("/cs-coordinate-system/:id", apiv1.PatchCsCoordinateSystem)
	v1.Delete("/cs-coordinate-system/:id", apiv1.DeleteCsCoordinateSystem)


	v1.Get("/ppdm-unit-of-measure", apiv1.GetPpdmUnitOfMeasure)
	v1.Post("/ppdm-unit-of-measure", apiv1.SetPpdmUnitOfMeasure)
	v1.Put("/ppdm-unit-of-measure/:id", apiv1.UpdatePpdmUnitOfMeasure)
	v1.Patch("/ppdm-unit-of-measure/:id", apiv1.PatchPpdmUnitOfMeasure)
	v1.Delete("/ppdm-unit-of-measure/:id", apiv1.DeletePpdmUnitOfMeasure)


	v1.Get("/r-coord-system-type", apiv1.GetRCoordSystemType)
	v1.Post("/r-coord-system-type", apiv1.SetRCoordSystemType)
	v1.Put("/r-coord-system-type/:id", apiv1.UpdateRCoordSystemType)
	v1.Patch("/r-coord-system-type/:id", apiv1.PatchRCoordSystemType)
	v1.Delete("/r-coord-system-type/:id", apiv1.DeleteRCoordSystemType)


	v1.Get("/r-projection-type", apiv1.GetRProjectionType)
	v1.Post("/r-projection-type", apiv1.SetRProjectionType)
	v1.Put("/r-projection-type/:id", apiv1.UpdateRProjectionType)
	v1.Patch("/r-projection-type/:id", apiv1.PatchRProjectionType)
	v1.Delete("/r-projection-type/:id", apiv1.DeleteRProjectionType)


	v1.Get("/r-vertical-datum-type", apiv1.GetRVerticalDatumType)
	v1.Post("/r-vertical-datum-type", apiv1.SetRVerticalDatumType)
	v1.Put("/r-vertical-datum-type/:id", apiv1.UpdateRVerticalDatumType)
	v1.Patch("/r-vertical-datum-type/:id", apiv1.PatchRVerticalDatumType)
	v1.Delete("/r-vertical-datum-type/:id", apiv1.DeleteRVerticalDatumType)


	v1.Get("/r-coord-capture", apiv1.GetRCoordCapture)
	v1.Post("/r-coord-capture", apiv1.SetRCoordCapture)
	v1.Put("/r-coord-capture/:id", apiv1.UpdateRCoordCapture)
	v1.Patch("/r-coord-capture/:id", apiv1.PatchRCoordCapture)
	v1.Delete("/r-coord-capture/:id", apiv1.DeleteRCoordCapture)


	v1.Get("/r-coord-compute", apiv1.GetRCoordCompute)
	v1.Post("/r-coord-compute", apiv1.SetRCoordCompute)
	v1.Put("/r-coord-compute/:id", apiv1.UpdateRCoordCompute)
	v1.Patch("/r-coord-compute/:id", apiv1.PatchRCoordCompute)
	v1.Delete("/r-coord-compute/:id", apiv1.DeleteRCoordCompute)


	v1.Get("/r-coord-quality", apiv1.GetRCoordQuality)
	v1.Post("/r-coord-quality", apiv1.SetRCoordQuality)
	v1.Put("/r-coord-quality/:id", apiv1.UpdateRCoordQuality)
	v1.Patch("/r-coord-quality/:id", apiv1.PatchRCoordQuality)
	v1.Delete("/r-coord-quality/:id", apiv1.DeleteRCoordQuality)


	v1.Get("/cs-coord-transform", apiv1.GetCsCoordTransform)
	v1.Post("/cs-coord-transform", apiv1.SetCsCoordTransform)
	v1.Put("/cs-coord-transform/:id", apiv1.UpdateCsCoordTransform)
	v1.Patch("/cs-coord-transform/:id", apiv1.PatchCsCoordTransform)
	v1.Delete("/cs-coord-transform/:id", apiv1.DeleteCsCoordTransform)


	v1.Get("/r-ppdm-uom-usage", apiv1.GetRPpdmUomUsage)
	v1.Post("/r-ppdm-uom-usage", apiv1.SetRPpdmUomUsage)
	v1.Put("/r-ppdm-uom-usage/:id", apiv1.UpdateRPpdmUomUsage)
	v1.Patch("/r-ppdm-uom-usage/:id", apiv1.PatchRPpdmUomUsage)
	v1.Delete("/r-ppdm-uom-usage/:id", apiv1.DeleteRPpdmUomUsage)


	v1.Get("/ppdm-measurement-system", apiv1.GetPpdmMeasurementSystem)
	v1.Post("/ppdm-measurement-system", apiv1.SetPpdmMeasurementSystem)
	v1.Put("/ppdm-measurement-system/:id", apiv1.UpdatePpdmMeasurementSystem)
	v1.Patch("/ppdm-measurement-system/:id", apiv1.PatchPpdmMeasurementSystem)
	v1.Delete("/ppdm-measurement-system/:id", apiv1.DeletePpdmMeasurementSystem)


	v1.Get("/ppdm-quantity", apiv1.GetPpdmQuantity)
	v1.Post("/ppdm-quantity", apiv1.SetPpdmQuantity)
	v1.Put("/ppdm-quantity/:id", apiv1.UpdatePpdmQuantity)
	v1.Patch("/ppdm-quantity/:id", apiv1.PatchPpdmQuantity)
	v1.Delete("/ppdm-quantity/:id", apiv1.DeletePpdmQuantity)


	v1.Get("/r-cs-transform-type", apiv1.GetRCsTransformType)
	v1.Post("/r-cs-transform-type", apiv1.SetRCsTransformType)
	v1.Put("/r-cs-transform-type/:id", apiv1.UpdateRCsTransformType)
	v1.Patch("/r-cs-transform-type/:id", apiv1.PatchRCsTransformType)
	v1.Delete("/r-cs-transform-type/:id", apiv1.DeleteRCsTransformType)


	v1.Get("/cs-coord-acquisition", apiv1.GetCsCoordAcquisition)
	v1.Post("/cs-coord-acquisition", apiv1.SetCsCoordAcquisition)
	v1.Put("/cs-coord-acquisition/:id", apiv1.UpdateCsCoordAcquisition)
	v1.Patch("/cs-coord-acquisition/:id", apiv1.PatchCsCoordAcquisition)
	v1.Delete("/cs-coord-acquisition/:id", apiv1.DeleteCsCoordAcquisition)


	v1.Get("/r-area-type", apiv1.GetRAreaType)
	v1.Post("/r-area-type", apiv1.SetRAreaType)
	v1.Put("/r-area-type/:id", apiv1.UpdateRAreaType)
	v1.Patch("/r-area-type/:id", apiv1.PatchRAreaType)
	v1.Delete("/r-area-type/:id", apiv1.DeleteRAreaType)


	v1.Get("/r-datum-origin", apiv1.GetRDatumOrigin)
	v1.Post("/r-datum-origin", apiv1.SetRDatumOrigin)
	v1.Put("/r-datum-origin/:id", apiv1.UpdateRDatumOrigin)
	v1.Patch("/r-datum-origin/:id", apiv1.PatchRDatumOrigin)
	v1.Delete("/r-datum-origin/:id", apiv1.DeleteRDatumOrigin)

	return app
}
