package dto

import (
	"time"
)

type Ba_consortium_service struct {
	Consortium_ba_id   string     `json:"consortium_ba_id"`
	Provided_by        string     `json:"provided_by"`
	Service_seq_no     int        `json:"service_seq_no"`
	Active_ind         *string    `json:"active_ind"`
	Contact_ba_id      *string    `json:"contact_ba_id"`
	Contract_id        *string    `json:"contract_id"`
	Description        *string    `json:"description"`
	Effective_date     *time.Time `json:"effective_date"`
	End_date           *time.Time `json:"end_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Provision_id       *string    `json:"provision_id"`
	Rate_schedule_id   *string    `json:"rate_schedule_id"`
	Remark             *string    `json:"remark"`
	Represented_ba_id  *string    `json:"represented_ba_id"`
	Service_quality    *string    `json:"service_quality"`
	Service_type       *string    `json:"service_type"`
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
