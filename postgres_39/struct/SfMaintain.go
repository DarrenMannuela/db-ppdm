package dto

import (
	"time"
)

type Sf_maintain struct {
	Sf_subtype          string     `json:"sf_subtype"`
	Support_facility_id string     `json:"support_facility_id"`
	Maintain_id         string     `json:"maintain_id"`
	Active_ind          *string    `json:"active_ind"`
	Actual_end_date     *time.Time `json:"actual_end_date"`
	Actual_start_date   *time.Time `json:"actual_start_date"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Maintain_ba_id      *string    `json:"maintain_ba_id"`
	Maintain_type       *string    `json:"maintain_type"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Schedule_end_date   *time.Time `json:"schedule_end_date"`
	Schedule_start_date *time.Time `json:"schedule_start_date"`
	Source              *string    `json:"source"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
