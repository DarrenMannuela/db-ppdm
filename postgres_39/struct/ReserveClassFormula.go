package dto

import (
	"time"
)

type Reserve_class_formula struct {
	Reserve_class_id   string     `json:"reserve_class_id"`
	Formula_id         string     `json:"formula_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Formula            *string    `json:"formula"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Use_rule           *string    `json:"use_rule"`
	Use_seq_no         *int       `json:"use_seq_no"`
	Use_type           *string    `json:"use_type"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
