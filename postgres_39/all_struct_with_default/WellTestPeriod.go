package dto

type Well_test_period struct {
	Uwi                  string   `json:"uwi" default:"source"`
	Source               string   `json:"source" default:"source"`
	Test_type            string   `json:"test_type" default:"source"`
	Run_num              string   `json:"run_num" default:"source"`
	Test_num             string   `json:"test_num" default:"source"`
	Period_type          string   `json:"period_type" default:"source"`
	Period_obs_no        int      `json:"period_obs_no" default:"1"`
	Active_ind           *string  `json:"active_ind" default:""`
	Casing_pressure      *float64 `json:"casing_pressure" default:""`
	Casing_pressure_ouom *string  `json:"casing_pressure_ouom" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Period_duration      *float64 `json:"period_duration" default:""`
	Period_duration_ouom *string  `json:"period_duration_ouom" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Remark               *string  `json:"remark" default:""`
	Tubing_pressure      *float64 `json:"tubing_pressure" default:""`
	Tubing_pressure_ouom *string  `json:"tubing_pressure_ouom" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
