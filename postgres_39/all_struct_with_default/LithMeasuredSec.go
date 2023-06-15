package dto

type Lith_measured_sec struct {
	Meas_section_id            string  `json:"meas_section_id" default:"source"`
	Source                     string  `json:"source" default:"source"`
	Active_ind                 *string `json:"active_ind" default:""`
	Author                     *string `json:"author" default:""`
	Description                *string `json:"description" default:""`
	Effective_date             *string `json:"effective_date" default:""`
	Expiry_date                *string `json:"expiry_date" default:""`
	Location_desc              *string `json:"location_desc" default:""`
	Location_qualifier         *string `json:"location_qualifier" default:""`
	Node_id                    *string `json:"node_id" default:""`
	Ppdm_guid                  *string `json:"ppdm_guid" default:""`
	Publication_reference_text *string `json:"publication_reference_text" default:""`
	Reference_date             *string `json:"reference_date" default:""`
	Remark                     *string `json:"remark" default:""`
	Source_document_id         *string `json:"source_document_id" default:""`
	Strat_name_set_id          *string `json:"strat_name_set_id" default:""`
	Strat_unit_id              *string `json:"strat_unit_id" default:""`
	Row_changed_by             *string `json:"row_changed_by" default:""`
	Row_changed_date           *string `json:"row_changed_date" default:""`
	Row_created_by             *string `json:"row_created_by" default:""`
	Row_created_date           *string `json:"row_created_date" default:""`
	Row_effective_date         *string `json:"row_effective_date" default:""`
	Row_expiry_date            *string `json:"row_expiry_date" default:""`
	Row_quality                *string `json:"row_quality" default:""`
}
