package dto

type Strat_equivalence struct {
	Strat_name_set_id         string  `json:"strat_name_set_id" default:"source"`
	Strat_unit_id             string  `json:"strat_unit_id" default:"source"`
	Equiv_strat_name_set_id   string  `json:"equiv_strat_name_set_id" default:"source"`
	Equiv_strat_unit_id       string  `json:"equiv_strat_unit_id" default:"source"`
	Equiv_id                  string  `json:"equiv_id" default:"source"`
	Source                    string  `json:"source" default:"source"`
	Active_ind                *string `json:"active_ind" default:""`
	Area_id                   *string `json:"area_id" default:""`
	Area_id2                  *string `json:"area_id_2" default:""`
	Area_type                 *string `json:"area_type" default:""`
	Area_type2                *string `json:"area_type_2" default:""`
	Effective_date            *string `json:"effective_date" default:""`
	Equiv_direction           *string `json:"equiv_direction" default:""`
	Equiv_interp_id           *string `json:"equiv_interp_id" default:""`
	Equiv_strat_column_id     *string `json:"equiv_strat_column_id" default:""`
	Equiv_strat_column_source *string `json:"equiv_strat_column_source" default:""`
	Expiry_date               *string `json:"expiry_date" default:""`
	Ppdm_guid                 *string `json:"ppdm_guid" default:""`
	Remark                    *string `json:"remark" default:""`
	Source_document_id        *string `json:"source_document_id" default:""`
	Strat_column_id           *string `json:"strat_column_id" default:""`
	Strat_column_source       *string `json:"strat_column_source" default:""`
	Strat_equivalence_type    *string `json:"strat_equivalence_type" default:""`
	Strat_interp_id           *string `json:"strat_interp_id" default:""`
	Row_changed_by            *string `json:"row_changed_by" default:""`
	Row_changed_date          *string `json:"row_changed_date" default:""`
	Row_created_by            *string `json:"row_created_by" default:""`
	Row_created_date          *string `json:"row_created_date" default:""`
	Row_effective_date        *string `json:"row_effective_date" default:""`
	Row_expiry_date           *string `json:"row_expiry_date" default:""`
	Row_quality               *string `json:"row_quality" default:""`
}
