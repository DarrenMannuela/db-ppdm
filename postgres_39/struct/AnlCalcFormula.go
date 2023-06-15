package dto

import (
	"time"
)

type Anl_calc_formula struct {
	Calculate_method_id string     `json:"calculate_method_id"`
	Formula_part_obs_no int        `json:"formula_part_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Assigned_variable   *string    `json:"assigned_variable"`
	Column_name         *string    `json:"column_name"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Substance_id        *string    `json:"substance_id"`
	System_id           *string    `json:"system_id"`
	Table_name          *string    `json:"table_name"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
