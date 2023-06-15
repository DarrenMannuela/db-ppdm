package dto

type Seis_streamer struct {
	Streamer_id          string   `json:"streamer_id" default:"source"`
	Acqtn_design_id      *string  `json:"acqtn_design_id" default:""`
	Active_ind           *string  `json:"active_ind" default:""`
	Cable_make           *string  `json:"cable_make" default:""`
	Description          *string  `json:"description" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Remark               *string  `json:"remark" default:""`
	Sf_subtype           *string  `json:"sf_subtype" default:""`
	Source               *string  `json:"source" default:""`
	Streamer_length      *float64 `json:"streamer_length" default:""`
	Streamer_length_ouom *string  `json:"streamer_length_ouom" default:""`
	Streamer_position    *string  `json:"streamer_position" default:""`
	Vessel_config_obs_no *int     `json:"vessel_config_obs_no" default:""`
	Vessel_id            *string  `json:"vessel_id" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
