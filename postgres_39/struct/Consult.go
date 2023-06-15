package dto

import (
	"time"
)

type Consult struct {
	Consult_id         string     `json:"consult_id"`
	Active_ind         *string    `json:"active_ind"`
	Complete_date      *time.Time `json:"complete_date"`
	Consult_reason     *string    `json:"consult_reason"`
	Consult_type       *string    `json:"consult_type"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Period_type        *string    `json:"period_type"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
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
