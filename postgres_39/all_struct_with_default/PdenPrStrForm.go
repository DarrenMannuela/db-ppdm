package dto

type Pden_pr_str_form struct {
	Pden_subtype       string  `json:"pden_subtype" default:"source"`
	Pden_id            string  `json:"pden_id" default:"source"`
	Pden_source        string  `json:"pden_source" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Pr_str_form_obs_no *int    `json:"pr_str_form_obs_no" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Strat_name_set_id  *string `json:"strat_name_set_id" default:""`
	Strat_unit_id      *string `json:"strat_unit_id" default:""`
	String_id          *string `json:"string_id" default:""`
	String_source      *string `json:"string_source" default:""`
	Uwi                *string `json:"uwi" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
