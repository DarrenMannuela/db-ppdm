package dto

import (
	"time"
)

type Well_log_pass struct {
	Uwi                string     `json:"uwi"`
	Source             string     `json:"source"`
	Job_id             string     `json:"job_id"`
	Trip_obs_no        int        `json:"trip_obs_no"`
	Log_tool_pass_no   int        `json:"log_tool_pass_no"`
	Active_ind         *string    `json:"active_ind"`
	Base_depth         *float64   `json:"base_depth"`
	Base_depth_ouom    *string    `json:"base_depth_ouom"`
	Base_log_ind       *string    `json:"base_log_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	End_time           *time.Time `json:"end_time"`
	End_timezone       *string    `json:"end_timezone"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Start_time         *time.Time `json:"start_time"`
	Start_timezone     *string    `json:"start_timezone"`
	Top_depth          *float64   `json:"top_depth"`
	Top_depth_ouom     *string    `json:"top_depth_ouom"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
