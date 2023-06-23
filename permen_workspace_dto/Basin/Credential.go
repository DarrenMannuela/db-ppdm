package dto

type Credential struct {
	GrantType string `form:"grant_type"`
	Username  string `form:"username"`
	Password  string `form:"password"`
}
