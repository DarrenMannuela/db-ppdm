package dto

type Well_drill_equipment struct {
	Uwi                 string   `json:"uwi" default:"source"`
	Equipment_id        string   `json:"equipment_id" default:"source"`
	Equipment_obs_no    int      `json:"equipment_obs_no" default:"1"`
	Active_ind          *string  `json:"active_ind" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Offsite_date        *string  `json:"offsite_date" default:""`
	Offsite_time        *string  `json:"offsite_time" default:""`
	Onsite_date         *string  `json:"onsite_date" default:""`
	Onsite_time         *string  `json:"onsite_time" default:""`
	Operated_by_ba_id   *string  `json:"operated_by_ba_id" default:""`
	Parent_equipment_id *string  `json:"parent_equipment_id" default:""`
	Period_on_well      *float64 `json:"period_on_well" default:""`
	Period_on_well_ouom *string  `json:"period_on_well_ouom" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Reference_num       *string  `json:"reference_num" default:""`
	Remark              *string  `json:"remark" default:""`
	Rented_ind          *string  `json:"rented_ind" default:""`
	Source              *string  `json:"source" default:""`
	Timezone            *string  `json:"timezone" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
