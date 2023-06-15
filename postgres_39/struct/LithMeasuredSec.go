package dto

import (
	"time"
)

type Lith_measured_sec struct {
	Meas_section_id            string     `json:"meas_section_id"`
	Source                     string     `json:"source"`
	Active_ind                 *string    `json:"active_ind"`
	Author                     *string    `json:"author"`
	Description                *string    `json:"description"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Location_desc              *string    `json:"location_desc"`
	Location_qualifier         *string    `json:"location_qualifier"`
	Node_id                    *string    `json:"node_id"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Publication_reference_text *string    `json:"publication_reference_text"`
	Reference_date             *time.Time `json:"reference_date"`
	Remark                     *string    `json:"remark"`
	Source_document_id         *string    `json:"source_document_id"`
	Strat_name_set_id          *string    `json:"strat_name_set_id"`
	Strat_unit_id              *string    `json:"strat_unit_id"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
