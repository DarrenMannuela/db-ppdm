package dto

import (
	"time"
)

type Fin_xref struct {
	Finance_id1          string     `json:"finance_id_1"`
	Finance_id2          string     `json:"finance_id_2"`
	Active_ind           *string    `json:"active_ind"`
	Distribution_percent *float64   `json:"distribution_percent"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Fin_xref_type        *string    `json:"fin_xref_type"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
