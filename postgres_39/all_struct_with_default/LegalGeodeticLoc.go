package dto

type Legal_geodetic_loc struct {
	Legal_geodetic_loc_id   string   `json:"legal_geodetic_loc_id" default:"source"`
	Location_type           string   `json:"location_type" default:"source"`
	Source                  string   `json:"source" default:"source"`
	Active_ind              *string  `json:"active_ind" default:""`
	Coord_system_id         *string  `json:"coord_system_id" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Geodetic_event_sequence *string  `json:"geodetic_event_sequence" default:""`
	Geodetic_loc_exception  *string  `json:"geodetic_loc_exception" default:""`
	Latitude                *float64 `json:"latitude" default:""`
	Longitude               *float64 `json:"longitude" default:""`
	Node_id                 *string  `json:"node_id" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Remark                  *string  `json:"remark" default:""`
	Uwi                     *string  `json:"uwi" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
