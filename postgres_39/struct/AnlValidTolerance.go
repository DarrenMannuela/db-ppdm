package dto

import (
	"time"
)

type Anl_valid_tolerance struct {
	Method_set_id      string     `json:"method_set_id"`
	Method_id          string     `json:"method_id"`
	Equip_obs_no       int        `json:"equip_obs_no"`
	Tolerance_obs_no   int        `json:"tolerance_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Max_tolerance      *float64   `json:"max_tolerance"`
	Max_tolerance_ouom *string    `json:"max_tolerance_ouom"`
	Max_tolerance_rep  *float64   `json:"max_tolerance_rep"`
	Max_tolerance_uom  *string    `json:"max_tolerance_uom"`
	Min_tolerance      *float64   `json:"min_tolerance"`
	Min_tolerance_ouom *string    `json:"min_tolerance_ouom"`
	Min_tolerance_rep  *string    `json:"min_tolerance_rep"`
	Min_tolerance_uom  *string    `json:"min_tolerance_uom"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Substance_id       *string    `json:"substance_id"`
	Tolerance_desc     *string    `json:"tolerance_desc"`
	Tolerance_type     *string    `json:"tolerance_type"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
