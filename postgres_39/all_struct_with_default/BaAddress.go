package dto

type Ba_address struct {
	Business_associate_id    string  `json:"business_associate_id" default:"source"`
	Source                   string  `json:"source" default:"source"`
	Address_obs_no           int     `json:"address_obs_no" default:"1"`
	Active_ind               *string `json:"active_ind" default:""`
	Addressee_text           *string `json:"addressee_text" default:""`
	Address_type             *string `json:"address_type" default:""`
	City_area_id             *string `json:"city_area_id" default:""`
	City_area_type           *string `json:"city_area_type" default:""`
	Country_area_id          *string `json:"country_area_id" default:""`
	Country_area_type        *string `json:"country_area_type" default:""`
	Effective_date           *string `json:"effective_date" default:""`
	Expiry_date              *string `json:"expiry_date" default:""`
	First_address_line       *string `json:"first_address_line" default:""`
	Office_type              *string `json:"office_type" default:""`
	Postal_zip_code          *string `json:"postal_zip_code" default:""`
	Ppdm_guid                *string `json:"ppdm_guid" default:""`
	Preferred_ind            *string `json:"preferred_ind" default:""`
	Province_state_area_id   *string `json:"province_state_area_id" default:""`
	Province_state_area_type *string `json:"province_state_area_type" default:""`
	Remark                   *string `json:"remark" default:""`
	Second_address_line      *string `json:"second_address_line" default:""`
	Service_period           *string `json:"service_period" default:""`
	Service_quality          *string `json:"service_quality" default:""`
	Third_address_line       *string `json:"third_address_line" default:""`
	Withholding_tax_ind      *string `json:"withholding_tax_ind" default:""`
	Row_changed_by           *string `json:"row_changed_by" default:""`
	Row_changed_date         *string `json:"row_changed_date" default:""`
	Row_created_by           *string `json:"row_created_by" default:""`
	Row_created_date         *string `json:"row_created_date" default:""`
	Row_effective_date       *string `json:"row_effective_date" default:""`
	Row_expiry_date          *string `json:"row_expiry_date" default:""`
	Row_quality              *string `json:"row_quality" default:""`
}
