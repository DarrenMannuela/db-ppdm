package dto

type Restriction struct {
	Restriction_id      string  `json:"restriction_id" default:"source"`
	Restriction_version int     `json:"restriction_version" default:"1"`
	Active_ind          *string `json:"active_ind" default:""`
	Description         *string `json:"description" default:""`
	Effective_date      *string `json:"effective_date" default:""`
	End_date            *string `json:"end_date" default:""`
	End_date_event      *string `json:"end_date_event" default:""`
	Expiry_date         *string `json:"expiry_date" default:""`
	Ppdm_guid           *string `json:"ppdm_guid" default:""`
	Regulatory_act      *string `json:"regulatory_act" default:""`
	Remark              *string `json:"remark" default:""`
	Restriction_class   *string `json:"restriction_class" default:""`
	Restriction_name    *string `json:"restriction_name" default:""`
	Restriction_type    *string `json:"restriction_type" default:""`
	Rest_duration_type  *string `json:"rest_duration_type" default:""`
	Source              *string `json:"source" default:""`
	Source_document_id  *string `json:"source_document_id" default:""`
	Start_date          *string `json:"start_date" default:""`
	Start_date_event    *string `json:"start_date_event" default:""`
	Row_changed_by      *string `json:"row_changed_by" default:""`
	Row_changed_date    *string `json:"row_changed_date" default:""`
	Row_created_by      *string `json:"row_created_by" default:""`
	Row_created_date    *string `json:"row_created_date" default:""`
	Row_effective_date  *string `json:"row_effective_date" default:""`
	Row_expiry_date     *string `json:"row_expiry_date" default:""`
	Row_quality         *string `json:"row_quality" default:""`
}
