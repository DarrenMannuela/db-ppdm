package dto

type R_well_status_xref struct {
	Status_xref_id     string  `json:"status_xref_id" default:"source"`
	Status_xref_obs_no int     `json:"status_xref_obs_no" default:"1"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Mapping_count      *int    `json:"mapping_count" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Qualifier_value1   *string `json:"qualifier_value_1" default:""`
	Qualifier_value2   *string `json:"qualifier_value_2" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Status1            *string `json:"status_1" default:""`
	Status2            *string `json:"status_2" default:""`
	Status_qualifier1  *string `json:"status_qualifier_1" default:""`
	Status_qualifier2  *string `json:"status_qualifier_2" default:""`
	Status_type1       *string `json:"status_type_1" default:""`
	Status_type2       *string `json:"status_type_2" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
