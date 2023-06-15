package dto

type Instrument_xref struct {
	Instrument_id        string  `json:"instrument_id" default:"source"`
	Parent_instrument_id string  `json:"parent_instrument_id" default:"source"`
	Active_ind           *string `json:"active_ind" default:""`
	Ammendment_desc      *string `json:"ammendment_desc" default:""`
	Effective_date       *string `json:"effective_date" default:""`
	Expiry_date          *string `json:"expiry_date" default:""`
	Ppdm_guid            *string `json:"ppdm_guid" default:""`
	Remark               *string `json:"remark" default:""`
	Source               *string `json:"source" default:""`
	Xref_type            *string `json:"xref_type" default:""`
	Row_changed_by       *string `json:"row_changed_by" default:""`
	Row_changed_date     *string `json:"row_changed_date" default:""`
	Row_created_by       *string `json:"row_created_by" default:""`
	Row_created_date     *string `json:"row_created_date" default:""`
	Row_effective_date   *string `json:"row_effective_date" default:""`
	Row_expiry_date      *string `json:"row_expiry_date" default:""`
	Row_quality          *string `json:"row_quality" default:""`
}
