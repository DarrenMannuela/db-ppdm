package dto

type Well_data struct{

Id      int  `json:"id" default:"" gorm:"primaryKey"`
Ba_long_name   *string  `json:"ba_long_name" default:""`
Ba_type   *string  `json:"ba_type" default:""`
Area_id   *string  `json:"area_id" default:""`
Area_type   *string  `json:"area_type" default:""`
Field_name   *string  `json:"field_name" default:""`
Well_name   *string  `json:"well_name" default:""`
Alias_long_name   *string  `json:"alias_long_name" default:""`
Uwi_   *string  `json:"uwi_" default:""`
Status_type   *string  `json:"status_type" default:""`
Current_class   *string  `json:"current_class" default:""`
Well_level_type   *string  `json:"well_level_type" default:""`
Profile_type   *string  `json:"profile_type" default:""`
Projected_strat_unit_id   *string  `json:"projected_strat_unit_id" default:""`
Surface_longitude   *string  `json:"surface_longitude" default:""`
Surface_latitude   *string  `json:"surface_latitude" default:""`
Easting_   *string  `json:"easting_" default:""`
Easting_ouom   *string  `json:"easting_ouom" default:""`
Northing_   *string  `json:"northing_" default:""`
Northing_ouom   *string  `json:"northing_ouom" default:""`
Utm_quadrant   *string  `json:"utm_quadrant" default:""`
Projection_type   *string  `json:"projection_type" default:""`
Geodetic_datum_name   *string  `json:"geodetic_datum_name" default:""`
Environment_type   *string  `json:"environment_type" default:""`
Line_name   *string  `json:"line_name" default:""`
Seis_point_id   *string  `json:"seis_point_id" default:""`
Spud_date   *string  `json:"spud_date" default:""`
Final_drill_date   *string  `json:"final_drill_date" default:""`
Completion_date   *string  `json:"completion_date" default:""`
Rotary_table_elev   *string  `json:"rotary_table_elev" default:""`
Rotary_table_elev_ouom   *string  `json:"rotary_table_elev_ouom" default:""`
Kb_elev   *string  `json:"kb_elev" default:""`
Kb_elev_ouom   *string  `json:"kb_elev_ouom" default:""`
Derrick_floor_elev   *string  `json:"derrick_floor_elev" default:""`
Derrick_floor_elev_ouom   *string  `json:"derrick_floor_elev_ouom" default:""`
Water_depth   *string  `json:"water_depth" default:""`
Water_depth_ouom   *string  `json:"water_depth_ouom" default:""`
Ground_elev   *string  `json:"ground_elev" default:""`
Ground_elev_ouom   *string  `json:"ground_elev_ouom" default:""`
Depth_datum_elev   *string  `json:"depth_datum_elev" default:""`
Drill_td   *string  `json:"drill_td" default:""`
Drill_td_ouom   *string  `json:"drill_td_ouom" default:""`
Log_td   *string  `json:"log_td" default:""`
Log_td_ouom   *string  `json:"log_td_ouom" default:""`
Max_tvd   *string  `json:"max_tvd" default:""`
Max_tvd_ouom   *string  `json:"max_tvd_ouom" default:""`
Projected_depth   *string  `json:"projected_depth" default:""`
Projected_depth_ouom   *string  `json:"projected_depth_ouom" default:""`
Final_td   *string  `json:"final_td" default:""`
Final_td_ouom   *string  `json:"final_td_ouom" default:""`
Operator_ba_id   *string  `json:"operator_ba_id" default:""`
Rig_name   *string  `json:"rig_name" default:""`
Rig_type   *string  `json:"rig_type" default:""`
Test_result_code   *string  `json:"test_result_code" default:""`
Remark_   *string  `json:"remark_" default:""`
Source_   *string  `json:"source_" default:""`
Qc_status   *string  `json:"qc_status" default:""`
Checked_by_ba_id   *string  `json:"checked_by_ba_id" default:""`
Tubing_obs_no   *int  `json:"tubing_obs_no" default:""`
Outside_diameter   *string  `json:"outside_diameter" default:""`
Outside_diameter_ouom   *string  `json:"outside_diameter_ouom" default:""`
Base_depth   *string  `json:"base_depth" default:""`
Base_depth_ouom   *string  `json:"base_depth_ouom" default:""`
}