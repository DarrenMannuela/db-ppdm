package dto

import (
	"time"
)

type Contest struct {
	Contest_id         string     `json:"contest_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	End_date           *time.Time `json:"end_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Lr_contest_type    *string    `json:"lr_contest_type"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Reason             *string    `json:"reason"`
	Remark             *string    `json:"remark"`
	Resolution_date    *time.Time `json:"resolution_date"`
	Resolution_method  *string    `json:"resolution_method"`
	Resolution_remark  *string    `json:"resolution_remark"`
	Source             *string    `json:"source"`
	Start_date         *time.Time `json:"start_date"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
