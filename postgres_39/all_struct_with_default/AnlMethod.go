package dto

type Anl_method struct {
	Method_set_id      string  `json:"method_set_id" default:"source"`
	Method_id          string  `json:"method_id" default:"source"`
	Abbreviation       *string `json:"abbreviation" default:""`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Long_name          *string `json:"long_name" default:""`
	Method_desc        *string `json:"method_desc" default:""`
	Method_seq_no      *int    `json:"method_seq_no" default:""`
	Method_type        *string `json:"method_type" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Preparation_class  *string `json:"preparation_class" default:""`
	Remark             *string `json:"remark" default:""`
	Short_name         *string `json:"short_name" default:""`
	Source             *string `json:"source" default:""`
	Source_document_id *string `json:"source_document_id" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
