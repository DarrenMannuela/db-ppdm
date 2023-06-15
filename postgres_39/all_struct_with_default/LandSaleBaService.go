package dto

type Land_sale_ba_service struct {
	Jurisdiction          string  `json:"jurisdiction" default:"source"`
	Land_sale_number      string  `json:"land_sale_number" default:"source"`
	Land_sale_offering_id string  `json:"land_sale_offering_id" default:"source"`
	Provided_by           string  `json:"provided_by" default:"source"`
	Service_seq_no        int     `json:"service_seq_no" default:"1"`
	Active_ind            *string `json:"active_ind" default:""`
	Ba_service_type       *string `json:"ba_service_type" default:""`
	Contact_ba_id         *string `json:"contact_ba_id" default:""`
	Description           *string `json:"description" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	End_date              *string `json:"end_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Organization_id       *string `json:"organization_id" default:""`
	Organization_seq_no   *int    `json:"organization_seq_no" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Provided_for_ba_id    *string `json:"provided_for_ba_id" default:""`
	Rate_schedule_id      *string `json:"rate_schedule_id" default:""`
	Remark                *string `json:"remark" default:""`
	Represented_ba_id     *string `json:"represented_ba_id" default:""`
	Service_quality       *string `json:"service_quality" default:""`
	Source                *string `json:"source" default:""`
	Start_date            *string `json:"start_date" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
