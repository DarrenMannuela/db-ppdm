package dto

import (
	"time"
)

type Ppdm_rule_detail struct {
	Rule_id                string     `json:"rule_id"`
	Detail_seq_no          int        `json:"detail_seq_no"`
	Active_ind             *string    `json:"active_ind"`
	Average_value          *float64   `json:"average_value"`
	Average_value_ouom     *string    `json:"average_value_ouom"`
	Average_value_uom      *string    `json:"average_value_uom"`
	Boolean_rule           *string    `json:"boolean_rule"`
	Business_rule          *string    `json:"business_rule"`
	Date_format_desc       *time.Time `json:"date_format_desc"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Max_date               *time.Time `json:"max_date"`
	Max_value              *float64   `json:"max_value"`
	Max_value_ouom         *string    `json:"max_value_ouom"`
	Max_value_uom          *string    `json:"max_value_uom"`
	Min_date               *time.Time `json:"min_date"`
	Min_value              *float64   `json:"min_value"`
	Min_value_ouom         *string    `json:"min_value_ouom"`
	Min_value_uom          *string    `json:"min_value_uom"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Referenced_column_name *string    `json:"referenced_column_name"`
	Reference_column_name2 *string    `json:"reference_column_name_2"`
	Reference_system_id    *string    `json:"reference_system_id"`
	Reference_table_name   *string    `json:"reference_table_name"`
	Reference_table_name2  *string    `json:"reference_table_name_2"`
	Reference_value        *float64   `json:"reference_value"`
	Reference_value_ouom   *string    `json:"reference_value_ouom"`
	Reference_value_type   *string    `json:"reference_value_type"`
	Reference_value_uom    *string    `json:"reference_value_uom"`
	Remark                 *string    `json:"remark"`
	Rule_desc              *string    `json:"rule_desc"`
	Rule_detail_type       *string    `json:"rule_detail_type"`
	Source                 *string    `json:"source"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
