package dto

import (
	"time"
)

type Ppdm_group_owner struct {
	Group_id               string     `json:"group_id"`
	Owner_obs_no           int        `json:"owner_obs_no"`
	Active_ind             *string    `json:"active_ind"`
	Default_unit_system_id *string    `json:"default_unit_system_id"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Organization_id        *string    `json:"organization_id"`
	Organization_seq_no    *int       `json:"organization_seq_no"`
	Owner_ba_id            *string    `json:"owner_ba_id"`
	Owner_role             *string    `json:"owner_role"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Sw_application_id      *string    `json:"sw_application_id"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
