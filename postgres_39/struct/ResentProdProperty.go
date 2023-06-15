package dto

import (
	"time"
)

type Resent_prod_property struct {
	Resent_id            string     `json:"resent_id"`
	Reserve_class_id     string     `json:"reserve_class_id"`
	Property_seq_no      int        `json:"property_seq_no"`
	Active_ind           *string    `json:"active_ind"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Heat_content         *float64   `json:"heat_content"`
	Heat_content_ouom    *string    `json:"heat_content_ouom"`
	Loss_factor          *float64   `json:"loss_factor"`
	Oil_density          *float64   `json:"oil_density"`
	Oil_density_ouom     *string    `json:"oil_density_ouom"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Product_type         *string    `json:"product_type"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Sulphur_content      *float64   `json:"sulphur_content"`
	Sulphur_content_ouom *string    `json:"sulphur_content_ouom"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
