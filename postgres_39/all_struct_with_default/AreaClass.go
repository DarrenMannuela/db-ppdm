package dto

type Area_class struct {
	Area_id_parent     string  `json:"area_id_parent" default:"source"`
	Area_type_parent   string  `json:"area_type_parent" default:"source"`
	Area_id            string  `json:"area_id" default:"source"`
	Area_type          string  `json:"area_type" default:"source"`
	Area_hierarchy_id  string  `json:"area_hierarchy_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
