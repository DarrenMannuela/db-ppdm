# pt_gtn

## How to use view_maker.py:
Make sure that the names of tables are aligned with the ones in the folder of (endpoints & schema)
 1. Make a variable containing all the tables used to make the view:
 `print_well_view = ['business_associate', 'area', 'field', 'well', 'rm_information_item', 'rm_creator', 'r_media_type', 'r_document_type',   'r_item_category', 'r_item_sub_category', 'rm_physical_item', 'anl_accuracy', 'rm_data_store', 'r_source','ppdm_quality_control']`
 2. Proceed to set the name of the view and pass both the list and the name of the view into the parameter `make_view('Print Well report', print_well_view)`
 3. After executing the code, the complete `.yaml` file can be found under the folder schema-view as `print_well_report.yaml`

# db-ppdm
# db-ppdm
