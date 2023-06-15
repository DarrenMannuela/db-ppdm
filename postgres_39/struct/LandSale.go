package dto

import (
	"time"
)

type Land_sale struct {
	Jurisdiction       string     `json:"jurisdiction"`
	Land_sale_number   string     `json:"land_sale_number"`
	Active_ind         *string    `json:"active_ind"`
	Close_date         *time.Time `json:"close_date"`
	Description        *string    `json:"description"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Land_sale_name     *string    `json:"land_sale_name"`
	Open_date          *time.Time `json:"open_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Publish_date       *time.Time `json:"publish_date"`
	Remark             *string    `json:"remark"`
	Sale_date          *time.Time `json:"sale_date"`
	Sold_size          *float64   `json:"sold_size"`
	Sold_size_ouom     *string    `json:"sold_size_ouom"`
	Source             *string    `json:"source"`
	Total_bonus        *float64   `json:"total_bonus"`
	Total_bonus_ouom   *string    `json:"total_bonus_ouom"`
	Total_size         *float64   `json:"total_size"`
	Total_size_ouom    *string    `json:"total_size_ouom"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
