package dto

import (
	"time"
)

type Project_equip_role struct {
	Project_id         string     `json:"project_id"`
	Equipment_obs_no   int        `json:"equipment_obs_no"`
	Equipment_role     string     `json:"equipment_role"`
	Role_seq_no        int        `json:"role_seq_no"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
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
