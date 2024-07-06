package model

type TimeTableFilter struct {
	Id               string `json:"id,omitempty"`
	InternName       string `json:"intern-name,omitempty"`
	StudentCode      string `json:"student-code,omitempty"`
	OfficeTimeFrom   string `json:"office-time-from,omitempty"`
	OfficeTimeTo     string `json:"office-time-to,omitempty"`
	Verified         string `json:"verified,omitempty"`
	StatusAttendance string `json:"status-attendance,omitempty"`
	OrderBy          string `json:"order-by,omitempty"`
	Role             string `json:"-"`
	Account_id       string `json:"-"`
}
