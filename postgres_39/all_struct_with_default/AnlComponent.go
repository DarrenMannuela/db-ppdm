package dto

type Anl_component struct {
	Analysis_id              string  `json:"analysis_id" default:"source"`
	Anl_source               string  `json:"anl_source" default:"source"`
	Component_obs_no         int     `json:"component_obs_no" default:"1"`
	Active_ind               *string `json:"active_ind" default:""`
	Activity_obs_no          *int    `json:"activity_obs_no" default:""`
	Application_id           *string `json:"application_id" default:""`
	Area_id                  *string `json:"area_id" default:""`
	Area_type                *string `json:"area_type" default:""`
	Authority_id             *string `json:"authority_id" default:""`
	Ba_organization_id       *string `json:"ba_organization_id" default:""`
	Ba_organization_seq_no   *int    `json:"ba_organization_seq_no" default:""`
	Business_associate_id    *string `json:"business_associate_id" default:""`
	Catalogue_additive_id    *string `json:"catalogue_additive_id" default:""`
	Catalogue_equip_id       *string `json:"catalogue_equip_id" default:""`
	Classification_system_id *string `json:"classification_system_id" default:""`
	Class_level_id           *string `json:"class_level_id" default:""`
	Component_type           *string `json:"component_type" default:""`
	Consent_id               *string `json:"consent_id" default:""`
	Consult_id               *string `json:"consult_id" default:""`
	Contest_id               *string `json:"contest_id" default:""`
	Contract_id              *string `json:"contract_id" default:""`
	Core_id                  *string `json:"core_id" default:""`
	Core_type                *string `json:"core_type" default:""`
	Description_obs_no       *int    `json:"description_obs_no" default:""`
	Ecozone_id               *string `json:"ecozone_id" default:""`
	Effective_date           *string `json:"effective_date" default:""`
	Employee_ba_id           *string `json:"employee_ba_id" default:""`
	Employee_obs_no          *int    `json:"employee_obs_no" default:""`
	Employer_ba_id           *string `json:"employer_ba_id" default:""`
	Entitlement_id           *string `json:"entitlement_id" default:""`
	Equipment_id             *string `json:"equipment_id" default:""`
	Expiry_date              *string `json:"expiry_date" default:""`
	Facility_id              *string `json:"facility_id" default:""`
	Facility_type            *string `json:"facility_type" default:""`
	Field_id                 *string `json:"field_id" default:""`
	Field_station_id         *string `json:"field_station_id" default:""`
	Finance_id               *string `json:"finance_id" default:""`
	Fossil_id                *string `json:"fossil_id" default:""`
	Incident_id              *string `json:"incident_id" default:""`
	Incident_set_id          *string `json:"incident_set_id" default:""`
	Incident_type            *string `json:"incident_type" default:""`
	Information_item_id      *string `json:"information_item_id" default:""`
	Info_item_subtype        *string `json:"info_item_subtype" default:""`
	Instrument_id            *string `json:"instrument_id" default:""`
	Interest_set_id          *string `json:"interest_set_id" default:""`
	Interest_set_seq_no      *int    `json:"interest_set_seq_no" default:""`
	Jurisdiction             *string `json:"jurisdiction" default:""`
	Land_right_id            *string `json:"land_right_id" default:""`
	Land_right_subtype       *string `json:"land_right_subtype" default:""`
	Land_sale_number         *string `json:"land_sale_number" default:""`
	Language                 *string `json:"language" default:""`
	Lithology_log_id         *string `json:"lithology_log_id" default:""`
	Lith_log_source          *string `json:"lith_log_source" default:""`
	Notification_id          *string `json:"notification_id" default:""`
	Obligation_id            *string `json:"obligation_id" default:""`
	Obligation_seq_no        *int    `json:"obligation_seq_no" default:""`
	Paleo_summary_id         *string `json:"paleo_summary_id" default:""`
	Pden_id                  *string `json:"pden_id" default:""`
	Pden_source              *string `json:"pden_source" default:""`
	Pden_subtype             *string `json:"pden_subtype" default:""`
	Physical_item_id         *string `json:"physical_item_id" default:""`
	Pool_id                  *string `json:"pool_id" default:""`
	Ppdm_guid                *string `json:"ppdm_guid" default:""`
	Ppdm_system_id           *string `json:"ppdm_system_id" default:""`
	Ppdm_table_name          *string `json:"ppdm_table_name" default:""`
	Product_type             *string `json:"product_type" default:""`
	Prod_string_source       *string `json:"prod_string_source" default:""`
	Project_id               *string `json:"project_id" default:""`
	Pr_str_form_obs_no       *int    `json:"pr_str_form_obs_no" default:""`
	Rate_schedule_id         *string `json:"rate_schedule_id" default:""`
	Recovery_obs_no          *int    `json:"recovery_obs_no" default:""`
	Referenced_guid          *string `json:"referenced_guid" default:""`
	Referenced_system_id     *string `json:"referenced_system_id" default:""`
	Referenced_table_name    *string `json:"referenced_table_name" default:""`
	Remark                   *string `json:"remark" default:""`
	Reported_ind             *string `json:"reported_ind" default:""`
	Resent_id                *string `json:"resent_id" default:""`
	Reserve_class_id         *string `json:"reserve_class_id" default:""`
	Run_num                  *string `json:"run_num" default:""`
	Sample_date              *string `json:"sample_date" default:""`
	Sample_id                *string `json:"sample_id" default:""`
	Sample_source            *string `json:"sample_source" default:""`
	Seis_set_id              *string `json:"seis_set_id" default:""`
	Seis_set_subtype         *string `json:"seis_set_subtype" default:""`
	Sf_subtype               *string `json:"sf_subtype" default:""`
	Source                   *string `json:"source" default:""`
	Spatial_description_id   *string `json:"spatial_description_id" default:""`
	Spatial_obs_no           *int    `json:"spatial_obs_no" default:""`
	Step_seq_no              *int    `json:"step_seq_no" default:""`
	Store_id                 *string `json:"store_id" default:""`
	Store_structure_obs_no   *int    `json:"store_structure_obs_no" default:""`
	Strat_name_set_id        *string `json:"strat_name_set_id" default:""`
	Strat_unit_id            *string `json:"strat_unit_id" default:""`
	String_id                *string `json:"string_id" default:""`
	Support_facility_id      *string `json:"support_facility_id" default:""`
	Test_num                 *string `json:"test_num" default:""`
	Test_source              *string `json:"test_source" default:""`
	Test_type                *string `json:"test_type" default:""`
	Thesaurus_id             *string `json:"thesaurus_id" default:""`
	Thesaurus_word_id        *string `json:"thesaurus_word_id" default:""`
	Uwi                      *string `json:"uwi" default:""`
	Well_activity_id         *string `json:"well_activity_id" default:""`
	Well_activity_set_id     *string `json:"well_activity_set_id" default:""`
	Well_activity_source     *string `json:"well_activity_source" default:""`
	Well_activity_type_id    *string `json:"well_activity_type_id" default:""`
	Well_core_desc_id        *string `json:"well_core_desc_id" default:""`
	Well_test_id             *string `json:"well_test_id" default:""`
	Well_test_recovery_id    *string `json:"well_test_recovery_id" default:""`
	Well_test_source         *string `json:"well_test_source" default:""`
	Well_test_type           *string `json:"well_test_type" default:""`
	Work_order_id            *string `json:"work_order_id" default:""`
	Row_changed_by           *string `json:"row_changed_by" default:""`
	Row_changed_date         *string `json:"row_changed_date" default:""`
	Row_created_by           *string `json:"row_created_by" default:""`
	Row_created_date         *string `json:"row_created_date" default:""`
	Row_effective_date       *string `json:"row_effective_date" default:""`
	Row_expiry_date          *string `json:"row_expiry_date" default:""`
	Row_quality              *string `json:"row_quality" default:""`
}
