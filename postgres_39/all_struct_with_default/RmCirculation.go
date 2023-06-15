package dto

type Rm_circulation struct {
	Circ_id             string  `json:"circ_id" default:"source"`
	Active_ind          *string `json:"active_ind" default:""`
	Authorized_by       *string `json:"authorized_by" default:""`
	Checked_out_by      *string `json:"checked_out_by" default:""`
	Circ_due_date       *string `json:"circ_due_date" default:""`
	Circ_in_date        *string `json:"circ_in_date" default:""`
	Circ_out_date       *string `json:"circ_out_date" default:""`
	Condition_in        *string `json:"condition_in" default:""`
	Condition_out       *string `json:"condition_out" default:""`
	Data_circ_status    *string `json:"data_circ_status" default:""`
	Data_content_seq_no *int    `json:"data_content_seq_no" default:""`
	Effective_date      *string `json:"effective_date" default:""`
	Expiry_date         *string `json:"expiry_date" default:""`
	Information_item_id *string `json:"information_item_id" default:""`
	Info_item_subtype   *string `json:"info_item_subtype" default:""`
	Physical_item_id    *string `json:"physical_item_id" default:""`
	Ppdm_guid           *string `json:"ppdm_guid" default:""`
	Project_id          *string `json:"project_id" default:""`
	Project_step_id     *string `json:"project_step_id" default:""`
	Reference_num       *string `json:"reference_num" default:""`
	Remark              *string `json:"remark" default:""`
	Reserved_by         *string `json:"reserved_by" default:""`
	Reserved_ind        *string `json:"reserved_ind" default:""`
	Source              *string `json:"source" default:""`
	Status_date         *string `json:"status_date" default:""`
	Row_changed_by      *string `json:"row_changed_by" default:""`
	Row_changed_date    *string `json:"row_changed_date" default:""`
	Row_created_by      *string `json:"row_created_by" default:""`
	Row_created_date    *string `json:"row_created_date" default:""`
	Row_effective_date  *string `json:"row_effective_date" default:""`
	Row_expiry_date     *string `json:"row_expiry_date" default:""`
	Row_quality         *string `json:"row_quality" default:""`
}
