package dto

import (
	"time"
)

type Well_test_recovery struct {
	Uwi                     string     `json:"uwi"`
	Source                  string     `json:"source"`
	Test_type               string     `json:"test_type"`
	Run_num                 string     `json:"run_num"`
	Test_num                string     `json:"test_num"`
	Recovery_obs_no         int        `json:"recovery_obs_no"`
	Active_ind              *string    `json:"active_ind"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Multiple_test_ind       *string    `json:"multiple_test_ind"`
	Period_obs_no           *int       `json:"period_obs_no"`
	Period_type             *string    `json:"period_type"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Recovery_amount         *float64   `json:"recovery_amount"`
	Recovery_amount_ouom    *string    `json:"recovery_amount_ouom"`
	Recovery_amount_percent *float64   `json:"recovery_amount_percent"`
	Recovery_amount_uom     *string    `json:"recovery_amount_uom"`
	Recovery_desc           *string    `json:"recovery_desc"`
	Recovery_method         *string    `json:"recovery_method"`
	Recovery_show_type      *string    `json:"recovery_show_type"`
	Recovery_type           *string    `json:"recovery_type"`
	Remark                  *string    `json:"remark"`
	Reverse_circulation_ind *string    `json:"reverse_circulation_ind"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
