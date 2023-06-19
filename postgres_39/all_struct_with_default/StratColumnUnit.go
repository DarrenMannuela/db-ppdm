package dto

type Strat_column_unit struct {
	Strat_column_id         string  `json:"strat_column_id" default:"source"`
	Strat_column_source     string  `json:"strat_column_source" default:"source"`
	Strat_name_set_id       string  `json:"strat_name_set_id" default:"source"`
	Strat_unit_id           string  `json:"strat_unit_id" default:"source"`
	Interp_id               string  `json:"interp_id" default:"source"`
	Active_ind              *string `json:"active_ind" default:""`
	Business_associate_id   *string `json:"business_associate_id" default:""`
	Certified_ind           *string `json:"certified_ind" default:""`
	Conformity_relationship *string `json:"conformity_relationship" default:""`
	Effective_date          *string `json:"effective_date" default:""`
	Expiry_date             *string `json:"expiry_date" default:""`
	Interpretation_version  *string `json:"interpretation_version" default:""`
	Lithology               *string `json:"lithology" default:""`
	Occurrence_type         *string `json:"occurrence_type" default:""`
	Ordinal_seq_no          *int    `json:"ordinal_seq_no" default:""`
	Overturned_ind          *string `json:"overturned_ind" default:""`
	Ppdm_guid               *string `json:"ppdm_guid" default:""`
	Preferred_ind           *string `json:"preferred_ind" default:""`
	Remark                  *string `json:"remark" default:""`
	Repeat_strat_occur_no   *int    `json:"repeat_strat_occur_no" default:""`
	Repeat_strat_type       *string `json:"repeat_strat_type" default:""`
	Source                  *string `json:"source" default:""`
	Source_document_id      *string `json:"source_document_id" default:""`
	Sour_gas_ind            *string `json:"sour_gas_ind" default:""`
	Strat_interpret_method  *string `json:"strat_interpret_method" default:""`
	Strat_topo_relation_ind *string `json:"strat_topo_relation_ind" default:""`
	Sw_application_id       *string `json:"sw_application_id" default:""`
	Version_obs_no          *int    `json:"version_obs_no" default:""`
	Row_changed_by          *string `json:"row_changed_by" default:""`
	Row_changed_date        *string `json:"row_changed_date" default:""`
	Row_created_by          *string `json:"row_created_by" default:""`
	Row_created_date        *string `json:"row_created_date" default:""`
	Row_effective_date      *string `json:"row_effective_date" default:""`
	Row_expiry_date         *string `json:"row_expiry_date" default:""`
	Row_quality             *string `json:"row_quality" default:""`
}