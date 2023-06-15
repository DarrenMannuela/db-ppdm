package dto

import (
	"time"
)

type Facility_version struct {
	Facility_id          string     `json:"facility_id"`
	Facility_type        string     `json:"facility_type"`
	Source               string     `json:"source"`
	Active_date          *time.Time `json:"active_date"`
	Active_ind           *string    `json:"active_ind"`
	Constructed_date     *time.Time `json:"constructed_date"`
	Current_operator     *string    `json:"current_operator"`
	Current_status_date  *time.Time `json:"current_status_date"`
	Description          *string    `json:"description"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Facility_long_name   *string    `json:"facility_long_name"`
	Facility_short_name  *string    `json:"facility_short_name"`
	Facility_status_id   *string    `json:"facility_status_id"`
	Inactive_date        *time.Time `json:"inactive_date"`
	Last_injection_date  *time.Time `json:"last_injection_date"`
	Last_production_date *time.Time `json:"last_production_date"`
	Last_reported_date   *time.Time `json:"last_reported_date"`
	Latitude             *float64   `json:"latitude"`
	Longitude            *float64   `json:"longitude"`
	On_injection_date    *time.Time `json:"on_injection_date"`
	On_production_date   *time.Time `json:"on_production_date"`
	Plot_name            *string    `json:"plot_name"`
	Plot_symbol          *string    `json:"plot_symbol"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Status_type          *string    `json:"status_type"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
