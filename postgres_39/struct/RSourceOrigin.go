package dto

import (
	"time"
)

type R_source_origin struct {
	Source             string     `json:"source"`
	Origin_obs_no      int        `json:"origin_obs_no"`
	Abbreviation       *string    `json:"abbreviation"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Long_name          *string    `json:"long_name"`
	Owner_ba_id        *string    `json:"owner_ba_id"`
	Physical_item_id   *string    `json:"physical_item_id"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Project_id         *string    `json:"project_id"`
	Remark             *string    `json:"remark"`
	Row_source         *string    `json:"row_source"`
	Short_name         *string    `json:"short_name"`
	Source_document    *string    `json:"source_document"`
	Sw_application_id  *string    `json:"sw_application_id"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
