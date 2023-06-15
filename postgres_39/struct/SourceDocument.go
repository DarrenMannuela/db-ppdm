package dto

import (
	"time"
)

type Source_document struct {
	Source_document_id string     `json:"source_document_id"`
	Abbreviation       *string    `json:"abbreviation"`
	Active_ind         *string    `json:"active_ind"`
	Document_title     *string    `json:"document_title"`
	Document_type      *string    `json:"document_type"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Figure_reference   *string    `json:"figure_reference"`
	Issue              *string    `json:"issue"`
	Language           *string    `json:"language"`
	Page_reference     *string    `json:"page_reference"`
	Plate_reference    *string    `json:"plate_reference"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Publication        *string    `json:"publication"`
	Publication_date   *time.Time `json:"publication_date"`
	Publication_year   *int       `json:"publication_year"`
	Publisher          *string    `json:"publisher"`
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
