package utils

import "time"

func CreateDateTimeCurrentFormated() string {
	currentTime := time.Now()
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	currentTimeInLocal := currentTime.In(loc)
	formattedTime := currentTimeInLocal.Format("2006-01-02 15:04:05")

	return formattedTime
}
