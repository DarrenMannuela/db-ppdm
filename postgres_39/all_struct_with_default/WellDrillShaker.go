package dto

type Well_drill_shaker struct {
	Uwi                string  `json:"uwi" default:"source"`
	Period_obs_no      int     `json:"period_obs_no" default:"1"`
	Shaker_id          string  `json:"shaker_id" default:"source"`
	Screen_obs_no      int     `json:"screen_obs_no" default:"1"`
	Active_ind         *string `json:"active_ind" default:""`
	Catalogue_equip_id *string `json:"catalogue_equip_id" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Equipment_id       *string `json:"equipment_id" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	New_screen_ind     *string `json:"new_screen_ind" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Screen_change_ind  *string `json:"screen_change_ind" default:""`
	Screen_location    *string `json:"screen_location" default:""`
	Screen_mesh_desc   *string `json:"screen_mesh_desc" default:""`
	Source             *string `json:"source" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
