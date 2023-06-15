package dto

import (
	"time"
)

type Well_sieve_analysis struct {
	Uwi                  string     `json:"uwi"`
	Source               string     `json:"source"`
	Analysis_obs_no      int        `json:"analysis_obs_no"`
	Active_ind           *string    `json:"active_ind"`
	Analysis_date        *time.Time `json:"analysis_date"`
	Base_depth           *float64   `json:"base_depth"`
	Base_depth_ouom      *string    `json:"base_depth_ouom"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Interval_depth       *float64   `json:"interval_depth"`
	Interval_depth_ouom  *string    `json:"interval_depth_ouom"`
	Interval_length      *float64   `json:"interval_length"`
	Interval_length_ouom *string    `json:"interval_length_ouom"`
	Laboratory           *string    `json:"laboratory"`
	Lab_file_num         *string    `json:"lab_file_num"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Top_depth            *float64   `json:"top_depth"`
	Top_depth_ouom       *string    `json:"top_depth_ouom"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
