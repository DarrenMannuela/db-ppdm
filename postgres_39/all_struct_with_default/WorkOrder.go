package dto

type Work_order struct {
	Work_order_id      string  `json:"work_order_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Complete_date      *string `json:"complete_date" default:""`
	Due_date           *string `json:"due_date" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Final_billing_date *string `json:"final_billing_date" default:""`
	Instructions       *string `json:"instructions" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Request_date       *string `json:"request_date" default:""`
	Rush_ind           *string `json:"rush_ind" default:""`
	Source             *string `json:"source" default:""`
	Work_order_number  *string `json:"work_order_number" default:""`
	Work_order_type    *string `json:"work_order_type" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
