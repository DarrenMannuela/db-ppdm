package dto

import (
	"time"
)

type Oblig_deduct_calc struct {
	Obligation_id       string     `json:"obligation_id"`
	Obligation_seq_no   int        `json:"obligation_seq_no"`
	Deduction_id        string     `json:"deduction_id"`
	Calculation_id      string     `json:"calculation_id"`
	Active_ind          *string    `json:"active_ind"`
	Calculation_formula *string    `json:"calculation_formula"`
	Calculation_method  *string    `json:"calculation_method"`
	Calculation_point   *string    `json:"calculation_point"`
	Contract_id         *string    `json:"contract_id"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Substance_id        *string    `json:"substance_id"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
