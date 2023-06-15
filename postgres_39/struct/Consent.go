package dto

import (
	"time"
)

type Consent struct {
	Consent_id          string     `json:"consent_id"`
	Active_ind          *string    `json:"active_ind"`
	Consent_desc        *string    `json:"consent_desc"`
	Consent_method_desc *string    `json:"consent_method_desc"`
	Consent_type        *string    `json:"consent_type"`
	Current_status      *string    `json:"current_status"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Receive_date        *time.Time `json:"receive_date"`
	Remark              *string    `json:"remark"`
	Request_date        *time.Time `json:"request_date"`
	Source              *string    `json:"source"`
	Status_remark       *string    `json:"status_remark"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
