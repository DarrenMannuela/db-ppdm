package dto

type Rm_info_item_origin struct {
	Info_item_subtype          string  `json:"info_item_subtype" default:"source"`
	Information_item_id        string  `json:"information_item_id" default:"source"`
	Origin_seq_no              int     `json:"origin_seq_no" default:"1"`
	Active_ind                 *string `json:"active_ind" default:""`
	Effective_date             *string `json:"effective_date" default:""`
	Expiry_date                *string `json:"expiry_date" default:""`
	Information_process        *string `json:"information_process" default:""`
	Parent_information_item_id *string `json:"parent_information_item_id" default:""`
	Parent_info_item_subtype   *string `json:"parent_info_item_subtype" default:""`
	Ppdm_guid                  *string `json:"ppdm_guid" default:""`
	Processor                  *string `json:"processor" default:""`
	Process_date               *string `json:"process_date" default:""`
	Remark                     *string `json:"remark" default:""`
	Source                     *string `json:"source" default:""`
	Version                    *string `json:"version" default:""`
	Row_changed_by             *string `json:"row_changed_by" default:""`
	Row_changed_date           *string `json:"row_changed_date" default:""`
	Row_created_by             *string `json:"row_created_by" default:""`
	Row_created_date           *string `json:"row_created_date" default:""`
	Row_effective_date         *string `json:"row_effective_date" default:""`
	Row_expiry_date            *string `json:"row_expiry_date" default:""`
	Row_quality                *string `json:"row_quality" default:""`
}
