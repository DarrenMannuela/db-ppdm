package dto

import (
	"time"
)

type Rm_custody struct {
	Physical_item_id   string     `json:"physical_item_id"`
	Custody_id         string     `json:"custody_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Receive_by         *string    `json:"receive_by"`
	Receive_date       *time.Time `json:"receive_date"`
	Registration_num   *string    `json:"registration_num"`
	Remark             *string    `json:"remark"`
	Send_by            *string    `json:"send_by"`
	Send_date          *time.Time `json:"send_date"`
	Send_method        *string    `json:"send_method"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
