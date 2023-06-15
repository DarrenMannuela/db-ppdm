package dto

import (
	"time"
)

type Well_drill_assembly struct {
	Uwi                 string     `json:"uwi"`
	Assembly_id         string     `json:"assembly_id"`
	Active_ind          *string    `json:"active_ind"`
	Assembly_ref_number *string    `json:"assembly_ref_number"`
	Calculated_length   *float64   `json:"calculated_length"`
	Calculated_weight   *float64   `json:"calculated_weight"`
	Effective_date      *time.Time `json:"effective_date"`
	End_time            *time.Time `json:"end_time"`
	End_timezone        *string    `json:"end_timezone"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Start_time          *time.Time `json:"start_time"`
	Start_timezone      *string    `json:"start_timezone"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
