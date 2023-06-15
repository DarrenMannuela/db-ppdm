package dto

import (
	"time"
)

type Well_dir_srvy_station struct {
	Uwi                     string     `json:"uwi"`
	Survey_id               string     `json:"survey_id"`
	Source                  string     `json:"source"`
	Depth_obs_no            int        `json:"depth_obs_no"`
	Accuracy_problem_ind    *string    `json:"accuracy_problem_ind"`
	Accuracy_problem_reason *string    `json:"accuracy_problem_reason"`
	Active_ind              *string    `json:"active_ind"`
	Azimuth                 *float64   `json:"azimuth"`
	Azimuth_ouom            *string    `json:"azimuth_ouom"`
	Closure_dist            *float64   `json:"closure_dist"`
	Closure_dist_dir        *float64   `json:"closure_dist_dir"`
	Closure_dist_dir_ouom   *string    `json:"closure_dist_dir_ouom"`
	Closure_dist_ouom       *string    `json:"closure_dist_ouom"`
	Delta_lat               *float64   `json:"delta_lat"`
	Delta_long              *float64   `json:"delta_long"`
	Dog_leg_severity        *float64   `json:"dog_leg_severity"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Inclination             *float64   `json:"inclination"`
	Inclination_ouom        *string    `json:"inclination_ouom"`
	Latitude                *float64   `json:"latitude"`
	Longitude               *float64   `json:"longitude"`
	Period_obs_no           *int       `json:"period_obs_no"`
	Point_type              *string    `json:"point_type"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Rad_uncert_dist         *float64   `json:"rad_uncert_dist"`
	Rad_uncert_dist_ouom    *string    `json:"rad_uncert_dist_ouom"`
	Rad_uncert_dist_reason  *string    `json:"rad_uncert_dist_reason"`
	Remark                  *string    `json:"remark"`
	Rpt_azimuth             *float64   `json:"rpt_azimuth"`
	Rpt_azimuth_ouom        *string    `json:"rpt_azimuth_ouom"`
	Station_id              *string    `json:"station_id"`
	Station_md              *float64   `json:"station_md"`
	Station_md_ouom         *string    `json:"station_md_ouom"`
	Station_tvd             *float64   `json:"station_tvd"`
	Station_tvdss           *float64   `json:"station_tvdss"`
	Station_tvdss_ouom      *string    `json:"station_tvdss_ouom"`
	Station_tvd_ouom        *string    `json:"station_tvd_ouom"`
	Survey_run_num          *string    `json:"survey_run_num"`
	Survey_type             *string    `json:"survey_type"`
	Vertical_section        *float64   `json:"vertical_section"`
	Vertical_section_ouom   *string    `json:"vertical_section_ouom"`
	X_offset                *float64   `json:"x_offset"`
	X_offset_ew_direction   *string    `json:"x_offset_ew_direction"`
	X_offset_ouom           *string    `json:"x_offset_ouom"`
	Y_offset                *float64   `json:"y_offset"`
	Y_offset_ns_direction   *string    `json:"y_offset_ns_direction"`
	Y_offset_ouom           *string    `json:"y_offset_ouom"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
