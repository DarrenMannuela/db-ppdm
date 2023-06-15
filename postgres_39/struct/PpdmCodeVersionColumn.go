package dto

import (
	"time"
)

type Ppdm_code_version_column struct {
	System_id          string     `json:"system_id"`
	Table_name         string     `json:"table_name"`
	Source             string     `json:"source"`
	Version_obs_no     int        `json:"version_obs_no"`
	Column_name        string     `json:"column_name"`
	Abbreviation       *string    `json:"abbreviation"`
	Active_ind         *string    `json:"active_ind"`
	Definition         *string    `json:"definition"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Language           *string    `json:"language"`
	Long_name          *string    `json:"long_name"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Primary_key        *string    `json:"primary_key"`
	Remark             *string    `json:"remark"`
	Short_name         *string    `json:"short_name"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
