package dto

type Land_sale_bid_set struct {
	Land_sale_bid_set_id string  `json:"land_sale_bid_set_id" default:"source"`
	Active_ind           *string `json:"active_ind" default:""`
	Bid_status           *string `json:"bid_status" default:""`
	Confidential_ind     *string `json:"confidential_ind" default:""`
	Contingency_desc     *string `json:"contingency_desc" default:""`
	Contingency_ind      *string `json:"contingency_ind" default:""`
	Effective_date       *string `json:"effective_date" default:""`
	Expiry_date          *string `json:"expiry_date" default:""`
	Owner_ba_id          *string `json:"owner_ba_id" default:""`
	Ppdm_guid            *string `json:"ppdm_guid" default:""`
	Remark               *string `json:"remark" default:""`
	Set_status_date      *string `json:"set_status_date" default:""`
	Source               *string `json:"source" default:""`
	Successful_ind       *string `json:"successful_ind" default:""`
	Row_changed_by       *string `json:"row_changed_by" default:""`
	Row_changed_date     *string `json:"row_changed_date" default:""`
	Row_created_by       *string `json:"row_created_by" default:""`
	Row_created_date     *string `json:"row_created_date" default:""`
	Row_effective_date   *string `json:"row_effective_date" default:""`
	Row_expiry_date      *string `json:"row_expiry_date" default:""`
	Row_quality          *string `json:"row_quality" default:""`
}
