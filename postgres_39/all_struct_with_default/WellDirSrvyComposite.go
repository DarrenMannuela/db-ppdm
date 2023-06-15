package dto

type Well_dir_srvy_composite struct {
	Composite_uwi       string  `json:"composite_uwi" default:"source"`
	Composite_survey_id string  `json:"composite_survey_id" default:"source"`
	Composite_source    string  `json:"composite_source" default:"source"`
	Composite_obs_no    int     `json:"composite_obs_no" default:"1"`
	Active_ind          *string `json:"active_ind" default:""`
	Depth_obs_no_from   *int    `json:"depth_obs_no_from" default:""`
	Depth_obs_no_to     *int    `json:"depth_obs_no_to" default:""`
	Effective_date      *string `json:"effective_date" default:""`
	Expiry_date         *string `json:"expiry_date" default:""`
	Input_source        *string `json:"input_source" default:""`
	Input_survey_id     *string `json:"input_survey_id" default:""`
	Input_uwi           *string `json:"input_uwi" default:""`
	Ppdm_guid           *string `json:"ppdm_guid" default:""`
	Remark              *string `json:"remark" default:""`
	Source              *string `json:"source" default:""`
	Row_changed_by      *string `json:"row_changed_by" default:""`
	Row_changed_date    *string `json:"row_changed_date" default:""`
	Row_created_by      *string `json:"row_created_by" default:""`
	Row_created_date    *string `json:"row_created_date" default:""`
	Row_effective_date  *string `json:"row_effective_date" default:""`
	Row_expiry_date     *string `json:"row_expiry_date" default:""`
	Row_quality         *string `json:"row_quality" default:""`
}
