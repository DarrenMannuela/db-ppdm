package dto

import (
	"time"
)

type Sf_component struct {
	Sf_subtype             string     `json:"sf_subtype"`
	Support_facility_id    string     `json:"support_facility_id"`
	Use_id                 string     `json:"use_id"`
	Component_obs_no       int        `json:"component_obs_no"`
	Active_ind             *string    `json:"active_ind"`
	Area_id                *string    `json:"area_id"`
	Area_type              *string    `json:"area_type"`
	Ba_licensee_ba_id      *string    `json:"ba_licensee_ba_id"`
	Ba_license_id          *string    `json:"ba_license_id"`
	Business_associate_id  *string    `json:"business_associate_id"`
	Consent_id             *string    `json:"consent_id"`
	Consult_id             *string    `json:"consult_id"`
	Contest_id             *string    `json:"contest_id"`
	Contract_id            *string    `json:"contract_id"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Facility_id            *string    `json:"facility_id"`
	Facility_license_id    *string    `json:"facility_license_id"`
	Facility_type          *string    `json:"facility_type"`
	Field_id               *string    `json:"field_id"`
	Field_station_id       *string    `json:"field_station_id"`
	Finance_id             *string    `json:"finance_id"`
	Hse_incident_id        *string    `json:"hse_incident_id"`
	Information_item_id    *string    `json:"information_item_id"`
	Info_item_subtype      *string    `json:"info_item_subtype"`
	Instrument_id          *string    `json:"instrument_id"`
	Interest_set_id        *string    `json:"interest_set_id"`
	Interest_set_seq_no    *int       `json:"interest_set_seq_no"`
	Land_right_id          *string    `json:"land_right_id"`
	Land_right_subtype     *string    `json:"land_right_subtype"`
	Land_sale_jurisdiction *string    `json:"land_sale_jurisdiction"`
	Land_sale_number       *string    `json:"land_sale_number"`
	Obligation_id          *string    `json:"obligation_id"`
	Obligation_seq_no      *int       `json:"obligation_seq_no"`
	Pden_id                *string    `json:"pden_id"`
	Pden_source            *string    `json:"pden_source"`
	Pden_subtype           *string    `json:"pden_subtype"`
	Pool_id                *string    `json:"pool_id"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Ppdm_system_id         *string    `json:"ppdm_system_id"`
	Ppdm_table_name        *string    `json:"ppdm_table_name"`
	Prod_string_id         *string    `json:"prod_string_id"`
	Prod_string_source     *string    `json:"prod_string_source"`
	Project_id             *string    `json:"project_id"`
	Pr_str_form_obs_no     *int       `json:"pr_str_form_obs_no"`
	Rate_schedule_id       *string    `json:"rate_schedule_id"`
	Remark                 *string    `json:"remark"`
	Seis_license_id        *string    `json:"seis_license_id"`
	Seis_set_id            *string    `json:"seis_set_id"`
	Seis_set_subtype       *string    `json:"seis_set_subtype"`
	Sf_component_type      *string    `json:"sf_component_type"`
	Source                 *string    `json:"source"`
	Source_document_id     *string    `json:"source_document_id"`
	Sw_application_id      *string    `json:"sw_application_id"`
	Uwi                    *string    `json:"uwi"`
	Well_license_id        *string    `json:"well_license_id"`
	Well_license_source    *string    `json:"well_license_source"`
	Well_set_id            *string    `json:"well_set_id"`
	Work_order_id          *string    `json:"work_order_id"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
