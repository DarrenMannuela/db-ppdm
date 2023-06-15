package dto

type Well_license_violation struct {
	Uwi                string  `json:"uwi" default:"source"`
	Well_lic_source    string  `json:"well_lic_source" default:"source"`
	License_id         string  `json:"license_id" default:"source"`
	Violation_id       string  `json:"violation_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Condition_id       *string `json:"condition_id" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Resolve_date       *string `json:"resolve_date" default:""`
	Resolve_desc       *string `json:"resolve_desc" default:""`
	Resolve_type       *string `json:"resolve_type" default:""`
	Source             *string `json:"source" default:""`
	Violation_date     *string `json:"violation_date" default:""`
	Violation_desc     *string `json:"violation_desc" default:""`
	Violation_type     *string `json:"violation_type" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
