package dto

import (
	"time"
)

type Well_zone_intrvl_value struct {
	Uwi                string     `json:"uwi"`
	Interval_source    string     `json:"interval_source"`
	Interval_id        string     `json:"interval_id"`
	Zone_id            string     `json:"zone_id"`
	Zone_source        string     `json:"zone_source"`
	Zone_value_obs_no  int        `json:"zone_value_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Date_format_desc   *time.Time `json:"date_format_desc"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Parameter          *string    `json:"parameter"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Value_type         *string    `json:"value_type"`
	Zone_value         *float64   `json:"zone_value"`
	Zone_value_ouom    *string    `json:"zone_value_ouom"`
	Zone_value_uom     *string    `json:"zone_value_uom"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
