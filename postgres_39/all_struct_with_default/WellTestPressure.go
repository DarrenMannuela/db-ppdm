package dto

type Well_test_pressure struct {
	Uwi                 string   `json:"uwi" default:"source"`
	Source              string   `json:"source" default:"source"`
	Test_type           string   `json:"test_type" default:"source"`
	Run_num             string   `json:"run_num" default:"source"`
	Test_num            string   `json:"test_num" default:"source"`
	Period_type         string   `json:"period_type" default:"source"`
	Period_obs_no       int      `json:"period_obs_no" default:"1"`
	Active_ind          *string  `json:"active_ind" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	End_pressure        *float64 `json:"end_pressure" default:""`
	End_pressure_ouom   *string  `json:"end_pressure_ouom" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Quality_code        *string  `json:"quality_code" default:""`
	Recorder_id         *string  `json:"recorder_id" default:""`
	Remark              *string  `json:"remark" default:""`
	Start_pressure      *float64 `json:"start_pressure" default:""`
	Start_pressure_ouom *string  `json:"start_pressure_ouom" default:""`
	Summary_ind         *string  `json:"summary_ind" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
