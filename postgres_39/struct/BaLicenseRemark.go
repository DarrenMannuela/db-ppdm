package dto

import (
	"time"
)

type Ba_license_remark struct {
	Licensee_ba_id     string     `json:"licensee_ba_id"`
	License_id         string     `json:"license_id"`
	Remark_id          string     `json:"remark_id"`
	Remark_seq_no      int        `json:"remark_seq_no"`
	Active_ind         *string    `json:"active_ind"`
	Author             *string    `json:"author"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Remark_date        *time.Time `json:"remark_date"`
	Remark_type        *string    `json:"remark_type"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
