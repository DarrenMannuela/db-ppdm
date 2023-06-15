package dto

type Source_doc_author struct {
	Source_document_id string  `json:"source_document_id" default:"source"`
	Author_id          string  `json:"author_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Author_ba_id       *string `json:"author_ba_id" default:""`
	Author_first_name  *string `json:"author_first_name" default:""`
	Author_initial     *string `json:"author_initial" default:""`
	Author_last_name   *string `json:"author_last_name" default:""`
	Author_seq_no      *int    `json:"author_seq_no" default:""`
	Author_type        *string `json:"author_type" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
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
