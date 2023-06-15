package dto

type Sp_boundary struct {
	Polygon_id         string   `json:"polygon_id" default:"source"`
	Point_seq_no       int      `json:"point_seq_no" default:"1"`
	Active_ind         *string  `json:"active_ind" default:""`
	Depth              *float64 `json:"depth" default:""`
	Depth_ouom         *string  `json:"depth_ouom" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Elevation          *float64 `json:"elevation" default:""`
	Elevation_ouom     *string  `json:"elevation_ouom" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Latitude           *float64 `json:"latitude" default:""`
	Location_quality   *string  `json:"location_quality" default:""`
	Longitude          *float64 `json:"longitude" default:""`
	Point_label        *string  `json:"point_label" default:""`
	Point_no           *float64 `json:"point_no" default:""`
	Point_position     *string  `json:"point_position" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Source             *string  `json:"source" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
