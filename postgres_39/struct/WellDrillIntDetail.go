package dto

import (
	"time"
)

type Well_drill_int_detail struct {
	Uwi                 string     `json:"uwi"`
	Bit_source          string     `json:"bit_source"`
	Bit_interval_obs_no int        `json:"bit_interval_obs_no"`
	Bit_detail_type     string     `json:"bit_detail_type"`
	Detail_obs_no       int        `json:"detail_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Average_value       *float64   `json:"average_value"`
	Average_value_ouom  *string    `json:"average_value_ouom"`
	Average_value_uom   *string    `json:"average_value_uom"`
	Base_depth          *float64   `json:"base_depth"`
	Base_depth_ouom     *string    `json:"base_depth_ouom"`
	Bit_detail_code     *string    `json:"bit_detail_code"`
	Date_format_desc    *time.Time `json:"date_format_desc"`
	Detail_desc         *string    `json:"detail_desc"`
	Effective_date      *time.Time `json:"effective_date"`
	End_date            *time.Time `json:"end_date"`
	End_time            *time.Time `json:"end_time"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Max_value           *float64   `json:"max_value"`
	Max_value_ouom      *string    `json:"max_value_ouom"`
	Max_value_uom       *string    `json:"max_value_uom"`
	Min_value           *float64   `json:"min_value"`
	Min_value_ouom      *string    `json:"min_value_ouom"`
	Min_value_uom       *string    `json:"min_value_uom"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Start_date          *time.Time `json:"start_date"`
	Start_time          *time.Time `json:"start_time"`
	Timezone            *string    `json:"timezone"`
	Top_depth           *float64   `json:"top_depth"`
	Top_depth_ouom      *string    `json:"top_depth_ouom"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
