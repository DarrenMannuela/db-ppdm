package dto

import (
	"time"
)

type Seis_segment struct {
	Seis_set_subtype      string     `json:"seis_set_subtype"`
	Segment_id            string     `json:"segment_id"`
	Acqtn_design_id       *string    `json:"acqtn_design_id"`
	Active_ind            *string    `json:"active_ind"`
	Area_id               *string    `json:"area_id"`
	Area_type             *string    `json:"area_type"`
	Business_associate_id *string    `json:"business_associate_id"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Refraction_reflection *string    `json:"refraction_reflection"`
	Remark                *string    `json:"remark"`
	Row_audit_by          *string    `json:"row_audit_by"`
	Row_audit_date        *time.Time `json:"row_audit_date"`
	Seis_dimension        *string    `json:"seis_dimension"`
	Seis_line_id          *string    `json:"seis_line_id"`
	Seis_line_set_subtype *string    `json:"seis_line_set_subtype"`
	Seis_segment_reason   *string    `json:"seis_segment_reason"`
	Seis_spectrum_type    *string    `json:"seis_spectrum_type"`
	Seis_station_type     *string    `json:"seis_station_type"`
	Source                *string    `json:"source"`
	Test_experimental     *string    `json:"test_experimental"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
