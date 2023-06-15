package dto

type Ppdm_system_map struct {
	Map_id             string  `json:"map_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Map_owner_ba_id    *string `json:"map_owner_ba_id" default:""`
	Map_version_num    *string `json:"map_version_num" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Preferred_ind      *string `json:"preferred_ind" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Sw_application_id  *string `json:"sw_application_id" default:""`
	System_id1         *string `json:"system_id_1" default:""`
	System_id2         *string `json:"system_id_2" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
