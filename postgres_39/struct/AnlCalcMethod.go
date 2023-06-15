package dto

import (
	"time"
)

type Anl_calc_method struct {
	Calculate_method_id string     `json:"calculate_method_id"`
	Active_ind          *string    `json:"active_ind"`
	Calculate_formula   *string    `json:"calculate_formula"`
	Calc_set_id         *string    `json:"calc_set_id"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Formula_type        *string    `json:"formula_type"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Preferred_ind       *string    `json:"preferred_ind"`
	Preferred_name      *string    `json:"preferred_name"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Source_document_id  *string    `json:"source_document_id"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
