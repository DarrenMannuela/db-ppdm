package dto

type Sp_desc_text struct {
	Spatial_description_id string  `json:"spatial_description_id" default:"source"`
	Spatial_obs_no         int     `json:"spatial_obs_no" default:"1"`
	Text_id                string  `json:"text_id" default:"source"`
	Description_seq_no     int     `json:"description_seq_no" default:"1"`
	Active_ind             *string `json:"active_ind" default:""`
	Description            *string `json:"description" default:""`
	Effective_date         *string `json:"effective_date" default:""`
	Expiry_date            *string `json:"expiry_date" default:""`
	Ppdm_guid              *string `json:"ppdm_guid" default:""`
	Remark                 *string `json:"remark" default:""`
	Source                 *string `json:"source" default:""`
	Row_changed_by         *string `json:"row_changed_by" default:""`
	Row_changed_date       *string `json:"row_changed_date" default:""`
	Row_created_by         *string `json:"row_created_by" default:""`
	Row_created_date       *string `json:"row_created_date" default:""`
	Row_effective_date     *string `json:"row_effective_date" default:""`
	Row_expiry_date        *string `json:"row_expiry_date" default:""`
	Row_quality            *string `json:"row_quality" default:""`
}
