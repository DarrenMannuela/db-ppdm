package dto

type Well_sieve_screen struct {
	Uwi                 string   `json:"uwi" default:"source"`
	Anl_source          string   `json:"anl_source" default:"source"`
	Analysis_obs_no     int      `json:"analysis_obs_no" default:"1"`
	Screen_obs_no       int      `json:"screen_obs_no" default:"1"`
	Active_ind          *string  `json:"active_ind" default:""`
	Cat_equip_id        *string  `json:"cat_equip_id" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Equipment_id        *string  `json:"equipment_id" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Particle_held_count *int     `json:"particle_held_count" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Remark              *string  `json:"remark" default:""`
	Screen_mesh_size    *float64 `json:"screen_mesh_size" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
