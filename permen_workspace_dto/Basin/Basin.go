package dto

type Basin struct{

Project_name       *string   `json:"project_name" default:""`
Strat_name_set_id       *string   `json:"strat_name_set_id" default:""`
Strat_status       *string   `json:"strat_status" default:""`
Product_type       *string   `json:"product_type" default:""`
Reserve_class_id       *string   `json:"reserve_class_id" default:""`
Open_balance       *string   `json:"open_balance" default:""`
Open_balance_ouom       *string   `json:"open_balance_ouom" default:""`
Size_type       *string   `json:"size_type" default:""`
Gross_size       *string   `json:"gross_size" default:""`
Size_ouom       *string   `json:"size_ouom" default:""`
Strat_type       *string   `json:"strat_type" default:""`
Fault_type       *string   `json:"fault_type" default:""`
Source       *string   `json:"source" default:""`
Qc_status       *string   `json:"qc_status" default:""`
Checked_by_ba_id       *string   `json:"checked_by_ba_id" default:""`
}