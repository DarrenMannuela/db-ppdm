package dto

import (
	"time"
)

type Seis_inspection struct {
	Inspection_id           string     `json:"inspection_id"`
	Active_ind              *string    `json:"active_ind"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Inspected_length        *float64   `json:"inspected_length"`
	Inspected_length_ouom   *string    `json:"inspected_length_ouom"`
	Inspecting_ba_id        *string    `json:"inspecting_ba_id"`
	Inspection_date         *time.Time `json:"inspection_date"`
	Inspection_status       *string    `json:"inspection_status"`
	Insp_loc_address_obs_no *int       `json:"insp_loc_address_obs_no"`
	Insp_loc_ba_id          *string    `json:"insp_loc_ba_id"`
	Insp_loc_source         *string    `json:"insp_loc_source"`
	Min_length              *float64   `json:"min_length"`
	Min_length_ouom         *string    `json:"min_length_ouom"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Reason_type             *string    `json:"reason_type"`
	Remark                  *string    `json:"remark"`
	Scheduled_date          *time.Time `json:"scheduled_date"`
	Source                  *string    `json:"source"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
