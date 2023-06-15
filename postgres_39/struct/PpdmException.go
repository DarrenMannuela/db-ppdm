package dto

import (
	"time"
)

type Ppdm_exception struct {
	Pe_id                string     `json:"pe_id"`
	Active_ind           *string    `json:"active_ind"`
	Constraint_full_name *string    `json:"constraint_full_name"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Owner                *string    `json:"owner"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Row_id               *string    `json:"row_id"`
	Source               *string    `json:"source"`
	System_id            *string    `json:"system_id"`
	Table_name           *string    `json:"table_name"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
