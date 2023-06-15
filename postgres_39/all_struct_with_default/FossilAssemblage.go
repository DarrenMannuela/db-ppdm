package dto

type Fossil_assemblage struct {
	Strat_name_set_id  string  `json:"strat_name_set_id" default:"source"`
	Strat_unit_id      string  `json:"strat_unit_id" default:"source"`
	Taxon_leaf_id      string  `json:"taxon_leaf_id" default:"source"`
	Interp_id          string  `json:"interp_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Assemblage_type    *string `json:"assemblage_type" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Oldest_ind         *string `json:"oldest_ind" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Primary_marker_ind *string `json:"primary_marker_ind" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Source_document_id *string `json:"source_document_id" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
