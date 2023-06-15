package dto

import (
	"time"
)

type Rm_circulation struct {
	Circ_id             string     `json:"circ_id"`
	Active_ind          *string    `json:"active_ind"`
	Authorized_by       *string    `json:"authorized_by"`
	Checked_out_by      *string    `json:"checked_out_by"`
	Circ_due_date       *time.Time `json:"circ_due_date"`
	Circ_in_date        *time.Time `json:"circ_in_date"`
	Circ_out_date       *time.Time `json:"circ_out_date"`
	Condition_in        *string    `json:"condition_in"`
	Condition_out       *string    `json:"condition_out"`
	Data_circ_status    *string    `json:"data_circ_status"`
	Data_content_seq_no *int       `json:"data_content_seq_no"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Information_item_id *string    `json:"information_item_id"`
	Info_item_subtype   *string    `json:"info_item_subtype"`
	Physical_item_id    *string    `json:"physical_item_id"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Project_id          *string    `json:"project_id"`
	Project_step_id     *string    `json:"project_step_id"`
	Reference_num       *string    `json:"reference_num"`
	Remark              *string    `json:"remark"`
	Reserved_by         *string    `json:"reserved_by"`
	Reserved_ind        *string    `json:"reserved_ind"`
	Source              *string    `json:"source"`
	Status_date         *time.Time `json:"status_date"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
