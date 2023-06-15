package dto

type Sample_lith_desc struct {
	Sample_id                string   `json:"sample_id" default:"source"`
	Description_id           string   `json:"description_id" default:"source"`
	Active_ind               *string  `json:"active_ind" default:""`
	Color_intensity          *string  `json:"color_intensity" default:""`
	Color_type               *string  `json:"color_type" default:""`
	Contaminant              *string  `json:"contaminant" default:""`
	Depositional_environment *string  `json:"depositional_environment" default:""`
	Diagenesis_type          *string  `json:"diagenesis_type" default:""`
	Ecozone_id               *string  `json:"ecozone_id" default:""`
	Effective_date           *string  `json:"effective_date" default:""`
	Expiry_date              *string  `json:"expiry_date" default:""`
	Lithology                *string  `json:"lithology" default:""`
	Oil_stain                *string  `json:"oil_stain" default:""`
	Percent                  *float64 `json:"percent" default:""`
	Permeability_type        *string  `json:"permeability_type" default:""`
	Porosity_type            *string  `json:"porosity_type" default:""`
	Ppdm_guid                *string  `json:"ppdm_guid" default:""`
	Remark                   *string  `json:"remark" default:""`
	Rock_matrix              *string  `json:"rock_matrix" default:""`
	Rock_type                *string  `json:"rock_type" default:""`
	Source                   *string  `json:"source" default:""`
	Strat_name_set_id        *string  `json:"strat_name_set_id" default:""`
	Strat_unit_id            *string  `json:"strat_unit_id" default:""`
	Row_changed_by           *string  `json:"row_changed_by" default:""`
	Row_changed_date         *string  `json:"row_changed_date" default:""`
	Row_created_by           *string  `json:"row_created_by" default:""`
	Row_created_date         *string  `json:"row_created_date" default:""`
	Row_effective_date       *string  `json:"row_effective_date" default:""`
	Row_expiry_date          *string  `json:"row_expiry_date" default:""`
	Row_quality              *string  `json:"row_quality" default:""`
}
