package dto

type Hse_incident_ba struct {
	Incident_id          string  `json:"incident_id" default:"source"`
	Party_obs_no         int     `json:"party_obs_no" default:"1"`
	Party_role_obs_no    int     `json:"party_role_obs_no" default:"1"`
	Active_ind           *string `json:"active_ind" default:""`
	Company_ba_id        *string `json:"company_ba_id" default:""`
	Detail_obs_no        *int    `json:"detail_obs_no" default:""`
	Effective_date       *string `json:"effective_date" default:""`
	Expiry_date          *string `json:"expiry_date" default:""`
	Involved_ba_role     *string `json:"involved_ba_role" default:""`
	Involved_ba_status   *string `json:"involved_ba_status" default:""`
	Involved_party_ba_id *string `json:"involved_party_ba_id" default:""`
	Period_obs_no        *int    `json:"period_obs_no" default:""`
	Ppdm_guid            *string `json:"ppdm_guid" default:""`
	Remark               *string `json:"remark" default:""`
	Source               *string `json:"source" default:""`
	Uwi                  *string `json:"uwi" default:""`
	Row_changed_by       *string `json:"row_changed_by" default:""`
	Row_changed_date     *string `json:"row_changed_date" default:""`
	Row_created_by       *string `json:"row_created_by" default:""`
	Row_created_date     *string `json:"row_created_date" default:""`
	Row_effective_date   *string `json:"row_effective_date" default:""`
	Row_expiry_date      *string `json:"row_expiry_date" default:""`
	Row_quality          *string `json:"row_quality" default:""`
}
