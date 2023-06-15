package dto

import (
	"time"
)

type Well_support_facility struct {
	Uwi                     string     `json:"uwi"`
	Sf_subtype              string     `json:"sf_subtype"`
	Support_facility_id     string     `json:"support_facility_id"`
	Sf_obs_no               int        `json:"sf_obs_no"`
	Active_ind              *string    `json:"active_ind"`
	Depth_datum             *string    `json:"depth_datum"`
	Depth_datum_elev        *float64   `json:"depth_datum_elev"`
	Depth_datum_elev_ouom   *string    `json:"depth_datum_elev_ouom"`
	Derrick_floor_elev      *float64   `json:"derrick_floor_elev"`
	Derrick_floor_elev_ouom *string    `json:"derrick_floor_elev_ouom"`
	Effective_date          *time.Time `json:"effective_date"`
	Elev_ref_datum          *string    `json:"elev_ref_datum"`
	Environment_type        *string    `json:"environment_type"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Kb_elev                 *float64   `json:"kb_elev"`
	Kb_elev_ouom            *string    `json:"kb_elev_ouom"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Remark                  *string    `json:"remark"`
	Rig_on_site_date        *time.Time `json:"rig_on_site_date"`
	Rig_release_date        *time.Time `json:"rig_release_date"`
	Rotary_table_elev       *float64   `json:"rotary_table_elev"`
	Rotary_table_elev_ouom  *string    `json:"rotary_table_elev_ouom"`
	Sf_use_type             *string    `json:"sf_use_type"`
	Source                  *string    `json:"source"`
	Subsea_elev_ref_type    *string    `json:"subsea_elev_ref_type"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
