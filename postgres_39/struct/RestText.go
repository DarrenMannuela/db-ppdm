package dto

import (
	"time"
)

type Rest_text struct {
	Restriction_id      string     `json:"restriction_id"`
	Restriction_version int        `json:"restriction_version"`
	Text_seq_no         int        `json:"text_seq_no"`
	Active_ind          *string    `json:"active_ind"`
	Body_of_text        *string    `json:"body_of_text"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
