package dto

import (
	"time"
)

type Cs_coord_acquisition struct {
	Acquisition_id           string     `json:"acquisition_id"`
	Acquired_by_ba_id        *string    `json:"acquired_by_ba_id"`
	Active_ind               *string    `json:"active_ind"`
	Capture_method           *string    `json:"capture_method"`
	Compute_method           *string    `json:"compute_method"`
	Coordinate_quality       *string    `json:"coordinate_quality"`
	Data_create_date         *time.Time `json:"data_create_date"`
	Digitized_scale          *int       `json:"digitized_scale"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Horizontal_accuracy      *int       `json:"horizontal_accuracy"`
	Horizontal_accuracy_ouom *string    `json:"horizontal_accuracy_ouom"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Remark                   *string    `json:"remark"`
	Source                   *string    `json:"source"`
	Source_scale             *int       `json:"source_scale"`
	Survey_accuracy          *int       `json:"survey_accuracy"`
	Transform_id             *string    `json:"transform_id"`
	Vertical_accuracy        *float64   `json:"vertical_accuracy"`
	Vertical_accuracy_ouom   *string    `json:"vertical_accuracy_ouom"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
