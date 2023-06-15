package dto

type Strat_column struct {
	Strat_column_id       string  `json:"strat_column_id" default:"source"`
	Source                string  `json:"source" default:"source"`
	Active_ind            *string `json:"active_ind" default:""`
	Area_id               *string `json:"area_id" default:""`
	Area_type             *string `json:"area_type" default:""`
	Business_associate_id *string `json:"business_associate_id" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Remark                *string `json:"remark" default:""`
	Source_document_id    *string `json:"source_document_id" default:""`
	Strat_column_name     *string `json:"strat_column_name" default:""`
	Strat_column_type     *string `json:"strat_column_type" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
