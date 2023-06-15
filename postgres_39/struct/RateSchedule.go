package dto

import (
	"time"
)

type Rate_schedule struct {
	Rate_schedule_id       string     `json:"rate_schedule_id"`
	Active_ind             *string    `json:"active_ind"`
	Area_id                *string    `json:"area_id"`
	Area_type              *string    `json:"area_type"`
	Change_notice          *string    `json:"change_notice"`
	Contract_id            *string    `json:"contract_id"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Owner_ba_add_obs_no    *int       `json:"owner_ba_add_obs_no"`
	Owner_ba_add_source    *string    `json:"owner_ba_add_source"`
	Owner_ba_id            *string    `json:"owner_ba_id"`
	Pay_method             *string    `json:"pay_method"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Preferred_currency_uom *string    `json:"preferred_currency_uom"`
	Provision_id           *string    `json:"provision_id"`
	Rate_schedule_type     *string    `json:"rate_schedule_type"`
	Remark                 *string    `json:"remark"`
	Review_frequency       *string    `json:"review_frequency"`
	Source                 *string    `json:"source"`
	Source_document_id     *string    `json:"source_document_id"`
	Spatial_description_id *string    `json:"spatial_description_id"`
	Spatial_obs_no         *int       `json:"spatial_obs_no"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
