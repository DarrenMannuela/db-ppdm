package dto

import (
	"time"
)

type Pr_str_form_completion struct {
	Uwi                    string     `json:"uwi"`
	Prod_string_source     string     `json:"prod_string_source"`
	String_id              string     `json:"string_id"`
	Pr_str_form_obs_no     int        `json:"pr_str_form_obs_no"`
	Completion_source      string     `json:"completion_source"`
	Completion_obs_no      int        `json:"completion_obs_no"`
	Form_completion_obs_no int        `json:"form_completion_obs_no"`
	Active_ind             *string    `json:"active_ind"`
	Completion_status      *string    `json:"completion_status"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Status_type            *string    `json:"status_type"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
