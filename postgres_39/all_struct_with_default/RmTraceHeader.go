package dto

type Rm_trace_header struct {
	Info_item_subtype   string   `json:"info_item_subtype" default:"source"`
	Information_item_id string   `json:"information_item_id" default:"source"`
	Header_id           string   `json:"header_id" default:"source"`
	Active_ind          *string  `json:"active_ind" default:""`
	Description         *string  `json:"description" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Header_format       *string  `json:"header_format" default:""`
	Header_offset       *float64 `json:"header_offset" default:""`
	Header_word         *string  `json:"header_word" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Remark              *string  `json:"remark" default:""`
	Source              *string  `json:"source" default:""`
	Word_length         *float64 `json:"word_length" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
