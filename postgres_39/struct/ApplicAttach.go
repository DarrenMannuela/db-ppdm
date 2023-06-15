package dto

import (
	"time"
)

type Applic_attach struct {
	Application_id         string     `json:"application_id"`
	Attachment_id          string     `json:"attachment_id"`
	Active_ind             *string    `json:"active_ind"`
	Attachment_description *string    `json:"attachment_description"`
	Attachment_type        *string    `json:"attachment_type"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Physical_item_id       *string    `json:"physical_item_id"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Received_date          *time.Time `json:"received_date"`
	Received_ind           *string    `json:"received_ind"`
	Remark                 *string    `json:"remark"`
	Sent_by                *string    `json:"sent_by"`
	Sent_date              *time.Time `json:"sent_date"`
	Sent_ind               *string    `json:"sent_ind"`
	Source                 *string    `json:"source"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
