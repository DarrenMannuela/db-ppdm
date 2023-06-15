package dto

type Sp_zone_definition struct {
	Zone_definition_id     string   `json:"zone_definition_id" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Base_depth             *float64 `json:"base_depth" default:""`
	Base_depth_ouom        *string  `json:"base_depth_ouom" default:""`
	Base_qualifier         *string  `json:"base_qualifier" default:""`
	Base_strat_unit_id     *string  `json:"base_strat_unit_id" default:""`
	Defined_strat_unit_id  *string  `json:"defined_strat_unit_id" default:""`
	Description            *string  `json:"description" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Owner_ba_id            *string  `json:"owner_ba_id" default:""`
	Owner_ref_num          *string  `json:"owner_ref_num" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Remark                 *string  `json:"remark" default:""`
	Source                 *string  `json:"source" default:""`
	Strat_name_set_id      *string  `json:"strat_name_set_id" default:""`
	Top_depth              *float64 `json:"top_depth" default:""`
	Top_depth_ouom         *string  `json:"top_depth_ouom" default:""`
	Top_qualifier          *string  `json:"top_qualifier" default:""`
	Top_strat_unit_id      *string  `json:"top_strat_unit_id" default:""`
	Uwi                    *string  `json:"uwi" default:""`
	Zone_derivation_method *string  `json:"zone_derivation_method" default:""`
	Zone_type              *string  `json:"zone_type" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
