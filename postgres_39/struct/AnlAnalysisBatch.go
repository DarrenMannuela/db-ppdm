package dto

import (
	"time"
)

type Anl_analysis_batch struct {
	Batch_id              string     `json:"batch_id"`
	Active_ind            *string    `json:"active_ind"`
	Batch_desc            *string    `json:"batch_desc"`
	Batch_name            *string    `json:"batch_name"`
	Batch_ref_num         *string    `json:"batch_ref_num"`
	Create_date           *time.Time `json:"create_date"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Sample_count          *int       `json:"sample_count"`
	Source                *string    `json:"source"`
	Standard_sample_count *int       `json:"standard_sample_count"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
