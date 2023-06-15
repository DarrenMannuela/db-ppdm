package dto

type Strat_topo_relation struct {
	Strat_name_set_id_1   string  `json:"strat_name_set_id_1" default:"source"`
	Strat_unit_id_1       string  `json:"strat_unit_id_1" default:"source"`
	Strat_name_set_id_2   string  `json:"strat_name_set_id_2" default:"source"`
	Strat_unit_id_2       string  `json:"strat_unit_id_2" default:"source"`
	Source                string  `json:"source" default:"source"`
	Topo_relation_id      string  `json:"topo_relation_id" default:"source"`
	Active_ind            *string `json:"active_ind" default:""`
	Area_id               *string `json:"area_id" default:""`
	Area_type             *string `json:"area_type" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Interp_id_1           *string `json:"interp_id_1" default:""`
	Interp_id_2           *string `json:"interp_id_2" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Remark                *string `json:"remark" default:""`
	Strat_column_id_1     *string `json:"strat_column_id_1" default:""`
	Strat_column_id_2     *string `json:"strat_column_id_2" default:""`
	Strat_column_source_1 *string `json:"strat_column_source_1" default:""`
	Strat_column_source_2 *string `json:"strat_column_source_2" default:""`
	Topo_relation_type    *string `json:"topo_relation_type" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
