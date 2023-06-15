package dto

import (
	"time"
)

type Well_core_shift struct {
	Uwi                   string     `json:"uwi"`
	Source                string     `json:"source"`
	Core_id               string     `json:"core_id"`
	Shift_obs_no          int        `json:"shift_obs_no"`
	Active_ind            *string    `json:"active_ind"`
	Base_depth            *float64   `json:"base_depth"`
	Base_depth_ouom       *string    `json:"base_depth_ouom"`
	Base_shift_depth      *float64   `json:"base_shift_depth"`
	Base_shift_depth_ouom *string    `json:"base_shift_depth_ouom"`
	Core_shift_company    *string    `json:"core_shift_company"`
	Core_shift_ind        *string    `json:"core_shift_ind"`
	Core_shift_method     *string    `json:"core_shift_method"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Top_depth             *float64   `json:"top_depth"`
	Top_depth_ouom        *string    `json:"top_depth_ouom"`
	Top_shift_depth       *float64   `json:"top_shift_depth"`
	Top_shift_depth_ouom  *string    `json:"top_shift_depth_ouom"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
