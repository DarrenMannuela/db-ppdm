package dto

import (
	"time"
)

type Pden_pr_str_allowable struct {
	Pden_subtype         string     `json:"pden_subtype"`
	Pden_id              string     `json:"pden_id"`
	Pden_source          string     `json:"pden_source"`
	Uwi                  string     `json:"uwi"`
	String_source        string     `json:"string_source"`
	String_id            string     `json:"string_id"`
	Pden_prs_xref_seq_no int        `json:"pden_prs_xref_seq_no"`
	Allowable_obs_no     int        `json:"allowable_obs_no"`
	Active_ind           *string    `json:"active_ind"`
	Allowable            *float64   `json:"allowable"`
	Allowable_date       *time.Time `json:"allowable_date"`
	Allowable_date_desc  *string    `json:"allowable_date_desc"`
	Allowable_ouom       *string    `json:"allowable_ouom"`
	Allowable_uom        *string    `json:"allowable_uom"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Period_type          *string    `json:"period_type"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Product_type         *string    `json:"product_type"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
