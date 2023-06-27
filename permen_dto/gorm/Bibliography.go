package dto

type Bibliography struct{

Id      int  `json:"id" default:"" gorm:"primaryKey"`
Publisher   *string  `json:"publisher" default:""`
Document_title   *string  `json:"document_title" default:""`
Issue   *string  `json:"issue" default:""`
Author_id   *string  `json:"author_id" default:""`
Publication_date   *string  `json:"publication_date" default:""`
Document_type   *string  `json:"document_type" default:""`
Data_store_name   *string  `json:"data_store_name" default:""`
}