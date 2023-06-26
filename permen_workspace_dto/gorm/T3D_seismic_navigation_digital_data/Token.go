package dto

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	ExpiryDate  string `json:"expiry_date"`
	Affiliation string `json:"affiliation"`
}
