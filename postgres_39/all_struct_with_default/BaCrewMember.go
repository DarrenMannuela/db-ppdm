package dto

type Ba_crew_member struct {
	Crew_company_ba_id string  `json:"crew_company_ba_id" default:"source"`
	Crew_id            string  `json:"crew_id" default:"source"`
	Member_obs_no      int     `json:"member_obs_no" default:"1"`
	Crew_member_ba_id  string  `json:"crew_member_ba_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Crew_position      *string `json:"crew_position" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
