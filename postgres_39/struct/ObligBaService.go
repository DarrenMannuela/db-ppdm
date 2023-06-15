package dto

import (
	"time"
)

type Oblig_ba_service struct {
	Obligation_id        string     `json:"obligation_id"`
	Obligation_seq_no    int        `json:"obligation_seq_no"`
	Provided_by          string     `json:"provided_by"`
	Oblig_service_seq_no int        `json:"oblig_service_seq_no"`
	Active_ind           *string    `json:"active_ind"`
	Contact_ba_id        *string    `json:"contact_ba_id"`
	Contract_id          *string    `json:"contract_id"`
	Description          *string    `json:"description"`
	Effective_date       *time.Time `json:"effective_date"`
	End_date             *time.Time `json:"end_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Provided_for         *string    `json:"provided_for"`
	Provision_id         *string    `json:"provision_id"`
	Rate_schedule_id     *string    `json:"rate_schedule_id"`
	Remark               *string    `json:"remark"`
	Service_quality      *string    `json:"service_quality"`
	Service_type         *string    `json:"service_type"`
	Source               *string    `json:"source"`
	Start_date           *time.Time `json:"start_date"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
