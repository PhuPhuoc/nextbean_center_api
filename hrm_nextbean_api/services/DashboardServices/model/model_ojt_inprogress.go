package model

type DashboardOJTInProgress struct {
	Id          int    `json:"id"`
	Semester    string `json:"semester"`
	University  string `json:"university"`
	Status      string `json:"status"`
	TotalIntern int    `json:"total-intern"`
}
