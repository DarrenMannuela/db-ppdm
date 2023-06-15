package dto

import (
	"time"
)

type Sample_component struct {
	Sample_id                string     `json:"sample_id"`
	Component_obs_no         int        `json:"component_obs_no"`
	Active_ind               *string    `json:"active_ind"`
	Activity_obs_no          *int       `json:"activity_obs_no"`
	Analysis_id              *string    `json:"analysis_id"`
	Anl_source               *string    `json:"anl_source"`
	Application_id           *string    `json:"application_id"`
	Area_id                  *string    `json:"area_id"`
	Area_type                *string    `json:"area_type"`
	Authority_id             *string    `json:"authority_id"`
	Ba_organization_id       *string    `json:"ba_organization_id"`
	Ba_organization_seq_no   *int       `json:"ba_organization_seq_no"`
	Business_associate_id    *string    `json:"business_associate_id"`
	Catalogue_additive_id    *string    `json:"catalogue_additive_id"`
	Catalogue_equip_id       *string    `json:"catalogue_equip_id"`
	Classification_system_id *string    `json:"classification_system_id"`
	Class_level_id           *string    `json:"class_level_id"`
	Component_type           *string    `json:"component_type"`
	Consent_id               *string    `json:"consent_id"`
	Consult_id               *string    `json:"consult_id"`
	Contest_id               *string    `json:"contest_id"`
	Contract_id              *string    `json:"contract_id"`
	Ecozone_id               *string    `json:"ecozone_id"`
	Effective_date           *time.Time `json:"effective_date"`
	Employee_ba_id           *string    `json:"employee_ba_id"`
	Employee_obs_no          *int       `json:"employee_obs_no"`
	Employer_ba_id           *string    `json:"employer_ba_id"`
	Entitlement_id           *string    `json:"entitlement_id"`
	Equipment_id             *string    `json:"equipment_id"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Facility_id              *string    `json:"facility_id"`
	Facility_type            *string    `json:"facility_type"`
	Field_id                 *string    `json:"field_id"`
	Field_station_id         *string    `json:"field_station_id"`
	Finance_id               *string    `json:"finance_id"`
	Fossil_id                *string    `json:"fossil_id"`
	Incident_id              *string    `json:"incident_id"`
	Incident_set_id          *string    `json:"incident_set_id"`
	Incident_type            *string    `json:"incident_type"`
	Information_item_id      *string    `json:"information_item_id"`
	Info_item_subtype        *string    `json:"info_item_subtype"`
	Instrument_id            *string    `json:"instrument_id"`
	Interest_set_id          *string    `json:"interest_set_id"`
	Interest_set_seq_no      *int       `json:"interest_set_seq_no"`
	Jurisdiction             *string    `json:"jurisdiction"`
	Land_right_id            *string    `json:"land_right_id"`
	Land_right_subtype       *string    `json:"land_right_subtype"`
	Land_sale_number         *string    `json:"land_sale_number"`
	Language                 *string    `json:"language"`
	Lithology_log_id         *string    `json:"lithology_log_id"`
	Lith_log_source          *string    `json:"lith_log_source"`
	Notification_id          *string    `json:"notification_id"`
	Obligation_id            *string    `json:"obligation_id"`
	Obligation_seq_no        *int       `json:"obligation_seq_no"`
	Paleo_summary_id         *string    `json:"paleo_summary_id"`
	Pden_id                  *string    `json:"pden_id"`
	Pden_source              *string    `json:"pden_source"`
	Pden_subtype             *string    `json:"pden_subtype"`
	Physical_item_id         *string    `json:"physical_item_id"`
	Pool_id                  *string    `json:"pool_id"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Ppdm_system_id           *string    `json:"ppdm_system_id"`
	Ppdm_table_name          *string    `json:"ppdm_table_name"`
	Product_type             *string    `json:"product_type"`
	Project_id               *string    `json:"project_id"`
	Pr_str_source            *string    `json:"pr_str_source"`
	Pr_str_uwi               *string    `json:"pr_str_uwi"`
	Rate_schedule_id         *string    `json:"rate_schedule_id"`
	Referenced_guid          *string    `json:"referenced_guid"`
	Referenced_system_id     *string    `json:"referenced_system_id"`
	Referenced_table_name    *string    `json:"referenced_table_name"`
	Remark                   *string    `json:"remark"`
	Resent_id                *string    `json:"resent_id"`
	Reserve_class_id         *string    `json:"reserve_class_id"`
	Seis_set_id              *string    `json:"seis_set_id"`
	Seis_set_subtype         *string    `json:"seis_set_subtype"`
	Sf_subtype               *string    `json:"sf_subtype"`
	Source                   *string    `json:"source"`
	Spatial_description_id   *string    `json:"spatial_description_id"`
	Spatial_obs_no           *int       `json:"spatial_obs_no"`
	Store_id                 *string    `json:"store_id"`
	Store_structure_obs_no   *int       `json:"store_structure_obs_no"`
	Strat_name_set_id        *string    `json:"strat_name_set_id"`
	Strat_unit_id            *string    `json:"strat_unit_id"`
	String_id                *string    `json:"string_id"`
	Support_facility_id      *string    `json:"support_facility_id"`
	Sw_application_id        *string    `json:"sw_application_id"`
	Thesaurus_id             *string    `json:"thesaurus_id"`
	Thesaurus_word_id        *string    `json:"thesaurus_word_id"`
	Uwi                      *string    `json:"uwi"`
	Well_activity_set_id     *string    `json:"well_activity_set_id"`
	Well_activity_source     *string    `json:"well_activity_source"`
	Well_activity_type_id    *string    `json:"well_activity_type_id"`
	Well_activity_uwi        *string    `json:"well_activity_uwi"`
	Well_core_id             *string    `json:"well_core_id"`
	Well_core_source         *string    `json:"well_core_source"`
	Well_core_uwi            *string    `json:"well_core_uwi"`
	Well_set_id              *string    `json:"well_set_id"`
	Work_order_id            *string    `json:"work_order_id"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
