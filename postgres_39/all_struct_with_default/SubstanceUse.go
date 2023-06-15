package dto

type Substance_use struct {
	Substance_id            string  `json:"substance_id" default:"source"`
	Substance_use_obs_no    int     `json:"substance_use_obs_no" default:"1"`
	Active_ind              *string `json:"active_ind" default:""`
	Area_id                 *string `json:"area_id" default:""`
	Area_type               *string `json:"area_type" default:""`
	Business_associate_id   *string `json:"business_associate_id" default:""`
	Contract_id             *string `json:"contract_id" default:""`
	Effective_date          *string `json:"effective_date" default:""`
	Expiry_date             *string `json:"expiry_date" default:""`
	Land_right_id           *string `json:"land_right_id" default:""`
	Land_right_subtype      *string `json:"land_right_subtype" default:""`
	Ppdm_column_name        *string `json:"ppdm_column_name" default:""`
	Ppdm_guid               *string `json:"ppdm_guid" default:""`
	Ppdm_system_id          *string `json:"ppdm_system_id" default:""`
	Ppdm_table_name         *string `json:"ppdm_table_name" default:""`
	Preferred_ind           *string `json:"preferred_ind" default:""`
	Project_id              *string `json:"project_id" default:""`
	Remark                  *string `json:"remark" default:""`
	Source                  *string `json:"source" default:""`
	Substance_alias_id      *string `json:"substance_alias_id" default:""`
	Substance_use_rule_desc *string `json:"substance_use_rule_desc" default:""`
	Substance_use_rule_type *string `json:"substance_use_rule_type" default:""`
	Row_changed_by          *string `json:"row_changed_by" default:""`
	Row_changed_date        *string `json:"row_changed_date" default:""`
	Row_created_by          *string `json:"row_created_by" default:""`
	Row_created_date        *string `json:"row_created_date" default:""`
	Row_effective_date      *string `json:"row_effective_date" default:""`
	Row_expiry_date         *string `json:"row_expiry_date" default:""`
	Row_quality             *string `json:"row_quality" default:""`
}
