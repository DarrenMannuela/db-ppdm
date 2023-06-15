package dto

type Well_drill_mud_additive struct {
	Uwi                  string   `json:"uwi" default:"source"`
	Drill_period_obs_no  int      `json:"drill_period_obs_no" default:"1"`
	Additive_id          string   `json:"additive_id" default:"source"`
	Additive_seq_no      int      `json:"additive_seq_no" default:"1"`
	Active_ind           *string  `json:"active_ind" default:""`
	Additive_method      *string  `json:"additive_method" default:""`
	Additive_period      *float64 `json:"additive_period" default:""`
	Additive_period_ouom *string  `json:"additive_period_ouom" default:""`
	Additive_period_uom  *string  `json:"additive_period_uom" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	End_time             *string  `json:"end_time" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Quantity_count       *float64 `json:"quantity_count" default:""`
	Quantity_value       *float64 `json:"quantity_value" default:""`
	Quantity_value_ouom  *string  `json:"quantity_value_ouom" default:""`
	Quantity_value_uom   *string  `json:"quantity_value_uom" default:""`
	Remark               *string  `json:"remark" default:""`
	Source               *string  `json:"source" default:""`
	Start_time           *string  `json:"start_time" default:""`
	Timezone             *string  `json:"timezone" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
