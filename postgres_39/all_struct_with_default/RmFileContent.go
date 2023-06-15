package dto

type Rm_file_content struct {
	File_id             string   `json:"file_id" default:"source"`
	Active_ind          *string  `json:"active_ind" default:""`
	Digital_format      *string  `json:"digital_format" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	File_content        *string  `json:"file_content" default:""`
	File_size           *float64 `json:"file_size" default:""`
	File_size_uom       *string  `json:"file_size_uom" default:""`
	Information_item_id *string  `json:"information_item_id" default:""`
	Info_item_subtype   *string  `json:"info_item_subtype" default:""`
	Physical_item_id    *string  `json:"physical_item_id" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Remark              *string  `json:"remark" default:""`
	Source              *string  `json:"source" default:""`
	Sw_application_id   *string  `json:"sw_application_id" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
