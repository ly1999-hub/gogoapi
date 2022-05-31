package apimodel

// UserAll ...
type UserAll struct {
	Page    int64  `query:"page"`
	Limit   int64  `query:"limit"`
	Keyword string `query:"keyword"`
	Banned  string `query:"banned"`
	FromAt  string `query:"fromAt"`
	ToAt    string `query:"toAt"`
}
