package dto

import (
	"time"
)

type Pden_vol_disposition struct {
	Pden_subtype           string     `json:"pden_subtype"`
	Pden_id                string     `json:"pden_id"`
	Pden_source            string     `json:"pden_source"`
	Volume_method          string     `json:"volume_method"`
	Activity_type          string     `json:"activity_type"`
	Product_type           string     `json:"product_type"`
	Period_type            string     `json:"period_type"`
	Disposition_obs_no     int        `json:"disposition_obs_no"`
	Amendment_seq_no       int        `json:"amendment_seq_no"`
	Active_ind             *string    `json:"active_ind"`
	Amend_reason           *string    `json:"amend_reason"`
	Business_associate_id  *string    `json:"business_associate_id"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Posted_date            *time.Time `json:"posted_date"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Receiving_pden_id      *string    `json:"receiving_pden_id"`
	Receiving_pden_subtype *string    `json:"receiving_pden_subtype"`
	Remark                 *string    `json:"remark"`
	Report_ind             *string    `json:"report_ind"`
	Source                 *string    `json:"source"`
	Volume                 *float64   `json:"volume"`
	Volume_end_date        *time.Time `json:"volume_end_date"`
	Volume_end_date_desc   *string    `json:"volume_end_date_desc"`
	Volume_ouom            *string    `json:"volume_ouom"`
	Volume_period          *float64   `json:"volume_period"`
	Volume_period_ouom     *string    `json:"volume_period_ouom"`
	Volume_quality         *float64   `json:"volume_quality"`
	Volume_quality_ouom    *string    `json:"volume_quality_ouom"`
	Volume_start_date      *time.Time `json:"volume_start_date"`
	Volume_uom             *string    `json:"volume_uom"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
