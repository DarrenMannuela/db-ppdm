package dto

type Cs_coord_transform struct {
	Transform_id           string  `json:"transform_id" default:"source"`
	Active_ind             *string `json:"active_ind" default:""`
	Effective_date         *string `json:"effective_date" default:""`
	Expiry_date            *string `json:"expiry_date" default:""`
	Ppdm_guid              *string `json:"ppdm_guid" default:""`
	Preferred_ind          *string `json:"preferred_ind" default:""`
	Remark                 *string `json:"remark" default:""`
	Source                 *string `json:"source" default:""`
	Source_coord_system_id *string `json:"source_coord_system_id" default:""`
	Source_document_id     *string `json:"source_document_id" default:""`
	Target_coord_system_id *string `json:"target_coord_system_id" default:""`
	Transform_name         *string `json:"transform_name" default:""`
	Transform_type         *string `json:"transform_type" default:""`
	Row_changed_by         *string `json:"row_changed_by" default:""`
	Row_changed_date       *string `json:"row_changed_date" default:""`
	Row_created_by         *string `json:"row_created_by" default:""`
	Row_created_date       *string `json:"row_created_date" default:""`
	Row_effective_date     *string `json:"row_effective_date" default:""`
	Row_expiry_date        *string `json:"row_expiry_date" default:""`
	Row_quality            *string `json:"row_quality" default:""`
}
