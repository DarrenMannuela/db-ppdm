package dto

type Sf_rig_shaker struct {
	Sf_subtype         string   `json:"sf_subtype" default:"source"`
	Rig_id             string   `json:"rig_id" default:"source"`
	Shaker_id          string   `json:"shaker_id" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Capacity           *float64 `json:"capacity" default:""`
	Capacity_ouom      *string  `json:"capacity_ouom" default:""`
	Catalogue_equip_id *string  `json:"catalogue_equip_id" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Equipment_id       *string  `json:"equipment_id" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Input_type         *string  `json:"input_type" default:""`
	Install_date       *string  `json:"install_date" default:""`
	Position_desc      *string  `json:"position_desc" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Reference_num      *string  `json:"reference_num" default:""`
	Remark             *string  `json:"remark" default:""`
	Remove_date        *string  `json:"remove_date" default:""`
	Shaker_count       *int     `json:"shaker_count" default:""`
	Shaker_size        *float64 `json:"shaker_size" default:""`
	Shaker_size_desc   *string  `json:"shaker_size_desc" default:""`
	Shaker_size_ouom   *string  `json:"shaker_size_ouom" default:""`
	Source             *string  `json:"source" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
