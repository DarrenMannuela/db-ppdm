package dto

type Sf_equipment struct {
	Sf_subtype          string  `json:"sf_subtype" default:"source"`
	Support_facility_id string  `json:"support_facility_id" default:"source"`
	Component_id        string  `json:"component_id" default:"source"`
	Active_ind          *string `json:"active_ind" default:""`
	Catalogue_equip_id  *string `json:"catalogue_equip_id" default:""`
	Component_count     *int    `json:"component_count" default:""`
	Description         *string `json:"description" default:""`
	Effective_date      *string `json:"effective_date" default:""`
	Equipment_id        *string `json:"equipment_id" default:""`
	Equipment_name      *string `json:"equipment_name" default:""`
	Expiry_date         *string `json:"expiry_date" default:""`
	Install_date        *string `json:"install_date" default:""`
	Ppdm_guid           *string `json:"ppdm_guid" default:""`
	Purchase_date       *string `json:"purchase_date" default:""`
	Reference_num       *string `json:"reference_num" default:""`
	Remark              *string `json:"remark" default:""`
	Remove_date         *string `json:"remove_date" default:""`
	Source              *string `json:"source" default:""`
	Row_changed_by      *string `json:"row_changed_by" default:""`
	Row_changed_date    *string `json:"row_changed_date" default:""`
	Row_created_by      *string `json:"row_created_by" default:""`
	Row_created_date    *string `json:"row_created_date" default:""`
	Row_effective_date  *string `json:"row_effective_date" default:""`
	Row_expiry_date     *string `json:"row_expiry_date" default:""`
	Row_quality         *string `json:"row_quality" default:""`
}
