package dto

type Ppdm_sw_app_function struct {
	Sw_application_id    string  `json:"sw_application_id" default:"source"`
	Sw_app_function      string  `json:"sw_app_function" default:"source"`
	Function_obs_no      int     `json:"function_obs_no" default:"1"`
	Abbreviation         *string `json:"abbreviation" default:""`
	Active_ind           *string `json:"active_ind" default:""`
	Effective_date       *string `json:"effective_date" default:""`
	Expiry_date          *string `json:"expiry_date" default:""`
	Long_name            *string `json:"long_name" default:""`
	Ppdm_guid            *string `json:"ppdm_guid" default:""`
	Remark               *string `json:"remark" default:""`
	Short_name           *string `json:"short_name" default:""`
	Source               *string `json:"source" default:""`
	Sw_app_function_type *string `json:"sw_app_function_type" default:""`
	Row_changed_by       *string `json:"row_changed_by" default:""`
	Row_changed_date     *string `json:"row_changed_date" default:""`
	Row_created_by       *string `json:"row_created_by" default:""`
	Row_created_date     *string `json:"row_created_date" default:""`
	Row_effective_date   *string `json:"row_effective_date" default:""`
	Row_expiry_date      *string `json:"row_expiry_date" default:""`
	Row_quality          *string `json:"row_quality" default:""`
}
