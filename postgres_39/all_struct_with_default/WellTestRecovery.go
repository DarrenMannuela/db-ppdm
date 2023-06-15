package dto

type Well_test_recovery struct {
	Uwi                     string   `json:"uwi" default:"source"`
	Source                  string   `json:"source" default:"source"`
	Test_type               string   `json:"test_type" default:"source"`
	Run_num                 string   `json:"run_num" default:"source"`
	Test_num                string   `json:"test_num" default:"source"`
	Recovery_obs_no         int      `json:"recovery_obs_no" default:"1"`
	Active_ind              *string  `json:"active_ind" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Multiple_test_ind       *string  `json:"multiple_test_ind" default:""`
	Period_obs_no           *int     `json:"period_obs_no" default:""`
	Period_type             *string  `json:"period_type" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Recovery_amount         *float64 `json:"recovery_amount" default:""`
	Recovery_amount_ouom    *string  `json:"recovery_amount_ouom" default:""`
	Recovery_amount_percent *float64 `json:"recovery_amount_percent" default:""`
	Recovery_amount_uom     *string  `json:"recovery_amount_uom" default:""`
	Recovery_desc           *string  `json:"recovery_desc" default:""`
	Recovery_method         *string  `json:"recovery_method" default:""`
	Recovery_show_type      *string  `json:"recovery_show_type" default:""`
	Recovery_type           *string  `json:"recovery_type" default:""`
	Remark                  *string  `json:"remark" default:""`
	Reverse_circulation_ind *string  `json:"reverse_circulation_ind" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
