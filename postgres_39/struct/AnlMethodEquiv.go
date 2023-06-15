package dto

import (
	"time"
)

type Anl_method_equiv struct {
	Method_set_id1     string     `json:"method_set_id_1"`
	Method_id1         string     `json:"method_id_1"`
	Method_set_id2     string     `json:"method_set_id_2"`
	Method_id2         string     `json:"method_id_2"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Method_equiv_type  *string    `json:"method_equiv_type"`
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
