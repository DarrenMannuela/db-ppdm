package dto

import (
	"time"
)

type Work_order_xref struct {
	Work_order_id         string     `json:"work_order_id"`
	Work_order_xref_id    string     `json:"work_order_xref_id"`
	Active_ind            *string    `json:"active_ind"`
	Business_associate_id *string    `json:"business_associate_id"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Parent_work_order_id  *string    `json:"parent_work_order_id"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Reference_id          *string    `json:"reference_id"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Wo_xref_type          *string    `json:"wo_xref_type"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
