package dto

type Resent_revision_cat struct {
	Revision_category_id string  `json:"revision_category_id" default:"source"`
	Active_ind           *string `json:"active_ind" default:""`
	Category_name        *string `json:"category_name" default:""`
	Category_type        *string `json:"category_type" default:""`
	Effective_date       *string `json:"effective_date" default:""`
	Expiry_date          *string `json:"expiry_date" default:""`
	Gross_ind            *string `json:"gross_ind" default:""`
	Net_ind              *string `json:"net_ind" default:""`
	Part_of_category_id  *string `json:"part_of_category_id" default:""`
	Ppdm_guid            *string `json:"ppdm_guid" default:""`
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
