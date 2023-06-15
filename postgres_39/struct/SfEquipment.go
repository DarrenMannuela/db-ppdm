package dto

import (
	"time"
)

type Sf_equipment struct {
	Sf_subtype          string     `json:"sf_subtype"`
	Support_facility_id string     `json:"support_facility_id"`
	Component_id        string     `json:"component_id"`
	Active_ind          *string    `json:"active_ind"`
	Catalogue_equip_id  *string    `json:"catalogue_equip_id"`
	Component_count     *int       `json:"component_count"`
	Description         *string    `json:"description"`
	Effective_date      *time.Time `json:"effective_date"`
	Equipment_id        *string    `json:"equipment_id"`
	Equipment_name      *string    `json:"equipment_name"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Install_date        *time.Time `json:"install_date"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Purchase_date       *time.Time `json:"purchase_date"`
	Reference_num       *string    `json:"reference_num"`
	Remark              *string    `json:"remark"`
	Remove_date         *time.Time `json:"remove_date"`
	Source              *string    `json:"source"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
