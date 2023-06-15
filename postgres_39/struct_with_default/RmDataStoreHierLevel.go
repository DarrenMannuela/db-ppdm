package dto

type Rm_data_store_hier_level struct {
	Data_store_hier_id   string  `json:"data_store_hier_id" default:"source"`
	Hier_level_id        string  `json:"hier_level_id" default:"source"`
	Active_ind           *string `json:"active_ind" default:""`
	Data_store_type      *string `json:"data_store_type" default:""`
	Effective_date       *string `json:"effective_date" default:""`
	Expiry_date          *string `json:"expiry_date" default:""`
	Level_name           *string `json:"level_name" default:""`
	Level_seq_no         *int    `json:"level_seq_no" default:""`
	Parent_hier_level_id *string `json:"parent_hier_level_id" default:""`
	Ppdm_guid            *string `json:"ppdm_guid" default:""`
	Remark               *string `json:"remark" default:""`
	Source               *string `json:"source" default:""`
	Row_changed_by       *string `json:"row_changed_by" default:""`
	Row_changed_date     *string `json:"row_changed_date" default:""`
	Row_created_by       *string `json:"row_created_by" default:""`
	Row_created_date     *string `json:"row_created_date" default:""`
	Row_effective_date   *string `json:"row_effective_date" default:""`
	Row_expiry_date      *string `json:"row_expiry_date" default:""`
	Row_quality          *string `json:"row_quality" default:""`
}
