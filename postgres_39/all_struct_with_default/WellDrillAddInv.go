package dto

type Well_drill_add_inv struct {
	Uwi                      string   `json:"uwi" default:"source"`
	Drill_period_obs_no      int      `json:"drill_period_obs_no" default:"1"`
	Additive_id              string   `json:"additive_id" default:"source"`
	Active_ind               *string  `json:"active_ind" default:""`
	Effective_date           *string  `json:"effective_date" default:""`
	Expiry_date              *string  `json:"expiry_date" default:""`
	Ppdm_guid                *string  `json:"ppdm_guid" default:""`
	Quantity_count_opening   *float64 `json:"quantity_count_opening" default:""`
	Quantity_count_ordered   *float64 `json:"quantity_count_ordered" default:""`
	Quantity_count_remaining *float64 `json:"quantity_count_remaining" default:""`
	Quantity_count_used      *float64 `json:"quantity_count_used" default:""`
	Quantity_value_opening   *float64 `json:"quantity_value_opening" default:""`
	Quantity_value_ordered   *float64 `json:"quantity_value_ordered" default:""`
	Quantity_value_ouom      *string  `json:"quantity_value_ouom" default:""`
	Quantity_value_remaining *float64 `json:"quantity_value_remaining" default:""`
	Quantity_value_uom       *string  `json:"quantity_value_uom" default:""`
	Quantity_value_used      *float64 `json:"quantity_value_used" default:""`
	Remark                   *string  `json:"remark" default:""`
	Report_time_type         *string  `json:"report_time_type" default:""`
	Sf_subtype               *string  `json:"sf_subtype" default:""`
	Source                   *string  `json:"source" default:""`
	Support_facility_id      *string  `json:"support_facility_id" default:""`
	Row_changed_by           *string  `json:"row_changed_by" default:""`
	Row_changed_date         *string  `json:"row_changed_date" default:""`
	Row_created_by           *string  `json:"row_created_by" default:""`
	Row_created_date         *string  `json:"row_created_date" default:""`
	Row_effective_date       *string  `json:"row_effective_date" default:""`
	Row_expiry_date          *string  `json:"row_expiry_date" default:""`
	Row_quality              *string  `json:"row_quality" default:""`
}
