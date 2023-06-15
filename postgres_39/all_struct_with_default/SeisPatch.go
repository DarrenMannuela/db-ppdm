package dto

type Seis_patch struct {
	Patch_id           string  `json:"patch_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Channel_count      *int    `json:"channel_count" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Gap_count          *int    `json:"gap_count" default:""`
	Patch_layout       *string `json:"patch_layout" default:""`
	Patch_type         *string `json:"patch_type" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Roll_along_method  *string `json:"roll_along_method" default:""`
	Shot_gap_ind       *string `json:"shot_gap_ind" default:""`
	Source             *string `json:"source" default:""`
	Symmetric_ind      *string `json:"symmetric_ind" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
