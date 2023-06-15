package dto

import (
	"time"
)

type Well_test_shutoff struct {
	Uwi                 string     `json:"uwi"`
	Source              string     `json:"source"`
	Test_type           string     `json:"test_type"`
	Run_num             string     `json:"run_num"`
	Test_num            string     `json:"test_num"`
	Shutoff_obs_no      int        `json:"shutoff_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Base_depth          *float64   `json:"base_depth"`
	Base_depth_ouom     *string    `json:"base_depth_ouom"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Plugback_depth      *float64   `json:"plugback_depth"`
	Plugback_depth_ouom *string    `json:"plugback_depth_ouom"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Shutoff_type        *string    `json:"shutoff_type"`
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
