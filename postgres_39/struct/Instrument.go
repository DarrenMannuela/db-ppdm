package dto

import (
	"time"
)

type Instrument struct {
	Instrument_id           string     `json:"instrument_id"`
	Active_ind              *string    `json:"active_ind"`
	Add_for_service_ba_id   *string    `json:"add_for_service_ba_id"`
	Area_id                 *string    `json:"area_id"`
	Area_type               *string    `json:"area_type"`
	Book_name               *string    `json:"book_name"`
	Book_number             *string    `json:"book_number"`
	Caveator_ba_id          *string    `json:"caveator_ba_id"`
	Completion_report_ind   *string    `json:"completion_report_ind"`
	Description             *string    `json:"description"`
	Discharge_date          *time.Time `json:"discharge_date"`
	Discharge_ind           *string    `json:"discharge_ind"`
	Discharge_num           *string    `json:"discharge_num"`
	Document_num            *string    `json:"document_num"`
	Drilling_intent_ind     *string    `json:"drilling_intent_ind"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Instrument_type         *string    `json:"instrument_type"`
	Jurisdiction            *string    `json:"jurisdiction"`
	Jurisdiction_add_obs_no *int       `json:"jurisdiction_add_obs_no"`
	Jurisdiction_add_source *string    `json:"jurisdiction_add_source"`
	Line_number             *string    `json:"line_number"`
	Page_number             *string    `json:"page_number"`
	Plug_and_abandon_ind    *string    `json:"plug_and_abandon_ind"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Production_lease_ind    *string    `json:"production_lease_ind"`
	Received_date           *time.Time `json:"received_date"`
	Recorded_date           *time.Time `json:"recorded_date"`
	Registered_by_ba_id     *string    `json:"registered_by_ba_id"`
	Registration_date       *time.Time `json:"registration_date"`
	Registration_num        *string    `json:"registration_num"`
	Remark                  *string    `json:"remark"`
	Source                  *string    `json:"source"`
	Valid_lease_ind         *string    `json:"valid_lease_ind"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
