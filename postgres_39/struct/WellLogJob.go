package dto

import (
	"time"
)

type Well_log_job struct {
	Uwi                    string     `json:"uwi"`
	Source                 string     `json:"source"`
	Job_id                 string     `json:"job_id"`
	Active_ind             *string    `json:"active_ind"`
	Area_id                *string    `json:"area_id"`
	Area_type              *string    `json:"area_type"`
	Casing_shoe_depth      *float64   `json:"casing_shoe_depth"`
	Casing_shoe_depth_ouom *string    `json:"casing_shoe_depth_ouom"`
	Drilling_md            *float64   `json:"drilling_md"`
	Drilling_md_ouom       *string    `json:"drilling_md_ouom"`
	Effective_date         *time.Time `json:"effective_date"`
	End_date               *time.Time `json:"end_date"`
	Engineer               *string    `json:"engineer"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Logging_company        *string    `json:"logging_company"`
	Logging_unit           *string    `json:"logging_unit"`
	Observer               *string    `json:"observer"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Start_date             *time.Time `json:"start_date"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
