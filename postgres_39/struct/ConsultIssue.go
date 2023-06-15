package dto

import (
	"time"
)

type Consult_issue struct {
	Consult_id         string     `json:"consult_id"`
	Detail_id          string     `json:"detail_id"`
	Active_ind         *string    `json:"active_ind"`
	Detail_desc        *string    `json:"detail_desc"`
	Discussion_id      *string    `json:"discussion_id"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Issue_type         *string    `json:"issue_type"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Remark_type        *string    `json:"remark_type"`
	Resolution         *string    `json:"resolution"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
