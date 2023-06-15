package dto

import (
	"time"
)

type Rm_info_item_group struct {
	Group_info_item_subtype   string     `json:"group_info_item_subtype"`
	Group_information_item_id string     `json:"group_information_item_id"`
	Info_item_subtype         string     `json:"info_item_subtype"`
	Information_item_id       string     `json:"information_item_id"`
	Group_obs_no              int        `json:"group_obs_no"`
	Active_ind                *string    `json:"active_ind"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Group_seq_no              *int       `json:"group_seq_no"`
	Group_type                *string    `json:"group_type"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Remark                    *string    `json:"remark"`
	Source                    *string    `json:"source"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
