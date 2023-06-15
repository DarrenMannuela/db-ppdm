package dto

import (
	"time"
)

type Ppdm_code_version_use struct {
	System_id          string     `json:"system_id"`
	Table_name         string     `json:"table_name"`
	Source             string     `json:"source"`
	Version_obs_no     int        `json:"version_obs_no"`
	Use_obs_no         int        `json:"use_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Group_id           *string    `json:"group_id"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Preferred_ind      *string    `json:"preferred_ind"`
	Procedure_id       *string    `json:"procedure_id"`
	Remark             *string    `json:"remark"`
	Source_document_id *string    `json:"source_document_id"`
	Sw_application_id  *string    `json:"sw_application_id"`
	Use_owner_ba_id    *string    `json:"use_owner_ba_id"`
	Use_rule_desc      *string    `json:"use_rule_desc"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
