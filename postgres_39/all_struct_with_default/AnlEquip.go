package dto

type Anl_equip struct {
	Analysis_id         string  `json:"analysis_id" default:"source"`
	Anl_source          string  `json:"anl_source" default:"source"`
	Equip_obs_no        int     `json:"equip_obs_no" default:"1"`
	Active_ind          *string `json:"active_ind" default:""`
	Callibration_record *string `json:"callibration_record" default:""`
	Catalogue_equip_id  *string `json:"catalogue_equip_id" default:""`
	Effective_date      *string `json:"effective_date" default:""`
	Equip_id            *string `json:"equip_id" default:""`
	Equip_role          *string `json:"equip_role" default:""`
	Expiry_date         *string `json:"expiry_date" default:""`
	Lab_ba_id           *string `json:"lab_ba_id" default:""`
	Ppdm_guid           *string `json:"ppdm_guid" default:""`
	Problem_ind         *string `json:"problem_ind" default:""`
	Remark              *string `json:"remark" default:""`
	Source              *string `json:"source" default:""`
	Step_seq_no         *int    `json:"step_seq_no" default:""`
	Row_changed_by      *string `json:"row_changed_by" default:""`
	Row_changed_date    *string `json:"row_changed_date" default:""`
	Row_created_by      *string `json:"row_created_by" default:""`
	Row_created_date    *string `json:"row_created_date" default:""`
	Row_effective_date  *string `json:"row_effective_date" default:""`
	Row_expiry_date     *string `json:"row_expiry_date" default:""`
	Row_quality         *string `json:"row_quality" default:""`
}
