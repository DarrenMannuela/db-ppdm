package dto

type Seis_inspection struct {
	Inspection_id           string   `json:"inspection_id" default:"source"`
	Active_ind              *string  `json:"active_ind" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Inspected_length        *float64 `json:"inspected_length" default:""`
	Inspected_length_ouom   *string  `json:"inspected_length_ouom" default:""`
	Inspecting_ba_id        *string  `json:"inspecting_ba_id" default:""`
	Inspection_date         *string  `json:"inspection_date" default:""`
	Inspection_status       *string  `json:"inspection_status" default:""`
	Insp_loc_address_obs_no *int     `json:"insp_loc_address_obs_no" default:""`
	Insp_loc_ba_id          *string  `json:"insp_loc_ba_id" default:""`
	Insp_loc_source         *string  `json:"insp_loc_source" default:""`
	Min_length              *float64 `json:"min_length" default:""`
	Min_length_ouom         *string  `json:"min_length_ouom" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Reason_type             *string  `json:"reason_type" default:""`
	Remark                  *string  `json:"remark" default:""`
	Scheduled_date          *string  `json:"scheduled_date" default:""`
	Source                  *string  `json:"source" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
