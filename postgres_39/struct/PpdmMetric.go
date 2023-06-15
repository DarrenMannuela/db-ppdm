package dto

import (
	"time"
)

type Ppdm_metric struct {
	Metric_id             string     `json:"metric_id"`
	Active_ind            *string    `json:"active_ind"`
	Effective_date        *time.Time `json:"effective_date"`
	End_date              *time.Time `json:"end_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Metric_name           *string    `json:"metric_name"`
	Metric_procedure      *string    `json:"metric_procedure"`
	Metric_type           *string    `json:"metric_type"`
	Organization_id       *string    `json:"organization_id"`
	Organization_seq_no   *int       `json:"organization_seq_no"`
	Owner_ba_id           *string    `json:"owner_ba_id"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Procedure_id          *string    `json:"procedure_id"`
	Procedure_system_id   *string    `json:"procedure_system_id"`
	Projected_final_count *int       `json:"projected_final_count"`
	Purpose_desc          *string    `json:"purpose_desc"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Start_date            *time.Time `json:"start_date"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
