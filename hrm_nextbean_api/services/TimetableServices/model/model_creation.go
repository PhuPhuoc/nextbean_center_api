package model

type TimtableCreation struct {
	OfficeTime string `json:"office-time" validate:"required,type=date"`
	EstStartAt string `json:"est-start-at" validate:"required,type=date"`
	EstEndAt   string `json:"est-end-at" validate:"required,type=date"`
}
