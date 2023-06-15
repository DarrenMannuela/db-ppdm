package dto

import (
	"time"
)

type Spacing_unit struct {
	Spacing_unit_id    string     `json:"spacing_unit_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Gross_size         *float64   `json:"gross_size"`
	Gross_size_ouom    *string    `json:"gross_size_ouom"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Spacing_unit_desc  *string    `json:"spacing_unit_desc"`
	Spacing_unit_name  *string    `json:"spacing_unit_name"`
	Spacing_unit_type  *string    `json:"spacing_unit_type"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
