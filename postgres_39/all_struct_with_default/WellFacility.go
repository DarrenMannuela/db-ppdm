package dto

type Well_facility struct {
	Uwi                string  `json:"uwi" default:"source"`
	Facility_id        string  `json:"facility_id" default:"source"`
	Facility_type      string  `json:"facility_type" default:"source"`
	Active_obs_no      int     `json:"active_obs_no" default:"1"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Facility_use_type  *string `json:"facility_use_type" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Single_source_ind  *string `json:"single_source_ind" default:""`
	Source             *string `json:"source" default:""`
	String_id          *string `json:"string_id" default:""`
	String_source      *string `json:"string_source" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
