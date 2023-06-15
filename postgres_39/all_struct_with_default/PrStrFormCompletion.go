package dto

type Pr_str_form_completion struct {
	Uwi                    string  `json:"uwi" default:"source"`
	Prod_string_source     string  `json:"prod_string_source" default:"source"`
	String_id              string  `json:"string_id" default:"source"`
	Pr_str_form_obs_no     int     `json:"pr_str_form_obs_no" default:"1"`
	Completion_source      string  `json:"completion_source" default:"source"`
	Completion_obs_no      int     `json:"completion_obs_no" default:"1"`
	Form_completion_obs_no int     `json:"form_completion_obs_no" default:"1"`
	Active_ind             *string `json:"active_ind" default:""`
	Completion_status      *string `json:"completion_status" default:""`
	Effective_date         *string `json:"effective_date" default:""`
	Expiry_date            *string `json:"expiry_date" default:""`
	Ppdm_guid              *string `json:"ppdm_guid" default:""`
	Remark                 *string `json:"remark" default:""`
	Source                 *string `json:"source" default:""`
	Status_type            *string `json:"status_type" default:""`
	Row_changed_by         *string `json:"row_changed_by" default:""`
	Row_changed_date       *string `json:"row_changed_date" default:""`
	Row_created_by         *string `json:"row_created_by" default:""`
	Row_created_date       *string `json:"row_created_date" default:""`
	Row_effective_date     *string `json:"row_effective_date" default:""`
	Row_expiry_date        *string `json:"row_expiry_date" default:""`
	Row_quality            *string `json:"row_quality" default:""`
}
