package dto

import (
	"time"
)

type Ppdm_map_load_error struct {
	Map_id                string     `json:"map_id"`
	Load_process_id       string     `json:"load_process_id"`
	Error_id              string     `json:"error_id"`
	Active_ind            *string    `json:"active_ind"`
	Delete_error_ind      *string    `json:"delete_error_ind"`
	Effective_date        *time.Time `json:"effective_date"`
	Error_cause_desc      *string    `json:"error_cause_desc"`
	Error_cause_type      *string    `json:"error_cause_type"`
	Error_code            *string    `json:"error_code"`
	Error_date            *time.Time `json:"error_date"`
	Error_message         *string    `json:"error_message"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Insert_error_ind      *string    `json:"insert_error_ind"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Ref_map_detail_obs_no *int       `json:"ref_map_detail_obs_no"`
	Remark                *string    `json:"remark"`
	Resolution_desc       *string    `json:"resolution_desc"`
	Resolution_type       *string    `json:"resolution_type"`
	Source                *string    `json:"source"`
	Update_error_ind      *string    `json:"update_error_ind"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
