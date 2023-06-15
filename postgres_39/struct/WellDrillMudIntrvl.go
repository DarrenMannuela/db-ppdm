package dto

import (
	"time"
)

type Well_drill_mud_intrvl struct {
	Uwi                   string     `json:"uwi"`
	Source                string     `json:"source"`
	Media_type            string     `json:"media_type"`
	Depth_obs_no          int        `json:"depth_obs_no"`
	Active_ind            *string    `json:"active_ind"`
	Casing_depth          *float64   `json:"casing_depth"`
	Casing_depth_ouom     *string    `json:"casing_depth_ouom"`
	Effective_date        *time.Time `json:"effective_date"`
	End_date              *time.Time `json:"end_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Max_mud_weight        *float64   `json:"max_mud_weight"`
	Max_mud_weight_ouom   *string    `json:"max_mud_weight_ouom"`
	Max_mud_weight_uom    *string    `json:"max_mud_weight_uom"`
	Max_weight_depth      *float64   `json:"max_weight_depth"`
	Max_weight_depth_ouom *string    `json:"max_weight_depth_ouom"`
	Mud_end_depth         *float64   `json:"mud_end_depth"`
	Mud_end_depth_ouom    *string    `json:"mud_end_depth_ouom"`
	Mud_start_depth       *float64   `json:"mud_start_depth"`
	Mud_start_depth_ouom  *string    `json:"mud_start_depth_ouom"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Start_date            *time.Time `json:"start_date"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
