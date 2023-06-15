package dto

import (
	"time"
)

type Ppdm_table struct {
	System_id          string     `json:"system_id"`
	Table_name         string     `json:"table_name"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Extension_ind      *string    `json:"extension_ind"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Primary_key_name   *string    `json:"primary_key_name"`
	Remark             *string    `json:"remark"`
	Row_count          *float64   `json:"row_count"`
	Row_count_date     *time.Time `json:"row_count_date"`
	Source             *string    `json:"source"`
	Table_comment      *string    `json:"table_comment"`
	Table_type         *string    `json:"table_type"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
