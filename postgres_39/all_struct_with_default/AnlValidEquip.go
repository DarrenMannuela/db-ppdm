package dto

type Anl_valid_equip struct {
	Method_set_id      string  `json:"method_set_id" default:"source"`
	Method_id          string  `json:"method_id" default:"source"`
	Equip_obs_no       int     `json:"equip_obs_no" default:"1"`
	Active_ind         *string `json:"active_ind" default:""`
	Catalogue_equip_id *string `json:"catalogue_equip_id" default:""`
	Confidence_type    *string `json:"confidence_type" default:""`
	Description        *string `json:"description" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Equip_id           *string `json:"equip_id" default:""`
	Equip_role         *string `json:"equip_role" default:""`
	Equip_setting      *string `json:"equip_setting" default:""`
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
