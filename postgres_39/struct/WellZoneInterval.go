package dto

import (
	"time"
)

type Well_zone_interval struct {
	Uwi                string     `json:"uwi"`
	Source             string     `json:"source"`
	Interval_id        string     `json:"interval_id"`
	Zone_id            string     `json:"zone_id"`
	Zone_source        string     `json:"zone_source"`
	Active_ind         *string    `json:"active_ind"`
	Base_md            *float64   `json:"base_md"`
	Base_md_ouom       *string    `json:"base_md_ouom"`
	Base_tvd           *float64   `json:"base_tvd"`
	Base_tvd_ouom      *string    `json:"base_tvd_ouom"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Top_delta_x        *float64   `json:"top_delta_x"`
	Top_delta_x_ouom   *string    `json:"top_delta_x_ouom"`
	Top_delta_y        *float64   `json:"top_delta_y"`
	Top_delta_y_ouom   *string    `json:"top_delta_y_ouom"`
	Top_md             *float64   `json:"top_md"`
	Top_md_ouom        *string    `json:"top_md_ouom"`
	Top_tvd            *float64   `json:"top_tvd"`
	Top_tvd_ouom       *string    `json:"top_tvd_ouom"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
