package dto

import (
	"time"
)

type Ppdm_procedure struct {
	System_id          string     `json:"system_id"`
	Procedure_id       string     `json:"procedure_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Procedure_desc     *string    `json:"procedure_desc"`
	Procedure_name     *string    `json:"procedure_name"`
	Procedure_type     *string    `json:"procedure_type"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Table_name         *string    `json:"table_name"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
