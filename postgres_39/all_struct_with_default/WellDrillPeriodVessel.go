package dto

type Well_drill_period_vessel struct {
	Uwi                string   `json:"uwi" default:"source"`
	Period_obs_no      int      `json:"period_obs_no" default:"1"`
	Sf_subtype         string   `json:"sf_subtype" default:"source"`
	Vessel_id          string   `json:"vessel_id" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Heading            *float64 `json:"heading" default:""`
	Heading_north_type *string  `json:"heading_north_type" default:""`
	Passengers_off     *int     `json:"passengers_off" default:""`
	Passengers_on      *int     `json:"passengers_on" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Riser_angle        *float64 `json:"riser_angle" default:""`
	Riser_tension      *float64 `json:"riser_tension" default:""`
	Riser_tension_ouom *string  `json:"riser_tension_ouom" default:""`
	Riser_tension_uom  *string  `json:"riser_tension_uom" default:""`
	Source             *string  `json:"source" default:""`
	Vessel_role        *string  `json:"vessel_role" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
