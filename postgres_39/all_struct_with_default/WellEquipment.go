package dto

type Well_equipment struct {
	Uwi                     string   `json:"uwi" default:"source"`
	Source                  string   `json:"source" default:"source"`
	Equipment_id            string   `json:"equipment_id" default:"source"`
	Equip_obs_no            int      `json:"equip_obs_no" default:"1"`
	Active_ind              *string  `json:"active_ind" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Install_base_depth      *float64 `json:"install_base_depth" default:""`
	Install_base_depth_ouom *string  `json:"install_base_depth_ouom" default:""`
	Install_date            *string  `json:"install_date" default:""`
	Install_seq_no          *int     `json:"install_seq_no" default:""`
	Install_top_depth       *float64 `json:"install_top_depth" default:""`
	Install_top_depth_ouom  *string  `json:"install_top_depth_ouom" default:""`
	Parent_equipment_id     *string  `json:"parent_equipment_id" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Remark                  *string  `json:"remark" default:""`
	Removal_date            *string  `json:"removal_date" default:""`
	String_id               *string  `json:"string_id" default:""`
	String_source           *string  `json:"string_source" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
