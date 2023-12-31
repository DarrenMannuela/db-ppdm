package dto

type Well_pressure_aof struct {
	Uwi                         string   `json:"uwi" default:"source"`
	Source                      string   `json:"source" default:"source"`
	Pressure_obs_no             int      `json:"pressure_obs_no" default:"1"`
	Aof_obs_no                  int      `json:"aof_obs_no" default:"1"`
	Active_ind                  *string  `json:"active_ind" default:""`
	Analysis_type               *string  `json:"analysis_type" default:""`
	Aof_calculate_method        *string  `json:"aof_calculate_method" default:""`
	Aof_potential               *float64 `json:"aof_potential" default:""`
	Aof_potential_ouom          *string  `json:"aof_potential_ouom" default:""`
	Aof_slope                   *float64 `json:"aof_slope" default:""`
	Bottom_hole_pressure_method *string  `json:"bottom_hole_pressure_method" default:""`
	Caof_rate                   *float64 `json:"caof_rate" default:""`
	Caof_rate_ouom              *string  `json:"caof_rate_ouom" default:""`
	Choke_size_desc             *string  `json:"choke_size_desc" default:""`
	Effective_date              *string  `json:"effective_date" default:""`
	Expiry_date                 *string  `json:"expiry_date" default:""`
	Flow_period_duration        *float64 `json:"flow_period_duration" default:""`
	Flow_period_duration_ouom   *string  `json:"flow_period_duration_ouom" default:""`
	Flow_pressure               *float64 `json:"flow_pressure" default:""`
	Flow_pressure_ouom          *string  `json:"flow_pressure_ouom" default:""`
	Flow_rate                   *float64 `json:"flow_rate" default:""`
	Flow_rate_ouom              *string  `json:"flow_rate_ouom" default:""`
	Four_point_caof_rate        *float64 `json:"four_point_caof_rate" default:""`
	Four_point_caof_rate_ouom   *string  `json:"four_point_caof_rate_ouom" default:""`
	Ppdm_guid                   *string  `json:"ppdm_guid" default:""`
	Remark                      *string  `json:"remark" default:""`
	Reservoir_pressure          *float64 `json:"reservoir_pressure" default:""`
	Reservoir_pressure_ouom     *string  `json:"reservoir_pressure_ouom" default:""`
	Shutin_pressure_type        *string  `json:"shutin_pressure_type" default:""`
	Test_date                   *string  `json:"test_date" default:""`
	Row_changed_by              *string  `json:"row_changed_by" default:""`
	Row_changed_date            *string  `json:"row_changed_date" default:""`
	Row_created_by              *string  `json:"row_created_by" default:""`
	Row_created_date            *string  `json:"row_created_date" default:""`
	Row_effective_date          *string  `json:"row_effective_date" default:""`
	Row_expiry_date             *string  `json:"row_expiry_date" default:""`
	Row_quality                 *string  `json:"row_quality" default:""`
}
