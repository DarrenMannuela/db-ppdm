package dto

import (
	"time"
)

type Seis_point_flow struct {
	Seis_set_subtype   string     `json:"seis_set_subtype"`
	Seis_set_id        string     `json:"seis_set_id"`
	Seis_point_id      string     `json:"seis_point_id"`
	Flow_id            string     `json:"flow_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Flow_depth         *float64   `json:"flow_depth"`
	Flow_depth_ouom    *string    `json:"flow_depth_ouom"`
	Flow_duration      *float64   `json:"flow_duration"`
	Flow_duration_uom  *string    `json:"flow_duration_uom"`
	Flow_volume        *float64   `json:"flow_volume"`
	Flow_volume_ouom   *string    `json:"flow_volume_ouom"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Remedial_ba_id     *string    `json:"remedial_ba_id"`
	Remedial_date      *time.Time `json:"remedial_date"`
	Remedial_ind       *string    `json:"remedial_ind"`
	Report_date        *time.Time `json:"report_date"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
