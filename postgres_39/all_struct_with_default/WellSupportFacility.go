package dto

type Well_support_facility struct {
	Uwi                     string   `json:"uwi" default:"source"`
	Sf_subtype              string   `json:"sf_subtype" default:"source"`
	Support_facility_id     string   `json:"support_facility_id" default:"source"`
	Sf_obs_no               int      `json:"sf_obs_no" default:"1"`
	Active_ind              *string  `json:"active_ind" default:""`
	Depth_datum             *string  `json:"depth_datum" default:""`
	Depth_datum_elev        *float64 `json:"depth_datum_elev" default:""`
	Depth_datum_elev_ouom   *string  `json:"depth_datum_elev_ouom" default:""`
	Derrick_floor_elev      *float64 `json:"derrick_floor_elev" default:""`
	Derrick_floor_elev_ouom *string  `json:"derrick_floor_elev_ouom" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	Elev_ref_datum          *string  `json:"elev_ref_datum" default:""`
	Environment_type        *string  `json:"environment_type" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Kb_elev                 *float64 `json:"kb_elev" default:""`
	Kb_elev_ouom            *string  `json:"kb_elev_ouom" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Remark                  *string  `json:"remark" default:""`
	Rig_on_site_date        *string  `json:"rig_on_site_date" default:""`
	Rig_release_date        *string  `json:"rig_release_date" default:""`
	Rotary_table_elev       *float64 `json:"rotary_table_elev" default:""`
	Rotary_table_elev_ouom  *string  `json:"rotary_table_elev_ouom" default:""`
	Sf_use_type             *string  `json:"sf_use_type" default:""`
	Source                  *string  `json:"source" default:""`
	Subsea_elev_ref_type    *string  `json:"subsea_elev_ref_type" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
