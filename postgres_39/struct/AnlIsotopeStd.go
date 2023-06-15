package dto

import (
	"time"
)

type Anl_isotope_std struct {
	Standard_id          string     `json:"standard_id"`
	Substance_id         string     `json:"substance_id"`
	Active_ind           *string    `json:"active_ind"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Source_document_id   *string    `json:"source_document_id"`
	Standard_name        *string    `json:"standard_name"`
	Standard_value       *float64   `json:"standard_value"`
	Standard_value_error *float64   `json:"standard_value_error"`
	Standard_value_ouom  *string    `json:"standard_value_ouom"`
	Standard_value_uom   *string    `json:"standard_value_uom"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
