package dto

type Sp_point struct {
	Point_id               string   `json:"point_id" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Coord_acquisition_id   *string  `json:"coord_acquisition_id" default:""`
	Coord_system_id        *string  `json:"coord_system_id" default:""`
	Datum_elev             *float64 `json:"datum_elev" default:""`
	Datum_elev_ouom        *string  `json:"datum_elev_ouom" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Elevation              *float64 `json:"elevation" default:""`
	Elevation_ouom         *string  `json:"elevation_ouom" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Latitude               *float64 `json:"latitude" default:""`
	Local_coord_system_id  *string  `json:"local_coord_system_id" default:""`
	Location_quality       *string  `json:"location_quality" default:""`
	Location_type          *string  `json:"location_type" default:""`
	Longitude              *float64 `json:"longitude" default:""`
	Measured_depth         *float64 `json:"measured_depth" default:""`
	Measured_depth_ouom    *string  `json:"measured_depth_ouom" default:""`
	Point_label            *string  `json:"point_label" default:""`
	Point_no               *float64 `json:"point_no" default:""`
	Point_position         *string  `json:"point_position" default:""`
	Point_seq_no           *int     `json:"point_seq_no" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Reference_datum        *string  `json:"reference_datum" default:""`
	Remark                 *string  `json:"remark" default:""`
	Source                 *string  `json:"source" default:""`
	Spatial_description_id *string  `json:"spatial_description_id" default:""`
	Spatial_obs_no         *int     `json:"spatial_obs_no" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
