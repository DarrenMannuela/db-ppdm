package dto

type Facility_version struct {
	Facility_id          string   `json:"facility_id" default:"source"`
	Facility_type        string   `json:"facility_type" default:"source"`
	Source               string   `json:"source" default:"source"`
	Active_date          *string  `json:"active_date" default:""`
	Active_ind           *string  `json:"active_ind" default:""`
	Constructed_date     *string  `json:"constructed_date" default:""`
	Current_operator     *string  `json:"current_operator" default:""`
	Current_status_date  *string  `json:"current_status_date" default:""`
	Description          *string  `json:"description" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Facility_long_name   *string  `json:"facility_long_name" default:""`
	Facility_short_name  *string  `json:"facility_short_name" default:""`
	Facility_status_id   *string  `json:"facility_status_id" default:""`
	Inactive_date        *string  `json:"inactive_date" default:""`
	Last_injection_date  *string  `json:"last_injection_date" default:""`
	Last_production_date *string  `json:"last_production_date" default:""`
	Last_reported_date   *string  `json:"last_reported_date" default:""`
	Latitude             *float64 `json:"latitude" default:""`
	Longitude            *float64 `json:"longitude" default:""`
	On_injection_date    *string  `json:"on_injection_date" default:""`
	On_production_date   *string  `json:"on_production_date" default:""`
	Plot_name            *string  `json:"plot_name" default:""`
	Plot_symbol          *string  `json:"plot_symbol" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Remark               *string  `json:"remark" default:""`
	Status_type          *string  `json:"status_type" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
