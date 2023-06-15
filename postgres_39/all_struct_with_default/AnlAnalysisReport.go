package dto

type Anl_analysis_report struct {
	Analysis_id        string   `json:"analysis_id" default:"source"`
	Anl_source         string   `json:"anl_source" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Analysis_date      *string  `json:"analysis_date" default:""`
	Analysis_purpose   *string  `json:"analysis_purpose" default:""`
	Analysis_quality   *string  `json:"analysis_quality" default:""`
	Base_depth         *float64 `json:"base_depth" default:""`
	Base_depth_ouom    *string  `json:"base_depth_ouom" default:""`
	Base_strat_unit_id *string  `json:"base_strat_unit_id" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	End_date           *string  `json:"end_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Received_date      *string  `json:"received_date" default:""`
	Remark             *string  `json:"remark" default:""`
	Reported_date      *string  `json:"reported_date" default:""`
	Reported_tvd       *float64 `json:"reported_tvd" default:""`
	Reported_tvd_ouom  *string  `json:"reported_tvd_ouom" default:""`
	Sample_date        *string  `json:"sample_date" default:""`
	Sample_location    *string  `json:"sample_location" default:""`
	Start_date         *string  `json:"start_date" default:""`
	Strat_name_set_id  *string  `json:"strat_name_set_id" default:""`
	Study_type         *string  `json:"study_type" default:""`
	Top_depth          *float64 `json:"top_depth" default:""`
	Top_depth_ouom     *string  `json:"top_depth_ouom" default:""`
	Top_strat_unit_id  *string  `json:"top_strat_unit_id" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
