package dto

import (
	"time"
)

type Report_hier_desc struct {
	Hierarchy_type_id  string     `json:"hierarchy_type_id"`
	Level_seq_no       int        `json:"level_seq_no"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Level_name         *string    `json:"level_name"`
	Level_type         *string    `json:"level_type"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	System_id          *string    `json:"system_id"`
	Table_name         *string    `json:"table_name"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
