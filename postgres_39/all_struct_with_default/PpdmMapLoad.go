package dto

type Ppdm_map_load struct {
	Map_id             string  `json:"map_id" default:"source"`
	Load_process_id    string  `json:"load_process_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Delete_allowed_ind *string `json:"delete_allowed_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	End_date           *string `json:"end_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Insert_allowed_ind *string `json:"insert_allowed_ind" default:""`
	Ppdm_group_id      *string `json:"ppdm_group_id" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Procedure_id       *string `json:"procedure_id" default:""`
	Procedure_name     *string `json:"procedure_name" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Source_system_id   *string `json:"source_system_id" default:""`
	Start_date         *string `json:"start_date" default:""`
	Sw_application_id  *string `json:"sw_application_id" default:""`
	Target_system_id   *string `json:"target_system_id" default:""`
	Update_allowed_ind *string `json:"update_allowed_ind" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
