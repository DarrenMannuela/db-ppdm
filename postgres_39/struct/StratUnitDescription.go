package dto

import (
	"time"
)

type Strat_unit_description struct {
	Strat_name_set_id  string     `json:"strat_name_set_id"`
	Strat_unit_id      string     `json:"strat_unit_id"`
	Description_seq_no int        `json:"description_seq_no"`
	Active_ind         *string    `json:"active_ind"`
	Average_value      *float64   `json:"average_value"`
	Average_value_ouom *string    `json:"average_value_ouom"`
	Average_value_uom  *string    `json:"average_value_uom"`
	Date_format_desc   *time.Time `json:"date_format_desc"`
	Description        *string    `json:"description"`
	Description_date   *time.Time `json:"description_date"`
	Description_type   *string    `json:"description_type"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Max_value          *float64   `json:"max_value"`
	Max_value_ouom     *string    `json:"max_value_ouom"`
	Max_value_uom      *string    `json:"max_value_uom"`
	Min_value          *float64   `json:"min_value"`
	Min_value_ouom     *string    `json:"min_value_ouom"`
	Min_value_uom      *string    `json:"min_value_uom"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Reference_pages    *string    `json:"reference_pages"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Source_document_id *string    `json:"source_document_id"`
	Strat_unit_desc    *string    `json:"strat_unit_desc"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
