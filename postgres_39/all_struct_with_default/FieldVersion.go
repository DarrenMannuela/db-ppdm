package dto

type Field_version struct {
	Field_id           string  `json:"field_id" default:"source"`
	Source             string  `json:"source" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Area_id            *string `json:"area_id" default:""`
	Area_type          *string `json:"area_type" default:""`
	Discovery_date     *string `json:"discovery_date" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Field_abbreviation *string `json:"field_abbreviation" default:""`
	Field_area_obs_no  *int    `json:"field_area_obs_no" default:""`
	Field_name         *string `json:"field_name" default:""`
	Field_type         *string `json:"field_type" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
