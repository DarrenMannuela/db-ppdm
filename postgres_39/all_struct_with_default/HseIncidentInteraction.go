package dto

type Hse_incident_interaction struct {
	Incident_id        string  `json:"incident_id" default:"source"`
	Interaction_obs_no int     `json:"interaction_obs_no" default:"1"`
	Active_ind         *string `json:"active_ind" default:""`
	Cause_obs_no       *int    `json:"cause_obs_no" default:""`
	Description        *string `json:"description" default:""`
	Detail_obs_no      *int    `json:"detail_obs_no" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Equip_obs_no       *int    `json:"equip_obs_no" default:""`
	Equip_role_obs_no  *int    `json:"equip_role_obs_no" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Incident_substance *string `json:"incident_substance" default:""`
	Interaction_type   *string `json:"interaction_type" default:""`
	Party_obs_no       *int    `json:"party_obs_no" default:""`
	Party_role_obs_no  *int    `json:"party_role_obs_no" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Response_obs_no    *int    `json:"response_obs_no" default:""`
	Source             *string `json:"source" default:""`
	Substance_seq_no   *int    `json:"substance_seq_no" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
