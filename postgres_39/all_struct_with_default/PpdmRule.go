package dto

type Ppdm_rule struct {
	Rule_id            string  `json:"rule_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Current_status     *string `json:"current_status" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Rule_class         *string `json:"rule_class" default:""`
	Rule_desc          *string `json:"rule_desc" default:""`
	Rule_purpose       *string `json:"rule_purpose" default:""`
	Rule_query         *string `json:"rule_query" default:""`
	Source             *string `json:"source" default:""`
	Use_condition_desc *string `json:"use_condition_desc" default:""`
	Use_condition_type *string `json:"use_condition_type" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
