package dto

import (
	"time"
)

type Anl_valid_equip struct {
	Method_set_id      string     `json:"method_set_id"`
	Method_id          string     `json:"method_id"`
	Equip_obs_no       int        `json:"equip_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Catalogue_equip_id *string    `json:"catalogue_equip_id"`
	Confidence_type    *string    `json:"confidence_type"`
	Description        *string    `json:"description"`
	Effective_date     *time.Time `json:"effective_date"`
	Equip_id           *string    `json:"equip_id"`
	Equip_role         *string    `json:"equip_role"`
	Equip_setting      *string    `json:"equip_setting"`
	Expiry_date        *time.Time `json:"expiry_date"`
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
