package dto

import (
	"time"
)

type Land_right_field struct {
	Land_right_subtype string     `json:"land_right_subtype"`
	Land_right_id      string     `json:"land_right_id"`
	Field_id           string     `json:"field_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Unit_operated_ind  *string    `json:"unit_operated_ind"`
	Xref_type          *string    `json:"xref_type"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}