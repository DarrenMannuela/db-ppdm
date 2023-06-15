package dto

import (
	"time"
)

type Ppdm_object_status struct {
	Status_id           string     `json:"status_id"`
	Status_obs_no       int        `json:"status_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Code_version_obs_no *int       `json:"code_version_obs_no"`
	Code_version_source *string    `json:"code_version_source"`
	Column_name         *string    `json:"column_name"`
	Constraint_name     *string    `json:"constraint_name"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Index_id            *string    `json:"index_id"`
	Object_name         *string    `json:"object_name"`
	Object_status       *string    `json:"object_status"`
	Object_type         *string    `json:"object_type"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Procedure_id        *string    `json:"procedure_id"`
	Property_set_id     *string    `json:"property_set_id"`
	Remark              *string    `json:"remark"`
	Rule_id             *string    `json:"rule_id"`
	Source              *string    `json:"source"`
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
