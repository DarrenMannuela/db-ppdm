package dto

type Well_test_comput_anal struct {
	Uwi                         string   `json:"uwi" default:"source"`
	Source                      string   `json:"source" default:"source"`
	Test_type                   string   `json:"test_type" default:"source"`
	Run_num                     string   `json:"run_num" default:"source"`
	Test_num                    string   `json:"test_num" default:"source"`
	Report_no                   int      `json:"report_no" default:"1"`
	Active_ind                  *string  `json:"active_ind" default:""`
	Analysis_company            *string  `json:"analysis_company" default:""`
	Computed_permeability       *float64 `json:"computed_permeability" default:""`
	Computed_permeability_ouom  *string  `json:"computed_permeability_ouom" default:""`
	Confidence_limit            *float64 `json:"confidence_limit" default:""`
	Confidence_limit_ouom       *string  `json:"confidence_limit_ouom" default:""`
	Effective_date              *string  `json:"effective_date" default:""`
	Est_damage_ratio            *float64 `json:"est_damage_ratio" default:""`
	Expiry_date                 *string  `json:"expiry_date" default:""`
	Extrap_pressure_percent     *float64 `json:"extrap_pressure_percent" default:""`
	Final_reservoir_pressure    *float64 `json:"final_reservoir_pressure" default:""`
	Final_reservoir_press_ouom  *string  `json:"final_reservoir_press_ouom" default:""`
	Gauge_depth                 *float64 `json:"gauge_depth" default:""`
	Gauge_depth_ouom            *string  `json:"gauge_depth_ouom" default:""`
	Investigation_radius        *float64 `json:"investigation_radius" default:""`
	Investigation_radius_ouom   *string  `json:"investigation_radius_ouom" default:""`
	Max_reservoir_pressure      *float64 `json:"max_reservoir_pressure" default:""`
	Max_reservoir_pressure_ouom *string  `json:"max_reservoir_pressure_ouom" default:""`
	Potmtrc_surf_height         *float64 `json:"potmtrc_surf_height" default:""`
	Potmtrc_surf_height_ouom    *string  `json:"potmtrc_surf_height_ouom" default:""`
	Ppdm_guid                   *string  `json:"ppdm_guid" default:""`
	Production_index_rate       *float64 `json:"production_index_rate" default:""`
	Production_index_rate_ouom  *string  `json:"production_index_rate_ouom" default:""`
	Recorder_id                 *string  `json:"recorder_id" default:""`
	Remark                      *string  `json:"remark" default:""`
	Row_changed_by              *string  `json:"row_changed_by" default:""`
	Row_changed_date            *string  `json:"row_changed_date" default:""`
	Row_created_by              *string  `json:"row_created_by" default:""`
	Row_created_date            *string  `json:"row_created_date" default:""`
	Row_effective_date          *string  `json:"row_effective_date" default:""`
	Row_expiry_date             *string  `json:"row_expiry_date" default:""`
	Row_quality                 *string  `json:"row_quality" default:""`
}
