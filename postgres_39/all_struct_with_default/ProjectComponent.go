package dto

type Project_component struct {
	Project_id               string  `json:"project_id" default:"source"`
	Component_obs_no         int     `json:"component_obs_no" default:"1"`
	Active_ind               *string `json:"active_ind" default:""`
	Activity_obs_no          *int    `json:"activity_obs_no" default:""`
	Analysis_id              *string `json:"analysis_id" default:""`
	Application_id           *string `json:"application_id" default:""`
	Area_id                  *string `json:"area_id" default:""`
	Area_type                *string `json:"area_type" default:""`
	Authority_id             *string `json:"authority_id" default:""`
	Ba_licensee_ba_id        *string `json:"ba_licensee_ba_id" default:""`
	Ba_license_id            *string `json:"ba_license_id" default:""`
	Ba_organization_id       *string `json:"ba_organization_id" default:""`
	Ba_organization_seq_no   *int    `json:"ba_organization_seq_no" default:""`
	Business_associate_id    *string `json:"business_associate_id" default:""`
	Catalogue_additive_id    *string `json:"catalogue_additive_id" default:""`
	Catalogue_equip_id       *string `json:"catalogue_equip_id" default:""`
	Classification_system_id *string `json:"classification_system_id" default:""`
	Class_level_id           *string `json:"class_level_id" default:""`
	Component_reason         *string `json:"component_reason" default:""`
	Component_type           *string `json:"component_type" default:""`
	Condition_id             *string `json:"condition_id" default:""`
	Consent_id               *string `json:"consent_id" default:""`
	Consult_id               *string `json:"consult_id" default:""`
	Contest_id               *string `json:"contest_id" default:""`
	Contract_id              *string `json:"contract_id" default:""`
	Description              *string `json:"description" default:""`
	Ecozone_id               *string `json:"ecozone_id" default:""`
	Effective_date           *string `json:"effective_date" default:""`
	Employee_ba_id           *string `json:"employee_ba_id" default:""`
	Employee_obs_no          *int    `json:"employee_obs_no" default:""`
	Employer_ba_id           *string `json:"employer_ba_id" default:""`
	Entitlement_id           *string `json:"entitlement_id" default:""`
	Equipment_id             *string `json:"equipment_id" default:""`
	Expiry_date              *string `json:"expiry_date" default:""`
	Facility_id              *string `json:"facility_id" default:""`
	Facility_license_id      *string `json:"facility_license_id" default:""`
	Facility_type            *string `json:"facility_type" default:""`
	Field_id                 *string `json:"field_id" default:""`
	Field_source             *string `json:"field_source" default:""`
	Field_station_id         *string `json:"field_station_id" default:""`
	Field_strat_interp_id    *string `json:"field_strat_interp_id" default:""`
	Field_strat_unit_id      *string `json:"field_strat_unit_id" default:""`
	Finance_id               *string `json:"finance_id" default:""`
	Fossil_id                *string `json:"fossil_id" default:""`
	Hse_incident_id          *string `json:"hse_incident_id" default:""`
	Incident_set_id          *string `json:"incident_set_id" default:""`
	Incident_type            *string `json:"incident_type" default:""`
	Information_item_id      *string `json:"information_item_id" default:""`
	Info_item_subtype        *string `json:"info_item_subtype" default:""`
	Input_ind                *string `json:"input_ind" default:""`
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
	Lith_sample_id           *string `json:"lith_sample_id" default:""`
	Notification_id          *string `json:"notification_id" default:""`
	Obligation_id            *string `json:"obligation_id" default:""`
	Obligation_seq_no        *int    `json:"obligation_seq_no" default:""`
	Other_project_id         *string `json:"other_project_id" default:""`
	Output_ind               *string `json:"output_ind" default:""`
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
	Prod_string_id           *string `json:"prod_string_id" default:""`
	Prod_string_source       *string `json:"prod_string_source" default:""`
	Project_plan_id          *string `json:"project_plan_id" default:""`
	Project_step_id          *string `json:"project_step_id" default:""`
	Provision_id             *string `json:"provision_id" default:""`
	Rate_schedule_id         *string `json:"rate_schedule_id" default:""`
	Referenced_guid          *string `json:"referenced_guid" default:""`
	Referenced_system_id     *string `json:"referenced_system_id" default:""`
	Referenced_table_name    *string `json:"referenced_table_name" default:""`
	Remark                   *string `json:"remark" default:""`
	Resent_id                *string `json:"resent_id" default:""`
	Reserve_class_id         *string `json:"reserve_class_id" default:""`
	Reserve_product_type     *string `json:"reserve_product_type" default:""`
	Reserve_revision_id      *string `json:"reserve_revision_id" default:""`
	Reserve_revision_obs_no  *int    `json:"reserve_revision_obs_no" default:""`
	Reserve_summary_id       *string `json:"reserve_summary_id" default:""`
	Reserve_summary_obs_no   *int    `json:"reserve_summary_obs_no" default:""`
	Sample_anal_source       *string `json:"sample_anal_source" default:""`
	Seis_license_id          *string `json:"seis_license_id" default:""`
	Seis_proc_plan_id        *string `json:"seis_proc_plan_id" default:""`
	Seis_set_id              *string `json:"seis_set_id" default:""`
	Seis_set_subtype         *string `json:"seis_set_subtype" default:""`
	Sf_subtype               *string `json:"sf_subtype" default:""`
	Source                   *string `json:"source" default:""`
	Spatial_description_id   *string `json:"spatial_description_id" default:""`
	Spatial_obs_no           *int    `json:"spatial_obs_no" default:""`
	Store_id                 *string `json:"store_id" default:""`
	Store_structure_obs_no   *int    `json:"store_structure_obs_no" default:""`
	Strat_column_id          *string `json:"strat_column_id" default:""`
	Strat_column_source      *string `json:"strat_column_source" default:""`
	Strat_name_set_id        *string `json:"strat_name_set_id" default:""`
	Strat_unit_id            *string `json:"strat_unit_id" default:""`
	Support_facility_id      *string `json:"support_facility_id" default:""`
	Sw_application_id        *string `json:"sw_application_id" default:""`
	Thesaurus_id             *string `json:"thesaurus_id" default:""`
	Thesaurus_word_id        *string `json:"thesaurus_word_id" default:""`
	Uwi                      *string `json:"uwi" default:""`
	Well_activity_set_id     *string `json:"well_activity_set_id" default:""`
	Well_activity_source     *string `json:"well_activity_source" default:""`
	Well_activity_type_id    *string `json:"well_activity_type_id" default:""`
	Well_int_interp_id       *string `json:"well_int_interp_id" default:""`
	Well_int_source          *string `json:"well_int_source" default:""`
	Well_int_strat_unit_id   *string `json:"well_int_strat_unit_id" default:""`
	Well_int_uwi             *string `json:"well_int_uwi" default:""`
	Well_license_id          *string `json:"well_license_id" default:""`
	Well_license_source      *string `json:"well_license_source" default:""`
	Well_set_id              *string `json:"well_set_id" default:""`
	Work_order_id            *string `json:"work_order_id" default:""`
	Row_changed_by           *string `json:"row_changed_by" default:""`
	Row_changed_date         *string `json:"row_changed_date" default:""`
	Row_created_by           *string `json:"row_created_by" default:""`
	Row_created_date         *string `json:"row_created_date" default:""`
	Row_effective_date       *string `json:"row_effective_date" default:""`
	Row_expiry_date          *string `json:"row_expiry_date" default:""`
	Row_quality              *string `json:"row_quality" default:""`
}
