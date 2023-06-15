package dto

import (
	"time"
)

type Work_order_instruction struct {
	Work_order_id          string     `json:"work_order_id"`
	Instruction_id         string     `json:"instruction_id"`
	Active_ind             *string    `json:"active_ind"`
	Ba_address_obs_no      *int       `json:"ba_address_obs_no"`
	Ba_address_source      *string    `json:"ba_address_source"`
	Ba_obs_no              *int       `json:"ba_obs_no"`
	Business_associate_id  *string    `json:"business_associate_id"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Instruction_desc       *string    `json:"instruction_desc"`
	Instruction_type       *string    `json:"instruction_type"`
	Instruction_value      *float64   `json:"instruction_value"`
	Instruction_value_ouom *string    `json:"instruction_value_ouom"`
	Instruction_value_uom  *string    `json:"instruction_value_uom"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
