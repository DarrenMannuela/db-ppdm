package dto

type Anl_oil_analysis struct {
	Analysis_id                 string   `json:"analysis_id" default:"source"`
	Anl_source                  string   `json:"anl_source" default:"source"`
	Oil_analysis_obs_no         int      `json:"oil_analysis_obs_no" default:"1"`
	Absolute_gravity            *float64 `json:"absolute_gravity" default:""`
	Absolute_gravity_ouom       *string  `json:"absolute_gravity_ouom" default:""`
	Active_ind                  *string  `json:"active_ind" default:""`
	Asphaltene_content          *float64 `json:"asphaltene_content" default:""`
	Asphaltene_content_ouom     *string  `json:"asphaltene_content_ouom" default:""`
	Asph_content_preferred_ind  *string  `json:"asph_content_preferred_ind" default:""`
	Asph_content_problem_ind    *string  `json:"asph_content_problem_ind" default:""`
	Asph_content_step_seq_no    *int     `json:"asph_content_step_seq_no" default:""`
	Characterize_factor         *string  `json:"characterize_factor" default:""`
	Characterize_factor_ouom    *string  `json:"characterize_factor_ouom" default:""`
	Cloud_point_preferred_ind   *string  `json:"cloud_point_preferred_ind" default:""`
	Cloud_point_problem_ind     *string  `json:"cloud_point_problem_ind" default:""`
	Cloud_point_step_seq_no     *int     `json:"cloud_point_step_seq_no" default:""`
	Cloud_point_temp            *float64 `json:"cloud_point_temp" default:""`
	Cloud_point_temp_ouom       *string  `json:"cloud_point_temp_ouom" default:""`
	Dew_point_preferred_ind     *string  `json:"dew_point_preferred_ind" default:""`
	Dew_point_problem_ind       *string  `json:"dew_point_problem_ind" default:""`
	Dew_point_step_seq_no       *int     `json:"dew_point_step_seq_no" default:""`
	Dew_point_temp              *float64 `json:"dew_point_temp" default:""`
	Dew_point_temp_ouom         *string  `json:"dew_point_temp_ouom" default:""`
	Distillation_base_type      *string  `json:"distillation_base_type" default:""`
	Distillation_preferred_ind  *string  `json:"distillation_preferred_ind" default:""`
	Distillation_problem_ind    *string  `json:"distillation_problem_ind" default:""`
	Distillation_step_seq_no    *int     `json:"distillation_step_seq_no" default:""`
	Effective_date              *string  `json:"effective_date" default:""`
	Expiry_date                 *string  `json:"expiry_date" default:""`
	Flash_point_preferred_ind   *string  `json:"flash_point_preferred_ind" default:""`
	Flash_point_problem_ind     *string  `json:"flash_point_problem_ind" default:""`
	Flash_point_step_seq_no     *int     `json:"flash_point_step_seq_no" default:""`
	Flash_point_temp            *float64 `json:"flash_point_temp" default:""`
	Flash_point_temp_ouom       *string  `json:"flash_point_temp_ouom" default:""`
	Gor                         *float64 `json:"gor" default:""`
	Gor_ouom                    *string  `json:"gor_ouom" default:""`
	Measured_pressure           *float64 `json:"measured_pressure" default:""`
	Measured_pressure_ouom      *string  `json:"measured_pressure_ouom" default:""`
	Measured_volume             *float64 `json:"measured_volume" default:""`
	Measured_volume_ouom        *string  `json:"measured_volume_ouom" default:""`
	Method_type                 *string  `json:"method_type" default:""`
	Oil_analysis_temp           *float64 `json:"oil_analysis_temp" default:""`
	Oil_analysis_temp_ouom      *string  `json:"oil_analysis_temp_ouom" default:""`
	Oil_color                   *string  `json:"oil_color" default:""`
	Oil_color_preferred_ind     *string  `json:"oil_color_preferred_ind" default:""`
	Oil_color_problem_ind       *string  `json:"oil_color_problem_ind" default:""`
	Oil_color_step_seq_no       *int     `json:"oil_color_step_seq_no" default:""`
	Oil_density                 *float64 `json:"oil_density" default:""`
	Oil_density_ouom            *string  `json:"oil_density_ouom" default:""`
	Oil_density_preferred_ind   *string  `json:"oil_density_preferred_ind" default:""`
	Oil_density_problem_ind     *string  `json:"oil_density_problem_ind" default:""`
	Oil_density_step_seq_no     *int     `json:"oil_density_step_seq_no" default:""`
	Oil_gravity                 *float64 `json:"oil_gravity" default:""`
	Oil_gravity_ouom            *string  `json:"oil_gravity_ouom" default:""`
	Oil_gravity_preferred_ind   *string  `json:"oil_gravity_preferred_ind" default:""`
	Oil_gravity_step_seq_no     *int     `json:"oil_gravity_step_seq_no" default:""`
	Oil_gravity_temp            *float64 `json:"oil_gravity_temp" default:""`
	Oil_gravity_temp_ouom       *string  `json:"oil_gravity_temp_ouom" default:""`
	Oil_gravity_temp_prefer_ind *string  `json:"oil_gravity_temp_prefer_ind" default:""`
	Oil_gravity_temp_prob_ind   *string  `json:"oil_gravity_temp_prob_ind" default:""`
	Oil_gravity_temp_step_seq   *int     `json:"oil_gravity_temp_step_seq" default:""`
	Oil_type                    *string  `json:"oil_type" default:""`
	Ppdm_guid                   *string  `json:"ppdm_guid" default:""`
	Relative_density            *float64 `json:"relative_density" default:""`
	Relative_density_ouom       *string  `json:"relative_density_ouom" default:""`
	Relative_std_density        *float64 `json:"relative_std_density" default:""`
	Relative_std_density_ouom   *string  `json:"relative_std_density_ouom" default:""`
	Remark                      *string  `json:"remark" default:""`
	Reservoir_press             *float64 `json:"reservoir_press" default:""`
	Reservoir_press_ouom        *string  `json:"reservoir_press_ouom" default:""`
	Reservoir_temp              *float64 `json:"reservoir_temp" default:""`
	Reservoir_temp_ouom         *string  `json:"reservoir_temp_ouom" default:""`
	Sediment_content            *float64 `json:"sediment_content" default:""`
	Sediment_content_ouom       *string  `json:"sediment_content_ouom" default:""`
	Sediment_preferred_ind      *string  `json:"sediment_preferred_ind" default:""`
	Sediment_problem_ind        *string  `json:"sediment_problem_ind" default:""`
	Sediment_step_seq_no        *int     `json:"sediment_step_seq_no" default:""`
	Shrinkage_factor            *float64 `json:"shrinkage_factor" default:""`
	Shrink_factor_preferred_ind *string  `json:"shrink_factor_preferred_ind" default:""`
	Shrink_factor_problem_ind   *string  `json:"shrink_factor_problem_ind" default:""`
	Shrink_factor_step_seq_no   *int     `json:"shrink_factor_step_seq_no" default:""`
	Source                      *string  `json:"source" default:""`
	Sulphur_content             *float64 `json:"sulphur_content" default:""`
	Vapour_press                *float64 `json:"vapour_press" default:""`
	Vapour_press_ouom           *string  `json:"vapour_press_ouom" default:""`
	Vapour_press_temp           *float64 `json:"vapour_press_temp" default:""`
	Vapour_press_temp_ouom      *string  `json:"vapour_press_temp_ouom" default:""`
	Water_content               *float64 `json:"water_content" default:""`
	Water_content_ouom          *string  `json:"water_content_ouom" default:""`
	Wax_content                 *float64 `json:"wax_content" default:""`
	Wax_content_ouom            *string  `json:"wax_content_ouom" default:""`
	Wax_content_preferred_ind   *string  `json:"wax_content_preferred_ind" default:""`
	Wax_content_problem_ind     *string  `json:"wax_content_problem_ind" default:""`
	Wax_content_step_seq_no     *int     `json:"wax_content_step_seq_no" default:""`
	Row_changed_by              *string  `json:"row_changed_by" default:""`
	Row_changed_date            *string  `json:"row_changed_date" default:""`
	Row_created_by              *string  `json:"row_created_by" default:""`
	Row_created_date            *string  `json:"row_created_date" default:""`
	Row_effective_date          *string  `json:"row_effective_date" default:""`
	Row_expiry_date             *string  `json:"row_expiry_date" default:""`
	Row_quality                 *string  `json:"row_quality" default:""`
}
