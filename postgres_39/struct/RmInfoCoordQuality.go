package dto

import (
	"time"
)

type Rm_info_coord_quality struct {
	Info_item_subtype        string     `json:"info_item_subtype"`
	Information_item_id      string     `json:"information_item_id"`
	Quality_id               string     `json:"quality_id"`
	Active_ind               *string    `json:"active_ind"`
	Completeness_desc        *string    `json:"completeness_desc"`
	Coord_accuracy_desc      *string    `json:"coord_accuracy_desc"`
	Coord_acquisition_id     *string    `json:"coord_acquisition_id"`
	Coord_quality            *string    `json:"coord_quality"`
	Coord_quality_desc       *string    `json:"coord_quality_desc"`
	Corrected_date           *time.Time `json:"corrected_date"`
	Deficiency_ind           *string    `json:"deficiency_ind"`
	Deficiency_type          *string    `json:"deficiency_type"`
	Described_by_ba_id       *string    `json:"described_by_ba_id"`
	Description              *string    `json:"description"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Horizontal_accuracy      *float64   `json:"horizontal_accuracy"`
	Horizontal_accuracy_desc *string    `json:"horizontal_accuracy_desc"`
	Horizontal_accuracy_ouom *string    `json:"horizontal_accuracy_ouom"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Quality_date             *time.Time `json:"quality_date"`
	Remark                   *string    `json:"remark"`
	Source                   *string    `json:"source"`
	Vertical_accuracy        *float64   `json:"vertical_accuracy"`
	Vertical_accuracy_desc   *string    `json:"vertical_accuracy_desc"`
	Vertical_accuracy_ouom   *string    `json:"vertical_accuracy_ouom"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
