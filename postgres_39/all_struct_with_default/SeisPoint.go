package dto

type Seis_point struct {
	Seis_set_subtype    string   `json:"seis_set_subtype" default:"source"`
	Seis_set_id         string   `json:"seis_set_id" default:"source"`
	Seis_point_id       string   `json:"seis_point_id" default:"source"`
	Acqtn_seq_no        *int     `json:"acqtn_seq_no" default:""`
	Active_ind          *string  `json:"active_ind" default:""`
	Bend_ind            *string  `json:"bend_ind" default:""`
	Datum_elev          *float64 `json:"datum_elev" default:""`
	Datum_elev_ouom     *string  `json:"datum_elev_ouom" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Elev                *float64 `json:"elev" default:""`
	Elev_ouom           *string  `json:"elev_ouom" default:""`
	End_ind             *string  `json:"end_ind" default:""`
	Exception_ind       *string  `json:"exception_ind" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Flowing_hole_ind    *string  `json:"flowing_hole_ind" default:""`
	Mapping_code        *string  `json:"mapping_code" default:""`
	Measured_depth      *float64 `json:"measured_depth" default:""`
	Measured_depth_ouom *string  `json:"measured_depth_ouom" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Reference_datum     *string  `json:"reference_datum" default:""`
	Remark              *string  `json:"remark" default:""`
	Seis_point_label    *string  `json:"seis_point_label" default:""`
	Seis_point_lat      *float64 `json:"seis_point_lat" default:""`
	Seis_point_long     *float64 `json:"seis_point_long" default:""`
	Seis_point_no       *int     `json:"seis_point_no" default:""`
	Seis_station_type   *string  `json:"seis_station_type" default:""`
	Source              *string  `json:"source" default:""`
	Spatial_seq_no      *int     `json:"spatial_seq_no" default:""`
	X_coordinate        *float64 `json:"x_coordinate" default:""`
	Y_coordinate        *float64 `json:"y_coordinate" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
