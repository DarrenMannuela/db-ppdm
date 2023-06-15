package dto

type Substance_composition struct {
	Substance_id             string  `json:"substance_id" default:"source"`
	Sub_substance_id         string  `json:"sub_substance_id" default:"source"`
	Composition_obs_no       int     `json:"composition_obs_no" default:"1"`
	Active_ind               *string `json:"active_ind" default:""`
	Defined_by_ba_id         *string `json:"defined_by_ba_id" default:""`
	Effective_date           *string `json:"effective_date" default:""`
	Expiry_date              *string `json:"expiry_date" default:""`
	Formula                  *string `json:"formula" default:""`
	Ppdm_guid                *string `json:"ppdm_guid" default:""`
	Remark                   *string `json:"remark" default:""`
	Source                   *string `json:"source" default:""`
	Source_document_id       *string `json:"source_document_id" default:""`
	Substance_component_type *string `json:"substance_component_type" default:""`
	Row_changed_by           *string `json:"row_changed_by" default:""`
	Row_changed_date         *string `json:"row_changed_date" default:""`
	Row_created_by           *string `json:"row_created_by" default:""`
	Row_created_date         *string `json:"row_created_date" default:""`
	Row_effective_date       *string `json:"row_effective_date" default:""`
	Row_expiry_date          *string `json:"row_expiry_date" default:""`
	Row_quality              *string `json:"row_quality" default:""`
}
