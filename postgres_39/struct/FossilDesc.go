package dto

import (
	"time"
)

type Fossil_desc struct {
	Fossil_id          string     `json:"fossil_id"`
	Fossil_desc_id     string     `json:"fossil_desc_id"`
	Active_ind         *string    `json:"active_ind"`
	Date_format_desc   *time.Time `json:"date_format_desc"`
	Description        *string    `json:"description"`
	Description_code   *string    `json:"description_code"`
	Description_type   *string    `json:"description_type"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Max_value          *float64   `json:"max_value"`
	Max_value_uom      *string    `json:"max_value_uom"`
	Min_value          *float64   `json:"min_value"`
	Min_value_uom      *string    `json:"min_value_uom"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Source_document_id *string    `json:"source_document_id"`
	Value_ouom         *string    `json:"value_ouom"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
