package dto

import (
	"time"
)

type Notification struct {
	Notification_id       string     `json:"notification_id"`
	Active_ind            *string    `json:"active_ind"`
	Business_associate_id *string    `json:"business_associate_id"`
	Contract_ind          *string    `json:"contract_ind"`
	Critical_date         *time.Time `json:"critical_date"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Land_right_id         *string    `json:"land_right_id"`
	Land_right_ind        *string    `json:"land_right_ind"`
	Land_right_subtype    *string    `json:"land_right_subtype"`
	Notification_type     *string    `json:"notification_type"`
	Obligation_id         *string    `json:"obligation_id"`
	Obligation_ind        *string    `json:"obligation_ind"`
	Obligation_seq_no     *int       `json:"obligation_seq_no"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Served_ind            *string    `json:"served_ind"`
	Source                *string    `json:"source"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
