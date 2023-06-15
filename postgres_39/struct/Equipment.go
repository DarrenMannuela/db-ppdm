package dto

import (
	"time"
)

type Equipment struct {
	Equipment_id        string     `json:"equipment_id"`
	Acquire_date_new    *time.Time `json:"acquire_date_new"`
	Active_ind          *string    `json:"active_ind"`
	Catalogue_equip_id  *string    `json:"catalogue_equip_id"`
	Commission_date     *time.Time `json:"commission_date"`
	Current_owner_ba_id *string    `json:"current_owner_ba_id"`
	Decommission_date   *time.Time `json:"decommission_date"`
	Description         *string    `json:"description"`
	Effective_date      *time.Time `json:"effective_date"`
	Equipment_group     *string    `json:"equipment_group"`
	Equipment_name      *string    `json:"equipment_name"`
	Equipment_type      *string    `json:"equipment_type"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Manufacturer_ba_id  *string    `json:"manufacturer_ba_id"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Purchase_date       *time.Time `json:"purchase_date"`
	Reference_num       *string    `json:"reference_num"`
	Remark              *string    `json:"remark"`
	Serial_num          *string    `json:"serial_num"`
	Source              *string    `json:"source"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
