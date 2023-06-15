package dto

type Business_associate struct {
	Business_associate_id string  `json:"business_associate_id" default:"source"`
	Active_ind            *string `json:"active_ind" default:""`
	Ba_abbreviation       *string `json:"ba_abbreviation" default:""`
	Ba_category           *string `json:"ba_category" default:""`
	Ba_code               *string `json:"ba_code" default:""`
	Ba_long_name          *string `json:"ba_long_name" default:""`
	Ba_short_name         *string `json:"ba_short_name" default:""`
	Ba_type               *string `json:"ba_type" default:""`
	Credit_check_date     *string `json:"credit_check_date" default:""`
	Credit_check_ind      *string `json:"credit_check_ind" default:""`
	Credit_check_source   *string `json:"credit_check_source" default:""`
	Credit_rating         *string `json:"credit_rating" default:""`
	Credit_rating_source  *string `json:"credit_rating_source" default:""`
	Current_status        *string `json:"current_status" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	First_name            *string `json:"first_name" default:""`
	Last_name             *string `json:"last_name" default:""`
	Middle_initial        *string `json:"middle_initial" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Remark                *string `json:"remark" default:""`
	Service_period        *string `json:"service_period" default:""`
	Source                *string `json:"source" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
