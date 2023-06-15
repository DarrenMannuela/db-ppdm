package dto

import (
	"time"
)

type Land_remark struct {
	Land_right_subtype string     `json:"land_right_subtype"`
	Land_right_id      string     `json:"land_right_id"`
	Remark_id          string     `json:"remark_id"`
	Remark_seq_no      int        `json:"remark_seq_no"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Recommend_impl_ind *string    `json:"recommend_impl_ind"`
	Remark             *string    `json:"remark"`
	Remark_ba_id       *string    `json:"remark_ba_id"`
	Remark_date        *time.Time `json:"remark_date"`
	Remark_type        *string    `json:"remark_type"`
	Remark_user        *string    `json:"remark_user"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
