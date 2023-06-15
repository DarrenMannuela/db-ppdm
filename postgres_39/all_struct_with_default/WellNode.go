package dto

type Well_node struct {
	Node_id               string   `json:"node_id" default:"source"`
	Active_ind            *string  `json:"active_ind" default:""`
	Coordinate_quality    *string  `json:"coordinate_quality" default:""`
	Coord_acquisition_id  *string  `json:"coord_acquisition_id" default:""`
	Coord_system_id       *string  `json:"coord_system_id" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Latitude              *float64 `json:"latitude" default:""`
	Legal_survey_type     *string  `json:"legal_survey_type" default:""`
	Local_coord_system_id *string  `json:"local_coord_system_id" default:""`
	Location_quality      *string  `json:"location_quality" default:""`
	Location_type         *string  `json:"location_type" default:""`
	Longitude             *float64 `json:"longitude" default:""`
	Node_position         *string  `json:"node_position" default:""`
	Original_obs_no       *int     `json:"original_obs_no" default:""`
	Original_xy_uom       *string  `json:"original_xy_uom" default:""`
	Platform_id           *string  `json:"platform_id" default:""`
	Platform_sf_subtype   *string  `json:"platform_sf_subtype" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Preferred_ind         *string  `json:"preferred_ind" default:""`
	Remark                *string  `json:"remark" default:""`
	Selected_obs_no       *int     `json:"selected_obs_no" default:""`
	Source                *string  `json:"source" default:""`
	Uwi                   *string  `json:"uwi" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
