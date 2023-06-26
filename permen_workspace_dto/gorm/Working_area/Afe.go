package dto

type Afe struct {
	Afe_number      int     `json:"afe_number" default:"1"`
	Workspace_name  *string `json:"workspace_name" default:""`
	Kkks_name       *string `json:"kkks_name" default:""`
	Working_area    *string `json:"working_area" default:""`
	Submission_type *string `json:"submission_type" default:""`
	Data_type       string  `json:"data_type" default:"Print Well Report"`
	Email           string  `json:"email" default:"lorem.ipsum@gmail.com"`
}
