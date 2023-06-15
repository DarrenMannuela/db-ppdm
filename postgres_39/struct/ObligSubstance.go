package dto

import (
	"time"
)

type Oblig_substance struct {
	Obligation_id      string     `json:"obligation_id"`
	Obligation_seq_no  int        `json:"obligation_seq_no"`
	Substance_id       string     `json:"substance_id"`
	Substance_seq_no   int        `json:"substance_seq_no"`
	Active_ind         *string    `json:"active_ind"`
	Contract_id        *string    `json:"contract_id"`
	Effective_date     *time.Time `json:"effective_date"`
	Excluded_ind       *string    `json:"excluded_ind"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Included_ind       *string    `json:"included_ind"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
