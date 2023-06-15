package dto

type Seis_vessel struct {
	Sf_subtype           string   `json:"sf_subtype" default:"source"`
	Vessel_id            string   `json:"vessel_id" default:"source"`
	Vessel_config_obs_no int      `json:"vessel_config_obs_no" default:"1"`
	Acqtn_design_id      *string  `json:"acqtn_design_id" default:""`
	Active_ind           *string  `json:"active_ind" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Fath_azimuth         *float64 `json:"fath_azimuth" default:""`
	Fath_offset          *float64 `json:"fath_offset" default:""`
	Master_vessel_ind    *string  `json:"master_vessel_ind" default:""`
	Offset_ouom          *string  `json:"offset_ouom" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Reference_point      *string  `json:"reference_point" default:""`
	Remark               *string  `json:"remark" default:""`
	Shot_offset          *float64 `json:"shot_offset" default:""`
	Slave_vessel_ind     *string  `json:"slave_vessel_ind" default:""`
	Source               *string  `json:"source" default:""`
	Streamer_far_offset  *float64 `json:"streamer_far_offset" default:""`
	Streamer_near_offset *float64 `json:"streamer_near_offset" default:""`
	Vessel_azimuth       *float64 `json:"vessel_azimuth" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
