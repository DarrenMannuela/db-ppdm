package dto

type Equipment struct {
	Equipment_id        string  `json:"equipment_id" default:"source"`
	Acquire_date_new    *string `json:"acquire_date_new" default:""`
	Active_ind          *string `json:"active_ind" default:""`
	Catalogue_equip_id  *string `json:"catalogue_equip_id" default:""`
	Commission_date     *string `json:"commission_date" default:""`
	Current_owner_ba_id *string `json:"current_owner_ba_id" default:""`
	Decommission_date   *string `json:"decommission_date" default:""`
	Description         *string `json:"description" default:""`
	Effective_date      *string `json:"effective_date" default:""`
	Equipment_group     *string `json:"equipment_group" default:""`
	Equipment_name      *string `json:"equipment_name" default:""`
	Equipment_type      *string `json:"equipment_type" default:""`
	Expiry_date         *string `json:"expiry_date" default:""`
	Manufacturer_ba_id  *string `json:"manufacturer_ba_id" default:""`
	Ppdm_guid           *string `json:"ppdm_guid" default:""`
	Purchase_date       *string `json:"purchase_date" default:""`
	Reference_num       *string `json:"reference_num" default:""`
	Remark              *string `json:"remark" default:""`
	Serial_num          *string `json:"serial_num" default:""`
	Source              *string `json:"source" default:""`
	Row_changed_by      *string `json:"row_changed_by" default:""`
	Row_changed_date    *string `json:"row_changed_date" default:""`
	Row_created_by      *string `json:"row_created_by" default:""`
	Row_created_date    *string `json:"row_created_date" default:""`
	Row_effective_date  *string `json:"row_effective_date" default:""`
	Row_expiry_date     *string `json:"row_expiry_date" default:""`
	Row_quality         *string `json:"row_quality" default:""`
}
