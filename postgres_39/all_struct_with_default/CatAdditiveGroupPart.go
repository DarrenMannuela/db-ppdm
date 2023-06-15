package dto

type Cat_additive_group_part struct {
	Additive_group_id     string  `json:"additive_group_id" default:"source"`
	Additive_type_id      string  `json:"additive_type_id" default:"source"`
	Additive_part_obs_no  int     `json:"additive_part_obs_no" default:"1"`
	Active_ind            *string `json:"active_ind" default:""`
	Catalogue_additive_id *string `json:"catalogue_additive_id" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Long_name             *string `json:"long_name" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Remark                *string `json:"remark" default:""`
	Short_name            *string `json:"short_name" default:""`
	Source                *string `json:"source" default:""`
	Source_document_id    *string `json:"source_document_id" default:""`
	Substance_id          *string `json:"substance_id" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
