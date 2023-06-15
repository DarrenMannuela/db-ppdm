package dto

type Seis_velocity_interval struct {
	Interval_id          string   `json:"interval_id" default:"source"`
	Active_ind           *string  `json:"active_ind" default:""`
	Base_depth           *float64 `json:"base_depth" default:""`
	Base_depth_ouom      *string  `json:"base_depth_ouom" default:""`
	Base_strat_unit_id   *string  `json:"base_strat_unit_id" default:""`
	Compute_method       *string  `json:"compute_method" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Last_update          *string  `json:"last_update" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Remark               *string  `json:"remark" default:""`
	Reported_tvd         *float64 `json:"reported_tvd" default:""`
	Reported_tvd_ouom    *string  `json:"reported_tvd_ouom" default:""`
	Seis_set_id          *string  `json:"seis_set_id" default:""`
	Seis_set_subtype     *string  `json:"seis_set_subtype" default:""`
	Seis_time_datum      *float64 `json:"seis_time_datum" default:""`
	Seis_time_datum_ouom *string  `json:"seis_time_datum_ouom" default:""`
	Source               *string  `json:"source" default:""`
	Strat_name_set_id    *string  `json:"strat_name_set_id" default:""`
	Top_depth            *float64 `json:"top_depth" default:""`
	Top_depth_ouom       *string  `json:"top_depth_ouom" default:""`
	Top_strat_unit_id    *string  `json:"top_strat_unit_id" default:""`
	Uwi                  *string  `json:"uwi" default:""`
	Velocity_quality     *string  `json:"velocity_quality" default:""`
	Velocity_type        *string  `json:"velocity_type" default:""`
	Velocity_value       *float64 `json:"velocity_value" default:""`
	Velocity_value_ouom  *string  `json:"velocity_value_ouom" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
