package dto

type Well_dir_srvy_st_version struct {
	Uwi                     string   `json:"uwi" default:"source"`
	Survey_id               string   `json:"survey_id" default:"source"`
	Source                  string   `json:"source" default:"source"`
	Version_obs_no          int      `json:"version_obs_no" default:"1"`
	Depth_obs_no            int      `json:"depth_obs_no" default:"1"`
	Accuracy_problem_ind    *string  `json:"accuracy_problem_ind" default:""`
	Accuracy_problem_reason *string  `json:"accuracy_problem_reason" default:""`
	Active_ind              *string  `json:"active_ind" default:""`
	Azimuth                 *float64 `json:"azimuth" default:""`
	Azimuth_ouom            *string  `json:"azimuth_ouom" default:""`
	Closure_dist            *float64 `json:"closure_dist" default:""`
	Closure_dist_dir        *float64 `json:"closure_dist_dir" default:""`
	Closure_dist_dir_ouom   *string  `json:"closure_dist_dir_ouom" default:""`
	Closure_dist_ouom       *string  `json:"closure_dist_ouom" default:""`
	Delta_lat               *float64 `json:"delta_lat" default:""`
	Delta_long              *float64 `json:"delta_long" default:""`
	Dog_leg_severity        *float64 `json:"dog_leg_severity" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Inclination             *float64 `json:"inclination" default:""`
	Inclination_ouom        *string  `json:"inclination_ouom" default:""`
	Latitude                *float64 `json:"latitude" default:""`
	Longitude               *float64 `json:"longitude" default:""`
	Period_obs_no           *int     `json:"period_obs_no" default:""`
	Point_type              *string  `json:"point_type" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Rad_uncert_dist         *float64 `json:"rad_uncert_dist" default:""`
	Rad_uncert_dist_ouom    *string  `json:"rad_uncert_dist_ouom" default:""`
	Rad_uncert_dist_reason  *string  `json:"rad_uncert_dist_reason" default:""`
	Remark                  *string  `json:"remark" default:""`
	Rpt_azimuth             *float64 `json:"rpt_azimuth" default:""`
	Rpt_azimuth_ouom        *string  `json:"rpt_azimuth_ouom" default:""`
	Station_id              *string  `json:"station_id" default:""`
	Station_md              *float64 `json:"station_md" default:""`
	Station_md_ouom         *string  `json:"station_md_ouom" default:""`
	Station_tvd             *float64 `json:"station_tvd" default:""`
	Station_tvdss           *float64 `json:"station_tvdss" default:""`
	Station_tvdss_ouom      *string  `json:"station_tvdss_ouom" default:""`
	Station_tvd_ouom        *string  `json:"station_tvd_ouom" default:""`
	Survey_run_num          *string  `json:"survey_run_num" default:""`
	Survey_type             *string  `json:"survey_type" default:""`
	Vertical_section        *float64 `json:"vertical_section" default:""`
	Vertical_section_ouom   *string  `json:"vertical_section_ouom" default:""`
	X_offset                *float64 `json:"x_offset" default:""`
	X_offset_ew_direction   *string  `json:"x_offset_ew_direction" default:""`
	X_offset_ouom           *string  `json:"x_offset_ouom" default:""`
	Y_offset                *float64 `json:"y_offset" default:""`
	Y_offset_ns_direction   *string  `json:"y_offset_ns_direction" default:""`
	Y_offset_ouom           *string  `json:"y_offset_ouom" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
