package dto

type Sp_polygon struct {
	Polygon_id              string   `json:"polygon_id" default:"source"`
	Acquisition_id          *string  `json:"acquisition_id" default:""`
	Active_ind              *string  `json:"active_ind" default:""`
	Boundary_direction      *string  `json:"boundary_direction" default:""`
	Contained_by_polygon_id *string  `json:"contained_by_polygon_id" default:""`
	Coord_system_id         *string  `json:"coord_system_id" default:""`
	Datum_elev              *float64 `json:"datum_elev" default:""`
	Datum_elev_ouom         *string  `json:"datum_elev_ouom" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	Exclusion_ind           *string  `json:"exclusion_ind" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Inclusion_ind           *string  `json:"inclusion_ind" default:""`
	Local_coord_system_id   *string  `json:"local_coord_system_id" default:""`
	Location_type           *string  `json:"location_type" default:""`
	Max_plot_scale          *string  `json:"max_plot_scale" default:""`
	Min_plot_scale          *string  `json:"min_plot_scale" default:""`
	Polygon_set_id          *string  `json:"polygon_set_id" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Preferred_ind           *string  `json:"preferred_ind" default:""`
	Reference_datum         *string  `json:"reference_datum" default:""`
	Remark                  *string  `json:"remark" default:""`
	Source                  *string  `json:"source" default:""`
	Spatial_description_id  *string  `json:"spatial_description_id" default:""`
	Spatial_obs_no          *int     `json:"spatial_obs_no" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
