package model

type Today struct {
	Checkin  string `json:"checkin,omitempty"`
	Checkout string `json:"checkout,omitempty"`
}
