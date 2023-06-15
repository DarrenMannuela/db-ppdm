package dto

import (
	"time"
)

type Reserve_class_calc struct {
	Reserve_class_id        string     `json:"reserve_class_id"`
	Formula_id              string     `json:"formula_id"`
	Calculation_seq_no      int        `json:"calculation_seq_no"`
	Active_ind              *string    `json:"active_ind"`
	Contribution_factor     *float64   `json:"contribution_factor"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Origin_reserve_class_id *string    `json:"origin_reserve_class_id"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Remark                  *string    `json:"remark"`
	Source                  *string    `json:"source"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
