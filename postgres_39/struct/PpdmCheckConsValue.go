package dto

import (
	"time"
)

type Ppdm_check_cons_value struct {
	System_id          string     `json:"system_id"`
	Table_name         string     `json:"table_name"`
	Column_name        string     `json:"column_name"`
	Check_value        string     `json:"check_value"`
	Active_ind         *string    `json:"active_ind"`
	Check_cons_name    *string    `json:"check_cons_name"`
	Constraint_name    *string    `json:"constraint_name"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Extension_ind      *string    `json:"extension_ind"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Value_long_name    *string    `json:"value_long_name"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
