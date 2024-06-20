package model

type Daily struct {
	WeekDay       string `json:"week-day"`
	Date          string `json:"date"`
	TotalApproved int    `json:"total-approved"`
	TotalDenied   int    `json:"total-denied"`
	TotalWaiting  int    `json:"total-waiting"`
}
