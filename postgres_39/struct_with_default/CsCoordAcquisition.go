package dto

type Cs_coord_acquisition struct {
	Acquisition_id           string   `json:"acquisition_id" default:"source"`
	Acquired_by_ba_id        *string  `json:"acquired_by_ba_id" default:""`
	Active_ind               *string  `json:"active_ind" default:""`
	Capture_method           *string  `json:"capture_method" default:""`
	Compute_method           *string  `json:"compute_method" default:""`
	Coordinate_quality       *string  `json:"coordinate_quality" default:""`
	Data_create_date         *string  `json:"data_create_date" default:""`
	Digitized_scale          *int     `json:"digitized_scale" default:""`
	Effective_date           *string  `json:"effective_date" default:""`
	Expiry_date              *string  `json:"expiry_date" default:""`
	Horizontal_accuracy      *int     `json:"horizontal_accuracy" default:""`
	Horizontal_accuracy_ouom *string  `json:"horizontal_accuracy_ouom" default:""`
	Ppdm_guid                *string  `json:"ppdm_guid" default:""`
	Remark                   *string  `json:"remark" default:""`
	Source                   *string  `json:"source" default:""`
	Source_scale             *int     `json:"source_scale" default:""`
	Survey_accuracy          *int     `json:"survey_accuracy" default:""`
	Transform_id             *string  `json:"transform_id" default:""`
	Vertical_accuracy        *float64 `json:"vertical_accuracy" default:""`
	Vertical_accuracy_ouom   *string  `json:"vertical_accuracy_ouom" default:""`
	Row_changed_by           *string  `json:"row_changed_by" default:""`
	Row_changed_date         *string  `json:"row_changed_date" default:""`
	Row_created_by           *string  `json:"row_created_by" default:""`
	Row_created_date         *string  `json:"row_created_date" default:""`
	Row_effective_date       *string  `json:"row_effective_date" default:""`
	Row_expiry_date          *string  `json:"row_expiry_date" default:""`
	Row_quality              *string  `json:"row_quality" default:""`
}
