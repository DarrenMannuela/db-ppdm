package dto

import (
	"time"
)

type Rm_phys_item_condition struct {
	Physical_item_id   string     `json:"physical_item_id"`
	Condition_id       string     `json:"condition_id"`
	Active_ind         *string    `json:"active_ind"`
	Condition_type     *string    `json:"condition_type"`
	Correction_method  *string    `json:"correction_method"`
	Description        *string    `json:"description"`
	Effective_date     *time.Time `json:"effective_date"`
	Error_count        *int       `json:"error_count"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Origin_seq_no      *int       `json:"origin_seq_no"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
