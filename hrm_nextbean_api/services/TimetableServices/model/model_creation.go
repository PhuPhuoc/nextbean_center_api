package model

type TimtableCreation struct {
	OfficeTime string `json:"office-time" validate:"required,type=date"`
	EstStart   string `json:"est-start" validate:"required,type=time"`
	EstEnd     string `json:"est-end" validate:"required,type=time"`
}
