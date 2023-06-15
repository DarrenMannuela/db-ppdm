package dto

import (
	"time"
)

type Oblig_xref struct {
	Obligation_id       string     `json:"obligation_id"`
	Obligation_seq_no   int        `json:"obligation_seq_no"`
	Obligation_id_2     string     `json:"obligation_id_2"`
	Obligation_seq_no_2 int        `json:"obligation_seq_no_2"`
	Xref_seq_no         int        `json:"xref_seq_no"`
	Active_ind          *string    `json:"active_ind"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Oblig_xref_type     *string    `json:"oblig_xref_type"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
