package dto

type R_source_origin struct {
	Source             string  `json:"source" default:"source"`
	Origin_obs_no      int     `json:"origin_obs_no" default:"1"`
	Abbreviation       *string `json:"abbreviation" default:""`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Long_name          *string `json:"long_name" default:""`
	Owner_ba_id        *string `json:"owner_ba_id" default:""`
	Physical_item_id   *string `json:"physical_item_id" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Project_id         *string `json:"project_id" default:""`
	Remark             *string `json:"remark" default:""`
	Row_source         *string `json:"row_source" default:""`
	Short_name         *string `json:"short_name" default:""`
	Source_document    *string `json:"source_document" default:""`
	Sw_application_id  *string `json:"sw_application_id" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
