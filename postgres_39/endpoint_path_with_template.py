import re


bibliography = ["Area", "SourceDocument", "RDocumentType", "RLanguage", "RPublicationName", 
    "RSource", "RPpdmRowQuality", "SourceDocAuthor", "BusinessAssociate", "RAuthorType", "RBaCategory", 
    "RBaType", "RBaStatus", "RmDataStore", "BaAddress", "RmDataStoreHier", "RmDataStoreHierLevel", 
    "RStoreStatus", "RAddressType", "RServiceQuality", "RDataStoreType", "CsGeodeticDatum", "CsEllipsoid",
    "CsPrincipalMeridian", "CsPrimeMeridian", "CsCoordinateSystem", "PpdmUnitOfMeasure", "RCoordSystemType",
    "RProjectionType", "RVerticalDatumType", "RCoordCapture", "RCoordCompute", "RCoordQuality", "CsCoordTransform",
    "RPpdmUomUsage", "PpdmMeasurementSystem", "PpdmQuantity", "RCsTransformType", "CsCoordAcquisition", "RAreaType",
    "RDatumOrigin"] 

print_well_report = ['business_associate', 'area', 'field', 'well', 'rm_information_item', 'rm_creator', 'r_media_type', 'r_document_type', 'r_item_category', 'r_item_sub_category', 'rm_physical_item', 'anl_accuracy', 'rm_data_store', 'r_source','ppdm_quality_control']


def process_list(cur_list: list[str]):
    process_list = []
    for item in cur_list:
        item = item.replace("_", " ")
        item = item.title()
        item = item.replace(" ", "")
        process_list.append(item)
    return process_list
    



def names(api_list: list):
    dictionary = {}
    for file in api_list:
        method_name = file.replace(" .go", "")
        endpoint_name = re.sub(r'(?<!^)(?=[A-Z])|(?<=\d)(?=[A-Z])', '-', file.replace(' .go', '')).lower()
        if "seis3-d" in endpoint_name:
            endpoint_name = endpoint_name.replace("seis3-d", "seis-3d")
        if "aof4-pt" in endpoint_name:
            endpoint_name = endpoint_name.replace("aof4-pt", "aof-4pt")
        dictionary[endpoint_name] = method_name
    return dictionary


def make_endpoints(api_list: list, link: str):
    save_to = link
    link = link.lower()
    link = link.replace(" ", "_")
    name = names(api_list)
    lst = []
    # data = ""
    data = rf"""package api

import (
	apiv1 "github.com/DarrenMannuela/pt_gtn_{link}/API/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)"""

    data +="""
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
	v1 := api.Group("/v1", middleware)"""
    for i, key in enumerate(name):
        get_endpoint = f"\n\tv1.Get(\"/{key}\", apiv1.Get{name[key]})\n"
        data += get_endpoint
        post_endpoint = f"\tv1.Post(\"/{key}\", apiv1.Set{name[key]})\n"
        data += post_endpoint
        put_endpoint = f"\tv1.Put(\"/{key}/:id\", apiv1.Update{name[key]})\n"
        data += put_endpoint
        patch_endpoint = f"\tv1.Patch(\"/{key}/:id\", apiv1.Patch{name[key]})\n"
        data += patch_endpoint
        delete_endpoint = f"\tv1.Delete(\"/{key}/:id\", apiv1.Delete{name[key]})\n\n"
        data += delete_endpoint

    data += """	return app
}
"""
    with open(rf"handler_maker/{save_to}/{save_to}_endpoints.go", 'w') as f:
        f.write(data)



if __name__ == '__main__':
    cur_list = process_list(print_well_report)
    make_endpoints(cur_list, 'print_well_report')
