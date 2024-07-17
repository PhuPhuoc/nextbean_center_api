package model

type ReportOJT struct {
	Semester    string `json:"semester"`
	University  string `json:"university"`
	StartAt     string `json:"start-at"`
	EndAt       string `json:"end-at"`
	TotalIntern int    `json:"total-intern"`
}
