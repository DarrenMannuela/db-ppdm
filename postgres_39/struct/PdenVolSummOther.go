package dto

import (
	"time"
)

type Pden_vol_summ_other struct {
	Pden_subtype        string     `json:"pden_subtype"`
	Pden_id             string     `json:"pden_id"`
	Pden_source         string     `json:"pden_source"`
	Volume_method       string     `json:"volume_method"`
	Activity_type       string     `json:"activity_type"`
	Period_type         string     `json:"period_type"`
	Pden_period_id      string     `json:"pden_period_id"`
	Amendment_seq_no    int        `json:"amendment_seq_no"`
	Product_type        string     `json:"product_type"`
	Active_ind          *string    `json:"active_ind"`
	Cum_volume          *float64   `json:"cum_volume"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Gas_energy          *float64   `json:"gas_energy"`
	Gas_energy_ouom     *string    `json:"gas_energy_ouom"`
	Mass                *float64   `json:"mass"`
	Mass_ouom           *string    `json:"mass_ouom"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Report_ind          *string    `json:"report_ind"`
	Source              *string    `json:"source"`
	Volume              *float64   `json:"volume"`
	Volume_date         *time.Time `json:"volume_date"`
	Volume_date_desc    *string    `json:"volume_date_desc"`
	Volume_ouom         *string    `json:"volume_ouom"`
	Volume_quality      *float64   `json:"volume_quality"`
	Volume_quality_ouom *string    `json:"volume_quality_ouom"`
	Volume_uom          *string    `json:"volume_uom"`
	Ytd_volume          *float64   `json:"ytd_volume"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
