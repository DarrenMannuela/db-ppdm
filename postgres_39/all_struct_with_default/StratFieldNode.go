package dto

type Strat_field_node struct {
	Field_station_id      string   `json:"field_station_id" default:"source"`
	Node_id               string   `json:"node_id" default:"source"`
	Source                string   `json:"source" default:"source"`
	Active_ind            *string  `json:"active_ind" default:""`
	Coord_acquisition_id  *string  `json:"coord_acquisition_id" default:""`
	Coord_system_id       *string  `json:"coord_system_id" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Elev                  *float64 `json:"elev" default:""`
	Elev_ouom             *string  `json:"elev_ouom" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Latitude              *float64 `json:"latitude" default:""`
	Local_coord_system_id *string  `json:"local_coord_system_id" default:""`
	Longitude             *float64 `json:"longitude" default:""`
	Node_loc_type         *string  `json:"node_loc_type" default:""`
	Original_obs_no       *int     `json:"original_obs_no" default:""`
	Original_xy_uom       *string  `json:"original_xy_uom" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Remark                *string  `json:"remark" default:""`
	Selected_obs_no       *int     `json:"selected_obs_no" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
