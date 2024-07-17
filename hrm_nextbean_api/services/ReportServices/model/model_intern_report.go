package model

type ReportIntern struct {
	Name                  string
	Code                  string
	ProjectsParticipated  int
	TotalTasks            int
	EstimatedTime         int
	ActualTime            int
	TotalDaysWorkOffline  int
	TotalTimesWorkOffline int
}
