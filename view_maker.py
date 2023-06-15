import yaml


# Main function to create the view, pass the name of the view on the 1st parameter and the 2nd is the list of tables needed to make the view
def make_view(view_name: str, tables: list):
    endpoint_holder = {'paths': {}}
    schema_holder = {}
    view_name = view_name.lower()
    title = view_name.title()
    base = {'openapi': '3.0.0', 'info': {'description': rf'This is the swagger API for {title}', 'version': '1.0.0', 'title': rf'{title}', 'termsOfService': 'http://swagger.io/terms/', 
            'contact':{'email': 'darren.mannuela@gmail.com'}, 'license': {'name': 'Apache 2.0', 'url': 'http://www.apache.org/licenses/LICENSE-2.0.html'}},
            'servers': [{'description': rf'{title}', 'url': 'http://localhost:8080/api/v1'},
                        {'description': 'SwaggerHub API Auto Mocking', 'url': rf'https://virtserver.swaggerhub.com/DarrenMannuela/{view_name.replace(" ", "")}/1.0.0'}], 
            
            
            
            'tags':[{'name': 'Query', 'description': rf'All endpoints related to get {title}'}, 
                    {'name': 'Create', 'description': rf'All endpoints related to creating {title}'},
                    {'name': 'Put', 'description': rf'All endpoints related to updating {title}'},
                    {'name': 'Patch', 'description': rf'All endpoints related to patching {title}'},
                    {'name': 'Delete', 'description': rf'All endpoints related to delete {title}'}]}
    

    cur_view = {'components': {'schemas': {}, 'securitySchemes': {'Authorization': {'type': 'http', 'scheme': 'bearer', 'bearerFormat': 'JWT', 'description': 'Please enter the access token'}}}}
    
    for table in tables:


        with open (rf'schema/{table}.yaml', "r") as f:
            schema = yaml.load(f, Loader=yaml.FullLoader)
            schema_holder.update(schema)

        with open (rf'endpoints/{table}.yaml', "r") as f:
            endpoint = yaml.load(f, Loader=yaml.FullLoader)
            endpoint_holder['paths'].update(endpoint)


    
    cur_view['components']['schemas'].update(schema_holder)
    base.update(endpoint_holder)
    base.update(cur_view)


    with open(rf'schema-views/{view_name.replace(" ", "_")}.yaml', "w") as f:
         yaml.dump(base, f, sort_keys=False, default_flow_style=False)


# The list to make the views should look like this, make sure the name of the list exist within the folder of endpoints and schema
print_well_report = ['business_associate', 'area', 'field', 'well', 'rm_information_item', 'rm_creator', 'r_media_type', 'r_document_type', 'r_item_category', 'r_item_sub_category', 'rm_physical_item', 'anl_accuracy', 'rm_data_store', 'r_source','ppdm_quality_control']

bibliography =  ["area", "source_document", "r_document_type", "r_language", "r_publication_name", 
"r_source", "r_ppdm_row_quality", "source_doc_author", "business_associate", "r_author_type", "r_ba_category", 
"r_ba_type", "r_ba_status", "rm_data_store", "ba_address", "rm_data_store_hier", "rm_data_store_hier_level", 
"r_store_status", "r_address_type", "r_service_quality", "r_data_store_type", "cs_geodetic_datum", "cs_ellipsoid",
"cs_principal_meridian", "cs_prime_meridian", "cs_coordinate_system", "ppdm_unit_of_measure", "r_coord_system_type",
"r_projection_type", "r_vertical_datum_type", "r_coord_capture", "r_coord_compute", "r_coord_quality", "cs_coord_transform",
"r_ppdm_uom_usage", "ppdm_measurement_system", "ppdm_quantity", "r_cs_transform_type", "cs_coord_acquisition", "r_area_type",
"r_datum_origin"]

if __name__ == "__main__":
    #Once the function is executed, the .yaml file will be stored within the folder schema-view
    # make_view('Print Well report', print_well_view)
    make_view('print_well_report', print_well_report)