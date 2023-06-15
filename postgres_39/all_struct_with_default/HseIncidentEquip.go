package dto

type Hse_incident_equip struct {
	Incident_id        string  `json:"incident_id" default:"source"`
	Equip_obs_no       int     `json:"equip_obs_no" default:"1"`
	Equip_role_obs_no  int     `json:"equip_role_obs_no" default:"1"`
	Active_ind         *string `json:"active_ind" default:""`
	Crew_role          *string `json:"crew_role" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Equipment_group    *string `json:"equipment_group" default:""`
	Equipment_id       *string `json:"equipment_id" default:""`
	Equipment_type     *string `json:"equipment_type" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Period_obs_no      *int    `json:"period_obs_no" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Uwi                *string `json:"uwi" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
