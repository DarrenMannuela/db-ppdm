package dto

import (
	"time"
)

type Work_order struct {
	Work_order_id      string     `json:"work_order_id"`
	Active_ind         *string    `json:"active_ind"`
	Complete_date      *time.Time `json:"complete_date"`
	Due_date           *time.Time `json:"due_date"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Final_billing_date *time.Time `json:"final_billing_date"`
	Instructions       *string    `json:"instructions"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Request_date       *time.Time `json:"request_date"`
	Rush_ind           *string    `json:"rush_ind"`
	Source             *string    `json:"source"`
	Work_order_number  *string    `json:"work_order_number"`
	Work_order_type    *string    `json:"work_order_type"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
