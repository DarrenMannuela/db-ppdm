package dto

type Rm_thesaurus_word_xref struct {
	Thesaurus_id1            string  `json:"thesaurus_id_1" default:"source"`
	Thesaurus_word_id1       string  `json:"thesaurus_word_id_1" default:"source"`
	Thesaurus_id2            string  `json:"thesaurus_id_2" default:"source"`
	Thesaurus_word_id2       string  `json:"thesaurus_word_id_2" default:"source"`
	Xref_obs_no              int     `json:"xref_obs_no" default:"1"`
	Active_ind               *string `json:"active_ind" default:""`
	Effective_date           *string `json:"effective_date" default:""`
	Expiry_date              *string `json:"expiry_date" default:""`
	Ppdm_guid                *string `json:"ppdm_guid" default:""`
	Relationship_description *string `json:"relationship_description" default:""`
	Remark                   *string `json:"remark" default:""`
	Source                   *string `json:"source" default:""`
	Source_document_id       *string `json:"source_document_id" default:""`
	Thesaurus_xref_type      *string `json:"thesaurus_xref_type" default:""`
	Row_changed_by           *string `json:"row_changed_by" default:""`
	Row_changed_date         *string `json:"row_changed_date" default:""`
	Row_created_by           *string `json:"row_created_by" default:""`
	Row_created_date         *string `json:"row_created_date" default:""`
	Row_effective_date       *string `json:"row_effective_date" default:""`
	Row_expiry_date          *string `json:"row_expiry_date" default:""`
	Row_quality              *string `json:"row_quality" default:""`
}
