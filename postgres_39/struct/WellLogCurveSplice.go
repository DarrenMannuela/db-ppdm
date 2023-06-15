package dto

import (
	"time"
)

type Well_log_curve_splice struct {
	Uwi                 string     `json:"uwi"`
	Curve_id            string     `json:"curve_id"`
	Splice_obs_no       int        `json:"splice_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Created_by_ba_id    *string    `json:"created_by_ba_id"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Information_item_id *string    `json:"information_item_id"`
	Info_item_subtype   *string    `json:"info_item_subtype"`
	Max_value_index     *float64   `json:"max_value_index"`
	Min_value_index     *float64   `json:"min_value_index"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Source_curve_id     *string    `json:"source_curve_id"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
