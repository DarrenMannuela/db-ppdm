package dto

type Well_activity_type_equiv struct {
	Activity_set_id    string  `json:"activity_set_id" default:"source"`
	Activity_type_id   string  `json:"activity_type_id" default:"source"`
	Activity_set_id2   string  `json:"activity_set_id_2" default:"source"`
	Activity_type_id2  string  `json:"activity_type_id_2" default:"source"`
	Equiv_obs_no       int     `json:"equiv_obs_no" default:"1"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Equiv_owner_ba_id  *string `json:"equiv_owner_ba_id" default:""`
	Equiv_type         *string `json:"equiv_type" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Full_name          *string `json:"full_name" default:""`
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
