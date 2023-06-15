package dto

type Well_show struct {
	Uwi                string   `json:"uwi" default:"source"`
	Source             string   `json:"source" default:"source"`
	Show_type          string   `json:"show_type" default:"source"`
	Show_obs_no        int      `json:"show_obs_no" default:"1"`
	Active_ind         *string  `json:"active_ind" default:""`
	Base_depth         *float64 `json:"base_depth" default:""`
	Base_depth_ouom    *string  `json:"base_depth_ouom" default:""`
	Base_strat_unit_id *string  `json:"base_strat_unit_id" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Lithology_desc     *string  `json:"lithology_desc" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Reservoir_ind      *string  `json:"reservoir_ind" default:""`
	Sample_type        *string  `json:"sample_type" default:""`
	Show_symbol        *string  `json:"show_symbol" default:""`
	Source_document_id *string  `json:"source_document_id" default:""`
	Strat_name_set_id  *string  `json:"strat_name_set_id" default:""`
	Top_depth          *float64 `json:"top_depth" default:""`
	Top_depth_ouom     *string  `json:"top_depth_ouom" default:""`
	Top_strat_unit_id  *string  `json:"top_strat_unit_id" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
