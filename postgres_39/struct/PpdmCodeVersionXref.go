package dto

import (
	"time"
)

type Ppdm_code_version_xref struct {
	System_id          string     `json:"system_id"`
	Table_name         string     `json:"table_name"`
	Source             string     `json:"source"`
	Version_obs_no     int        `json:"version_obs_no"`
	System_id2         string     `json:"system_id_2"`
	Table_name2        string     `json:"table_name_2"`
	Source2            string     `json:"source_2"`
	Version_obs_no2    int        `json:"version_obs_no_2"`
	Active_ind         *string    `json:"active_ind"`
	Code_xref_type     *string    `json:"code_xref_type"`
	Effective_date     *time.Time `json:"effective_date"`
	Equivalence_desc   *string    `json:"equivalence_desc"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source_document_id *string    `json:"source_document_id"`
	User_rule_desc     *string    `json:"user_rule_desc"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
