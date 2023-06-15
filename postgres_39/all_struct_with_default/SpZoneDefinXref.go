package dto

type Sp_zone_defin_xref struct {
	Zone_definition_id_1 string  `json:"zone_definition_id_1" default:"source"`
	Zone_definition_id_2 string  `json:"zone_definition_id_2" default:"source"`
	Active_ind           *string `json:"active_ind" default:""`
	Description          *string `json:"description" default:""`
	Effective_date       *string `json:"effective_date" default:""`
	Expiry_date          *string `json:"expiry_date" default:""`
	Ppdm_guid            *string `json:"ppdm_guid" default:""`
	Remark               *string `json:"remark" default:""`
	Responsible_ba_id    *string `json:"responsible_ba_id" default:""`
	Source               *string `json:"source" default:""`
	Xref_reason          *string `json:"xref_reason" default:""`
	Row_changed_by       *string `json:"row_changed_by" default:""`
	Row_changed_date     *string `json:"row_changed_date" default:""`
	Row_created_by       *string `json:"row_created_by" default:""`
	Row_created_date     *string `json:"row_created_date" default:""`
	Row_effective_date   *string `json:"row_effective_date" default:""`
	Row_expiry_date      *string `json:"row_expiry_date" default:""`
	Row_quality          *string `json:"row_quality" default:""`
}
