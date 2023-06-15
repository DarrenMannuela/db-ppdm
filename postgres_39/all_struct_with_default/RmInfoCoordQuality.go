package dto

type Rm_info_coord_quality struct {
	Info_item_subtype        string   `json:"info_item_subtype" default:"source"`
	Information_item_id      string   `json:"information_item_id" default:"source"`
	Quality_id               string   `json:"quality_id" default:"source"`
	Active_ind               *string  `json:"active_ind" default:""`
	Completeness_desc        *string  `json:"completeness_desc" default:""`
	Coord_accuracy_desc      *string  `json:"coord_accuracy_desc" default:""`
	Coord_acquisition_id     *string  `json:"coord_acquisition_id" default:""`
	Coord_quality            *string  `json:"coord_quality" default:""`
	Coord_quality_desc       *string  `json:"coord_quality_desc" default:""`
	Corrected_date           *string  `json:"corrected_date" default:""`
	Deficiency_ind           *string  `json:"deficiency_ind" default:""`
	Deficiency_type          *string  `json:"deficiency_type" default:""`
	Described_by_ba_id       *string  `json:"described_by_ba_id" default:""`
	Description              *string  `json:"description" default:""`
	Effective_date           *string  `json:"effective_date" default:""`
	Expiry_date              *string  `json:"expiry_date" default:""`
	Horizontal_accuracy      *float64 `json:"horizontal_accuracy" default:""`
	Horizontal_accuracy_desc *string  `json:"horizontal_accuracy_desc" default:""`
	Horizontal_accuracy_ouom *string  `json:"horizontal_accuracy_ouom" default:""`
	Ppdm_guid                *string  `json:"ppdm_guid" default:""`
	Quality_date             *string  `json:"quality_date" default:""`
	Remark                   *string  `json:"remark" default:""`
	Source                   *string  `json:"source" default:""`
	Vertical_accuracy        *float64 `json:"vertical_accuracy" default:""`
	Vertical_accuracy_desc   *string  `json:"vertical_accuracy_desc" default:""`
	Vertical_accuracy_ouom   *string  `json:"vertical_accuracy_ouom" default:""`
	Row_changed_by           *string  `json:"row_changed_by" default:""`
	Row_changed_date         *string  `json:"row_changed_date" default:""`
	Row_created_by           *string  `json:"row_created_by" default:""`
	Row_created_date         *string  `json:"row_created_date" default:""`
	Row_effective_date       *string  `json:"row_effective_date" default:""`
	Row_expiry_date          *string  `json:"row_expiry_date" default:""`
	Row_quality              *string  `json:"row_quality" default:""`
}
