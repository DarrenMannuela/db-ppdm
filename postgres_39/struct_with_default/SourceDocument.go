package dto

type Source_document struct {
	Source_document_id string  `json:"source_document_id" default:"source"`
	Abbreviation       *string `json:"abbreviation" default:""`
	Active_ind         *string `json:"active_ind" default:""`
	Document_title     *string `json:"document_title" default:""`
	Document_type      *string `json:"document_type" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Figure_reference   *string `json:"figure_reference" default:""`
	Issue              *string `json:"issue" default:""`
	Language           *string `json:"language" default:""`
	Page_reference     *string `json:"page_reference" default:""`
	Plate_reference    *string `json:"plate_reference" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Publication        *string `json:"publication" default:""`
	Publication_date   *string `json:"publication_date" default:""`
	Publication_year   *int    `json:"publication_year" default:""`
	Publisher          *string `json:"publisher" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
