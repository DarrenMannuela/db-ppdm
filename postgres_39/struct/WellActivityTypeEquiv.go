package dto

import (
	"time"
)

type Well_activity_type_equiv struct {
	Activity_set_id    string     `json:"activity_set_id"`
	Activity_type_id   string     `json:"activity_type_id"`
	Activity_set_id2   string     `json:"activity_set_id_2"`
	Activity_type_id2  string     `json:"activity_type_id_2"`
	Equiv_obs_no       int        `json:"equiv_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Equiv_owner_ba_id  *string    `json:"equiv_owner_ba_id"`
	Equiv_type         *string    `json:"equiv_type"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Full_name          *string    `json:"full_name"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
