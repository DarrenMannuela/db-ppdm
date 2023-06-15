package dto

type Well_pressure_aof_4pt struct {
	Uwi                        string   `json:"uwi" default:"source"`
	Source                     string   `json:"source" default:"source"`
	Pressure_obs_no            int      `json:"pressure_obs_no" default:"1"`
	Aof_obs_no                 int      `json:"aof_obs_no" default:"1"`
	Point_obs_no               int      `json:"point_obs_no" default:"1"`
	Active_ind                 *string  `json:"active_ind" default:""`
	Bhfp                       *float64 `json:"bhfp" default:""`
	Bhfp_ouom                  *string  `json:"bhfp_ouom" default:""`
	Choke_size_desc            *string  `json:"choke_size_desc" default:""`
	Condensate_flow_rate       *float64 `json:"condensate_flow_rate" default:""`
	Condensate_flow_rate_ouom  *string  `json:"condensate_flow_rate_ouom" default:""`
	Effective_date             *string  `json:"effective_date" default:""`
	Expiry_date                *string  `json:"expiry_date" default:""`
	Final_shutin_pressure      *float64 `json:"final_shutin_pressure" default:""`
	Final_shutin_pressure_ouom *string  `json:"final_shutin_pressure_ouom" default:""`
	Flow_duration              *float64 `json:"flow_duration" default:""`
	Flow_duration_ouom         *string  `json:"flow_duration_ouom" default:""`
	Flow_pressure              *float64 `json:"flow_pressure" default:""`
	Flow_pressure_ouom         *string  `json:"flow_pressure_ouom" default:""`
	Flow_rate                  *float64 `json:"flow_rate" default:""`
	Flow_rate_ouom             *string  `json:"flow_rate_ouom" default:""`
	Ppdm_guid                  *string  `json:"ppdm_guid" default:""`
	Remark                     *string  `json:"remark" default:""`
	Water_flow_rate            *float64 `json:"water_flow_rate" default:""`
	Water_flow_rate_ouom       *string  `json:"water_flow_rate_ouom" default:""`
	Row_changed_by             *string  `json:"row_changed_by" default:""`
	Row_changed_date           *string  `json:"row_changed_date" default:""`
	Row_created_by             *string  `json:"row_created_by" default:""`
	Row_created_date           *string  `json:"row_created_date" default:""`
	Row_effective_date         *string  `json:"row_effective_date" default:""`
	Row_expiry_date            *string  `json:"row_expiry_date" default:""`
	Row_quality                *string  `json:"row_quality" default:""`
}
