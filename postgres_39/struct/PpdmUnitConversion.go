package dto

import (
	"time"
)

type Ppdm_unit_conversion struct {
	From_uom_id          string     `json:"from_uom_id"`
	To_uom_id            string     `json:"to_uom_id"`
	Active_ind           *string    `json:"active_ind"`
	Effective_date       *time.Time `json:"effective_date"`
	Exact_conversion_ind *string    `json:"exact_conversion_ind"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Factor_denominator   *float64   `json:"factor_denominator"`
	Factor_numerator     *float64   `json:"factor_numerator"`
	Post_offset          *float64   `json:"post_offset"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Pre_offset           *float64   `json:"pre_offset"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Source_document_id   *string    `json:"source_document_id"`
	Unit_expression      *float64   `json:"unit_expression"`
	Unit_quantity_type   *string    `json:"unit_quantity_type"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
