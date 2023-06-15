package dto

type Paleo_obs_qualifier struct {
	Qualifier_id         string  `json:"qualifier_id" default:"source"`
	Active_ind           *string `json:"active_ind" default:""`
	Effective_date       *string `json:"effective_date" default:""`
	Expiry_date          *string `json:"expiry_date" default:""`
	Ppdm_guid            *string `json:"ppdm_guid" default:""`
	Qualifier_code       *string `json:"qualifier_code" default:""`
	Qualifier_long_name  *string `json:"qualifier_long_name" default:""`
	Qualifier_owner      *string `json:"qualifier_owner" default:""`
	Qualifier_short_name *string `json:"qualifier_short_name" default:""`
	Qualifier_type       *string `json:"qualifier_type" default:""`
	Remark               *string `json:"remark" default:""`
	Source               *string `json:"source" default:""`
	Row_changed_by       *string `json:"row_changed_by" default:""`
	Row_changed_date     *string `json:"row_changed_date" default:""`
	Row_created_by       *string `json:"row_created_by" default:""`
	Row_created_date     *string `json:"row_created_date" default:""`
	Row_effective_date   *string `json:"row_effective_date" default:""`
	Row_expiry_date      *string `json:"row_expiry_date" default:""`
	Row_quality          *string `json:"row_quality" default:""`
}
