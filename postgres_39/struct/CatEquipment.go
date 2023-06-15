package dto

import (
	"time"
)

type Cat_equipment struct {
	Catalogue_equip_id string     `json:"catalogue_equip_id"`
	Active_ind         *string    `json:"active_ind"`
	Cat_equip_group    *string    `json:"cat_equip_group"`
	Cat_equip_type     *string    `json:"cat_equip_type"`
	Effective_date     *time.Time `json:"effective_date"`
	Equipment_name     *string    `json:"equipment_name"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Install_loc_type   *string    `json:"install_loc_type"`
	Manufacturer_ba_id *string    `json:"manufacturer_ba_id"`
	Model_num          *string    `json:"model_num"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
