package dto

type Facility_equipment struct {
	Facility_id        string  `json:"facility_id" default:"source"`
	Facility_type      string  `json:"facility_type" default:"source"`
	Equipment_id       string  `json:"equipment_id" default:"source"`
	Install_obs_no     int     `json:"install_obs_no" default:"1"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Install_date       *string `json:"install_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Remove_date        *string `json:"remove_date" default:""`
	Remove_reason      *string `json:"remove_reason" default:""`
	Source             *string `json:"source" default:""`
	Use_description    *string `json:"use_description" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
