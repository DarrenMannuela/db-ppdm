package dto

import (
	"time"
)

type Lith_log_ba_service struct {
	Lithology_log_id      string     `json:"lithology_log_id"`
	Lith_log_source       string     `json:"lith_log_source"`
	Provided_by           string     `json:"provided_by"`
	Service_seq_no        int        `json:"service_seq_no"`
	Active_ind            *string    `json:"active_ind"`
	Ba_service_type       *string    `json:"ba_service_type"`
	Contact_ba_id         *string    `json:"contact_ba_id"`
	Contract_id           *string    `json:"contract_id"`
	Contract_provision_id *string    `json:"contract_provision_id"`
	Description           *string    `json:"description"`
	Effective_date        *time.Time `json:"effective_date"`
	End_date              *time.Time `json:"end_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Rate_schedule_id      *string    `json:"rate_schedule_id"`
	Remark                *string    `json:"remark"`
	Represented_ba_id     *string    `json:"represented_ba_id"`
	Service_quality       *string    `json:"service_quality"`
	Source                *string    `json:"source"`
	Start_date            *time.Time `json:"start_date"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
