package dto

type Seis_streamer_build struct {
	Streamer_id           string   `json:"streamer_id" default:"source"`
	Streamer_component_id string   `json:"streamer_component_id" default:"source"`
	Active_ind            *string  `json:"active_ind" default:""`
	Component_type        *string  `json:"component_type" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Position_ouom         *string  `json:"position_ouom" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Remark                *string  `json:"remark" default:""`
	Source                *string  `json:"source" default:""`
	Streamer_position     *float64 `json:"streamer_position" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
