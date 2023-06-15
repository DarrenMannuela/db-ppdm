package dto

import (
	"time"
)

type Ppdm_map_load struct {
	Map_id             string     `json:"map_id"`
	Load_process_id    string     `json:"load_process_id"`
	Active_ind         *string    `json:"active_ind"`
	Delete_allowed_ind *string    `json:"delete_allowed_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	End_date           *time.Time `json:"end_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Insert_allowed_ind *string    `json:"insert_allowed_ind"`
	Ppdm_group_id      *string    `json:"ppdm_group_id"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Procedure_id       *string    `json:"procedure_id"`
	Procedure_name     *string    `json:"procedure_name"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Source_system_id   *string    `json:"source_system_id"`
	Start_date         *time.Time `json:"start_date"`
	Sw_application_id  *string    `json:"sw_application_id"`
	Target_system_id   *string    `json:"target_system_id"`
	Update_allowed_ind *string    `json:"update_allowed_ind"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
