package model

type Timtable struct {
	Id                string `json:"id"`
	InternName        string `json:"intern_name"`
	StudentCode       string `json:"student-code"`
	OfficeTime        string `json:"office-time"`
	Verified          string `json:"verified"`
	EstStartTime      string `json:"est-start-time"`
	EstEndTime        string `json:"est-end-time"`
	ActClockin        string `json:"act-clockin"`
	ClockinValidated  string `json:"clockin-validated"`
	ActClockout       string `json:"act-clockout"`
	ClockoutValidated string `json:"clockout-validated"`
	StatusAttendance  string `json:"status-attendance"`
}
