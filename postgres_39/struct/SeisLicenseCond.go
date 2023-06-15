package dto

import (
	"time"
)

type Seis_license_cond struct {
	Seis_set_subtype    string     `json:"seis_set_subtype"`
	Seis_set_id         string     `json:"seis_set_id"`
	License_id          string     `json:"license_id"`
	Condition_id        string     `json:"condition_id"`
	Active_ind          *string    `json:"active_ind"`
	Condition_code      *string    `json:"condition_code"`
	Condition_desc      *string    `json:"condition_desc"`
	Condition_type      *string    `json:"condition_type"`
	Condition_value     *float64   `json:"condition_value"`
	Condition_value_uom *string    `json:"condition_value_uom"`
	Contact_ba_id       *string    `json:"contact_ba_id"`
	Date_format_desc    *time.Time `json:"date_format_desc"`
	Description         *string    `json:"description"`
	Due_condition       *string    `json:"due_condition"`
	Due_date            *time.Time `json:"due_date"`
	Due_frequency       *string    `json:"due_frequency"`
	Due_term            *float64   `json:"due_term"`
	Due_term_uom        *string    `json:"due_term_uom"`
	Effective_date      *time.Time `json:"effective_date"`
	Exempt_ind          *string    `json:"exempt_ind"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Fulfilled_by_ba_id  *string    `json:"fulfilled_by_ba_id"`
	Fulfilled_date      *time.Time `json:"fulfilled_date"`
	Fulfilled_ind       *string    `json:"fulfilled_ind"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Restriction_id      *string    `json:"restriction_id"`
	Restriction_version *int       `json:"restriction_version"`
	Source              *string    `json:"source"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
