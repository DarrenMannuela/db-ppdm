package dto

type Rm_data_store_media struct {
	Store_id           string  `json:"store_id" default:"source"`
	Structure_obs_no   int     `json:"structure_obs_no" default:"1"`
	Media_obs_no       int     `json:"media_obs_no" default:"1"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Media_type         *string `json:"media_type" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Total_capacity     *int    `json:"total_capacity" default:""`
	Used_capacity      *int    `json:"used_capacity" default:""`
	Used_capacity_date *string `json:"used_capacity_date" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
