package dto

type Sf_rig_overhead_equip struct {
	Sf_subtype           string   `json:"sf_subtype" default:"source"`
	Rig_id               string   `json:"rig_id" default:"source"`
	Equip_id             string   `json:"equip_id" default:"source"`
	Active_ind           *string  `json:"active_ind" default:""`
	Capacity             *float64 `json:"capacity" default:""`
	Capacity_ouom        *string  `json:"capacity_ouom" default:""`
	Capacity_type        *string  `json:"capacity_type" default:""`
	Capacity_uom         *string  `json:"capacity_uom" default:""`
	Catalogue_equip_id   *string  `json:"catalogue_equip_id" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Equipment_id         *string  `json:"equipment_id" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Input_type           *string  `json:"input_type" default:""`
	Inside_diameter      *float64 `json:"inside_diameter" default:""`
	Inside_diameter_ouom *string  `json:"inside_diameter_ouom" default:""`
	Install_date         *string  `json:"install_date" default:""`
	Overhead_count       *int     `json:"overhead_count" default:""`
	Overhead_equip_type  *string  `json:"overhead_equip_type" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Remark               *string  `json:"remark" default:""`
	Remove_date          *string  `json:"remove_date" default:""`
	Source               *string  `json:"source" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
