package dto

import (
	"time"
)

type Resent_vol_revision struct {
	Resent_id            string     `json:"resent_id"`
	Reserve_class_id     string     `json:"reserve_class_id"`
	Revision_id          string     `json:"revision_id"`
	Revision_obs_no      int        `json:"revision_obs_no"`
	Active_ind           *string    `json:"active_ind"`
	Analyst_ba_id        *string    `json:"analyst_ba_id"`
	Approved_by_ba_id    *string    `json:"approved_by_ba_id"`
	Approved_date        *time.Time `json:"approved_date"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Gross_revision_ind   *string    `json:"gross_revision_ind"`
	Interest_set_id      *string    `json:"interest_set_id"`
	Interest_set_seq_no  *int       `json:"interest_set_seq_no"`
	Net_revision_ind     *string    `json:"net_revision_ind"`
	New_volume           *float64   `json:"new_volume"`
	New_volume_ouom      *string    `json:"new_volume_ouom"`
	New_volume_uom       *string    `json:"new_volume_uom"`
	Partner_ba_id        *string    `json:"partner_ba_id"`
	Partner_obs_no       *int       `json:"partner_obs_no"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Product_type         *string    `json:"product_type"`
	Project_id           *string    `json:"project_id"`
	Remark               *string    `json:"remark"`
	Report_ind           *string    `json:"report_ind"`
	Revision_category_id *string    `json:"revision_category_id"`
	Revision_date        *time.Time `json:"revision_date"`
	Revision_date_desc   *string    `json:"revision_date_desc"`
	Revision_method      *string    `json:"revision_method"`
	Revision_volume      *float64   `json:"revision_volume"`
	Revision_volume_ouom *string    `json:"revision_volume_ouom"`
	Revision_volume_uom  *string    `json:"revision_volume_uom"`
	Source               *string    `json:"source"`
	Summary_id           *string    `json:"summary_id"`
	Summary_obs_no       *int       `json:"summary_obs_no"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
