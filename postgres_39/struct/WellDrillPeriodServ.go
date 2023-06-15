package dto

import (
	"time"
)

type Well_drill_period_serv struct {
	Uwi                string     `json:"uwi"`
	Period_obs_no      int        `json:"period_obs_no"`
	Provided_by        string     `json:"provided_by"`
	Service_seq_no     int        `json:"service_seq_no"`
	Segment_obs_no     int        `json:"segment_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Description        *string    `json:"description"`
	Effective_date     *time.Time `json:"effective_date"`
	End_time           *time.Time `json:"end_time"`
	End_timezone       *string    `json:"end_timezone"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Metric_code        *string    `json:"metric_code"`
	Metric_max_value   *float64   `json:"metric_max_value"`
	Metric_min_value   *float64   `json:"metric_min_value"`
	Metric_value       *float64   `json:"metric_value"`
	Metric_value_ouom  *string    `json:"metric_value_ouom"`
	Metric_value_uom   *string    `json:"metric_value_uom"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Service_desc       *string    `json:"service_desc"`
	Service_metric     *string    `json:"service_metric"`
	Source             *string    `json:"source"`
	Start_time         *time.Time `json:"start_time"`
	Start_timezone     *string    `json:"start_timezone"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
