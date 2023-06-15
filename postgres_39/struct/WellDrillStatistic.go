package dto

import (
	"time"
)

type Well_drill_statistic struct {
	Uwi                  string     `json:"uwi"`
	Period_obs_no        int        `json:"period_obs_no"`
	Statistic_type       string     `json:"statistic_type"`
	Statistic_obs_no     int        `json:"statistic_obs_no"`
	Active_ind           *string    `json:"active_ind"`
	Currency_conversion  *float64   `json:"currency_conversion"`
	Date_format_desc     *time.Time `json:"date_format_desc"`
	Description          *string    `json:"description"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Max_value            *float64   `json:"max_value"`
	Max_value_ouom       *string    `json:"max_value_ouom"`
	Max_value_uom        *string    `json:"max_value_uom"`
	Min_value            *float64   `json:"min_value"`
	Min_value_ouom       *string    `json:"min_value_ouom"`
	Min_value_uom        *string    `json:"min_value_uom"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Statistic_code       *string    `json:"statistic_code"`
	Statistic_date       *time.Time `json:"statistic_date"`
	Statistic_time       *time.Time `json:"statistic_time"`
	Statistic_timezone   *string    `json:"statistic_timezone"`
	Statistic_value      *float64   `json:"statistic_value"`
	Statistic_value_ouom *string    `json:"statistic_value_ouom"`
	Statistic_value_uom  *string    `json:"statistic_value_uom"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
