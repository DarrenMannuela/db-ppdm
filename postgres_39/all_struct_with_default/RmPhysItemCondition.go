package dto

type Rm_phys_item_condition struct {
	Physical_item_id   string  `json:"physical_item_id" default:"source"`
	Condition_id       string  `json:"condition_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Condition_type     *string `json:"condition_type" default:""`
	Correction_method  *string `json:"correction_method" default:""`
	Description        *string `json:"description" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Error_count        *int    `json:"error_count" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Origin_seq_no      *int    `json:"origin_seq_no" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
