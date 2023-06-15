package dto

import (
	"time"
)

type Ba_address struct {
	Business_associate_id    string     `json:"business_associate_id"`
	Source                   string     `json:"source"`
	Address_obs_no           int        `json:"address_obs_no"`
	Active_ind               *string    `json:"active_ind"`
	Addressee_text           *string    `json:"addressee_text"`
	Address_type             *string    `json:"address_type"`
	City_area_id             *string    `json:"city_area_id"`
	City_area_type           *string    `json:"city_area_type"`
	Country_area_id          *string    `json:"country_area_id"`
	Country_area_type        *string    `json:"country_area_type"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	First_address_line       *string    `json:"first_address_line"`
	Office_type              *string    `json:"office_type"`
	Postal_zip_code          *string    `json:"postal_zip_code"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Preferred_ind            *string    `json:"preferred_ind"`
	Province_state_area_id   *string    `json:"province_state_area_id"`
	Province_state_area_type *string    `json:"province_state_area_type"`
	Remark                   *string    `json:"remark"`
	Second_address_line      *string    `json:"second_address_line"`
	Service_period           *string    `json:"service_period"`
	Service_quality          *string    `json:"service_quality"`
	Third_address_line       *string    `json:"third_address_line"`
	Withholding_tax_ind      *string    `json:"withholding_tax_ind"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
