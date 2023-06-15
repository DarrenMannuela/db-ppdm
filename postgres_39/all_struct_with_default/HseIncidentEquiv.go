package dto

type Hse_incident_equiv struct {
	Incident_set_id    string  `json:"incident_set_id" default:"source"`
	Incident_type_id   string  `json:"incident_type_id" default:"source"`
	Incident_set_id2   string  `json:"incident_set_id_2" default:"source"`
	Incident_type_id2  string  `json:"incident_type_id_2" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Description        *string `json:"description" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
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
