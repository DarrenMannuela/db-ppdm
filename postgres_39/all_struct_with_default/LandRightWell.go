package dto

type Land_right_well struct {
	Uwi                    string   `json:"uwi" default:"source"`
	Land_right_subtype     string   `json:"land_right_subtype" default:"source"`
	Land_right_id          string   `json:"land_right_id" default:"source"`
	Lr_well_seq_no         int      `json:"lr_well_seq_no" default:"1"`
	Active_ind             *string  `json:"active_ind" default:""`
	Completion_obs_no      *int     `json:"completion_obs_no" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Gas_percent_psu        *float64 `json:"gas_percent_psu" default:""`
	Key_well_ind           *string  `json:"key_well_ind" default:""`
	Oil_percent_psu        *float64 `json:"oil_percent_psu" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Pr_str_form_obs_no     *int     `json:"pr_str_form_obs_no" default:""`
	Pr_str_source          *string  `json:"pr_str_source" default:""`
	Remark                 *string  `json:"remark" default:""`
	Source                 *string  `json:"source" default:""`
	Spacing_complete_ind   *string  `json:"spacing_complete_ind" default:""`
	Spacing_unit_id        *string  `json:"spacing_unit_id" default:""`
	String_id              *string  `json:"string_id" default:""`
	Well_in_tract_ind      *string  `json:"well_in_tract_ind" default:""`
	Well_relationship_type *string  `json:"well_relationship_type" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
