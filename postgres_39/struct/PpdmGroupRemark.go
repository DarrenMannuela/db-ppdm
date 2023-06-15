package dto

import (
	"time"
)

type Ppdm_group_remark struct {
	Group_id            string     `json:"group_id"`
	Remark_obs_no       int        `json:"remark_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Group_remark        *string    `json:"group_remark"`
	Made_by_ba_id       *string    `json:"made_by_ba_id"`
	Organization_id     *string    `json:"organization_id"`
	Organization_seq_no *int       `json:"organization_seq_no"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Remark_type         *string    `json:"remark_type"`
	Source              *string    `json:"source"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
