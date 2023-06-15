package dto

import (
	"time"
)

type Project_equipment struct {
	Project_id         string     `json:"project_id"`
	Equipment_obs_no   int        `json:"equipment_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Catalogue_equip_id *string    `json:"catalogue_equip_id"`
	Effective_date     *time.Time `json:"effective_date"`
	Equipment_id       *string    `json:"equipment_id"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Operated_by_ba_id  *string    `json:"operated_by_ba_id"`
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
