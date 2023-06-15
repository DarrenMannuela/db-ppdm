package dto

type Equipment_status struct {
	Equipment_id       string   `json:"equipment_id" default:"source"`
	Equip_status_id    string   `json:"equip_status_id" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	End_time           *string  `json:"end_time" default:""`
	Equip_status       *string  `json:"equip_status" default:""`
	Equip_status_type  *string  `json:"equip_status_type" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Percent_capability *float64 `json:"percent_capability" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Preferred_ind      *string  `json:"preferred_ind" default:""`
	Remark             *string  `json:"remark" default:""`
	Source             *string  `json:"source" default:""`
	Start_time         *string  `json:"start_time" default:""`
	Timezone           *string  `json:"timezone" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
