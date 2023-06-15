package dto

import (
	"time"
)

type Obligation_component struct {
	Obligation_id             string     `json:"obligation_id"`
	Obligation_seq_no         int        `json:"obligation_seq_no"`
	Component_obs_no          int        `json:"component_obs_no"`
	Active_ind                *string    `json:"active_ind"`
	Activity_obs_no           *int       `json:"activity_obs_no"`
	Analysis_id               *string    `json:"analysis_id"`
	Application_id            *string    `json:"application_id"`
	Area_id                   *string    `json:"area_id"`
	Area_type                 *string    `json:"area_type"`
	Authority_id              *string    `json:"authority_id"`
	Ba_licensee_ba_id         *string    `json:"ba_licensee_ba_id"`
	Ba_license_id             *string    `json:"ba_license_id"`
	Ba_organization_id        *string    `json:"ba_organization_id"`
	Ba_organization_seq_no    *int       `json:"ba_organization_seq_no"`
	Ba_service_ba_id          *string    `json:"ba_service_ba_id"`
	Ba_service_provider       *string    `json:"ba_service_provider"`
	Ba_service_seq_no         *int       `json:"ba_service_seq_no"`
	Business_associate_id     *string    `json:"business_associate_id"`
	Calculation_method        *string    `json:"calculation_method"`
	Catalogue_additive_id     *string    `json:"catalogue_additive_id"`
	Catalogue_equip_id        *string    `json:"catalogue_equip_id"`
	Classification_system_id  *string    `json:"classification_system_id"`
	Class_level_id            *string    `json:"class_level_id"`
	Component_reason          *string    `json:"component_reason"`
	Consent_id                *string    `json:"consent_id"`
	Consult_id                *string    `json:"consult_id"`
	Contest_id                *string    `json:"contest_id"`
	Contract_extension_id     *string    `json:"contract_extension_id"`
	Contract_id               *string    `json:"contract_id"`
	Contract_provision_id     *string    `json:"contract_provision_id"`
	Cont_service_ba_id        *string    `json:"cont_service_ba_id"`
	Cont_service_seq_no       *int       `json:"cont_service_seq_no"`
	Critical_date             *time.Time `json:"critical_date"`
	Description               *string    `json:"description"`
	Ecozone_id                *string    `json:"ecozone_id"`
	Effective_date            *time.Time `json:"effective_date"`
	Employee_ba_id            *string    `json:"employee_ba_id"`
	Employee_obs_no           *int       `json:"employee_obs_no"`
	Employer_ba_id            *string    `json:"employer_ba_id"`
	Entitlement_id            *string    `json:"entitlement_id"`
	Equipment_id              *string    `json:"equipment_id"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Facility_id               *string    `json:"facility_id"`
	Facility_license_id       *string    `json:"facility_license_id"`
	Facility_service_ba_id    *string    `json:"facility_service_ba_id"`
	Facility_service_seq_no   *int       `json:"facility_service_seq_no"`
	Facility_type             *string    `json:"facility_type"`
	Field_id                  *string    `json:"field_id"`
	Field_station_id          *string    `json:"field_station_id"`
	Finance_id                *string    `json:"finance_id"`
	Fossil_id                 *string    `json:"fossil_id"`
	Fulfilled_date            *time.Time `json:"fulfilled_date"`
	Fulfilled_ind             *string    `json:"fulfilled_ind"`
	Fulfilled_remark          *string    `json:"fulfilled_remark"`
	Hse_incident_id           *string    `json:"hse_incident_id"`
	Incident_set_id           *string    `json:"incident_set_id"`
	Incident_type             *string    `json:"incident_type"`
	Information_item_id       *string    `json:"information_item_id"`
	Info_item_subtype         *string    `json:"info_item_subtype"`
	Instrument_id             *string    `json:"instrument_id"`
	Interest_set_id           *string    `json:"interest_set_id"`
	Interest_set_seq_no       *int       `json:"interest_set_seq_no"`
	Land_cost_id              *string    `json:"land_cost_id"`
	Land_offering_bid_id      *string    `json:"land_offering_bid_id"`
	Land_request_id           *string    `json:"land_request_id"`
	Land_right_id             *string    `json:"land_right_id"`
	Land_right_subtype        *string    `json:"land_right_subtype"`
	Land_sale_bid_set_id      *string    `json:"land_sale_bid_set_id"`
	Land_sale_jurisdiction    *string    `json:"land_sale_jurisdiction"`
	Land_sale_number          *string    `json:"land_sale_number"`
	Land_sale_offering_id     *string    `json:"land_sale_offering_id"`
	Land_service_provided_by  *string    `json:"land_service_provided_by"`
	Land_service_seq_no       *int       `json:"land_service_seq_no"`
	Language                  *string    `json:"language"`
	Lithology_log_id          *string    `json:"lithology_log_id"`
	Lith_log_source           *string    `json:"lith_log_source"`
	Lr_termination_id         *string    `json:"lr_termination_id"`
	Lr_termination_seq_no     *int       `json:"lr_termination_seq_no"`
	Notification_id           *string    `json:"notification_id"`
	Obligation_component_type *string    `json:"obligation_component_type"`
	Paleo_summary_id          *string    `json:"paleo_summary_id"`
	Pden_id                   *string    `json:"pden_id"`
	Pden_source               *string    `json:"pden_source"`
	Pden_subtype              *string    `json:"pden_subtype"`
	Physical_item_id          *string    `json:"physical_item_id"`
	Pool_id                   *string    `json:"pool_id"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Ppdm_system_id            *string    `json:"ppdm_system_id"`
	Ppdm_table_name           *string    `json:"ppdm_table_name"`
	Product_type              *string    `json:"product_type"`
	Prod_string_id            *string    `json:"prod_string_id"`
	Prod_string_source        *string    `json:"prod_string_source"`
	Project_id                *string    `json:"project_id"`
	Rate_schedule_id          *string    `json:"rate_schedule_id"`
	Referenced_guid           *string    `json:"referenced_guid"`
	Referenced_system_id      *string    `json:"referenced_system_id"`
	Referenced_table_name     *string    `json:"referenced_table_name"`
	Remark                    *string    `json:"remark"`
	Resent_id                 *string    `json:"resent_id"`
	Reserve_class_id          *string    `json:"reserve_class_id"`
	Sample_anal_source        *string    `json:"sample_anal_source"`
	Seis_license_id           *string    `json:"seis_license_id"`
	Seis_point_flow_id        *string    `json:"seis_point_flow_id"`
	Seis_point_id             *string    `json:"seis_point_id"`
	Seis_service_ba_id        *string    `json:"seis_service_ba_id"`
	Seis_service_seq_no       *int       `json:"seis_service_seq_no"`
	Seis_set_id               *string    `json:"seis_set_id"`
	Seis_set_subtype          *string    `json:"seis_set_subtype"`
	Seis_transaction_id       *string    `json:"seis_transaction_id"`
	Seis_transaction_type     *string    `json:"seis_transaction_type"`
	Sf_subtype                *string    `json:"sf_subtype"`
	Source                    *string    `json:"source"`
	Spatial_description_id    *string    `json:"spatial_description_id"`
	Spatial_obs_no            *int       `json:"spatial_obs_no"`
	Store_id                  *string    `json:"store_id"`
	Store_structure_obs_no    *int       `json:"store_structure_obs_no"`
	Strat_name_set_id         *string    `json:"strat_name_set_id"`
	Strat_unit_id             *string    `json:"strat_unit_id"`
	Support_facility_id       *string    `json:"support_facility_id"`
	Sw_application_id         *string    `json:"sw_application_id"`
	Thesaurus_id              *string    `json:"thesaurus_id"`
	Thesaurus_word_id         *string    `json:"thesaurus_word_id"`
	Uwi                       *string    `json:"uwi"`
	Well_activity_set_id      *string    `json:"well_activity_set_id"`
	Well_activity_source      *string    `json:"well_activity_source"`
	Well_activity_type_id     *string    `json:"well_activity_type_id"`
	Well_license_id           *string    `json:"well_license_id"`
	Well_license_source       *string    `json:"well_license_source"`
	Well_service_ba_id        *string    `json:"well_service_ba_id"`
	Well_service_seq_no       *int       `json:"well_service_seq_no"`
	Well_set_id               *string    `json:"well_set_id"`
	Work_order_delivery_id    *string    `json:"work_order_delivery_id"`
	Work_order_id             *string    `json:"work_order_id"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
