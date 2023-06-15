package dto

type Strat_hierarchy_desc struct {
	Hierarchy_desc_id      string  `json:"hierarchy_desc_id" default:"source"`
	Hierarchy_seq_no       int     `json:"hierarchy_seq_no" default:"1"`
	Active_ind             *string `json:"active_ind" default:""`
	Effective_date         *string `json:"effective_date" default:""`
	Expiry_date            *string `json:"expiry_date" default:""`
	Parent_strat_unit_type *string `json:"parent_strat_unit_type" default:""`
	Ppdm_guid              *string `json:"ppdm_guid" default:""`
	Remark                 *string `json:"remark" default:""`
	Source                 *string `json:"source" default:""`
	Source_document_id     *string `json:"source_document_id" default:""`
	Strat_hierarchy_type   *string `json:"strat_hierarchy_type" default:""`
	Strat_type             *string `json:"strat_type" default:""`
	Strat_unit_type        *string `json:"strat_unit_type" default:""`
	Row_changed_by         *string `json:"row_changed_by" default:""`
	Row_changed_date       *string `json:"row_changed_date" default:""`
	Row_created_by         *string `json:"row_created_by" default:""`
	Row_created_date       *string `json:"row_created_date" default:""`
	Row_effective_date     *string `json:"row_effective_date" default:""`
	Row_expiry_date        *string `json:"row_expiry_date" default:""`
	Row_quality            *string `json:"row_quality" default:""`
}
