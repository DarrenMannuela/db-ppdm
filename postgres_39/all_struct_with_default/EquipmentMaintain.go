package dto

type Equipment_maintain struct {
	Equipment_id               string  `json:"equipment_id" default:"source"`
	Equip_maint_id             string  `json:"equip_maint_id" default:"source"`
	Active_ind                 *string `json:"active_ind" default:""`
	Actual_end_date            *string `json:"actual_end_date" default:""`
	Actual_start_date          *string `json:"actual_start_date" default:""`
	Catalogue_equip_id         *string `json:"catalogue_equip_id" default:""`
	Completed_by_ba_id         *string `json:"completed_by_ba_id" default:""`
	Effective_date             *string `json:"effective_date" default:""`
	End_date                   *string `json:"end_date" default:""`
	Expiry_date                *string `json:"expiry_date" default:""`
	Failure_ind                *string `json:"failure_ind" default:""`
	Location_ba_address_obs_no *int    `json:"location_ba_address_obs_no" default:""`
	Location_ba_id             *string `json:"location_ba_id" default:""`
	Location_ba_source         *string `json:"location_ba_source" default:""`
	Maint_location_type        *string `json:"maint_location_type" default:""`
	Maint_reason               *string `json:"maint_reason" default:""`
	Maint_type                 *string `json:"maint_type" default:""`
	Ppdm_guid                  *string `json:"ppdm_guid" default:""`
	Project_id                 *string `json:"project_id" default:""`
	Remark                     *string `json:"remark" default:""`
	Scheduled_date             *string `json:"scheduled_date" default:""`
	Scheduled_ind              *string `json:"scheduled_ind" default:""`
	Source                     *string `json:"source" default:""`
	Start_date                 *string `json:"start_date" default:""`
	System_condition           *string `json:"system_condition" default:""`
	Row_changed_by             *string `json:"row_changed_by" default:""`
	Row_changed_date           *string `json:"row_changed_date" default:""`
	Row_created_by             *string `json:"row_created_by" default:""`
	Row_created_date           *string `json:"row_created_date" default:""`
	Row_effective_date         *string `json:"row_effective_date" default:""`
	Row_expiry_date            *string `json:"row_expiry_date" default:""`
	Row_quality                *string `json:"row_quality" default:""`
}
