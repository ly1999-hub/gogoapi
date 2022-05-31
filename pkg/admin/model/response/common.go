package responsemodel

// ResponseCreate ...
type ResponseCreate struct {
	ID string `json:"id"`
}

// ResponseUpdate ...
type ResponseUpdate struct {
	ID string `json:"id"`
}

// ResponseDelete ...
type ResponseDelete struct {
	ID string `json:"id"`
}

// ResponseList ...
type ResponseList struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
	Limit int64       `json:"limit"`
}