package dto

type Anl_valid_ba struct {
	Method_set_id      string  `json:"method_set_id" default:"source"`
	Method_id          string  `json:"method_id" default:"source"`
	Valid_ba_id        string  `json:"valid_ba_id" default:"source"`
	Valid_ba_obs_no    int     `json:"valid_ba_obs_no" default:"1"`
	Active_ind         *string `json:"active_ind" default:""`
	Ba_role_type       *string `json:"ba_role_type" default:""`
	Confidence_type    *string `json:"confidence_type" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
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
