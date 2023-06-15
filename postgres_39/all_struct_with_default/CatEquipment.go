package dto

type Cat_equipment struct {
	Catalogue_equip_id string  `json:"catalogue_equip_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Cat_equip_group    *string `json:"cat_equip_group" default:""`
	Cat_equip_type     *string `json:"cat_equip_type" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Equipment_name     *string `json:"equipment_name" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Install_loc_type   *string `json:"install_loc_type" default:""`
	Manufacturer_ba_id *string `json:"manufacturer_ba_id" default:""`
	Model_num          *string `json:"model_num" default:""`
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
