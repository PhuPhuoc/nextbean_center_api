package common

type Pagination struct {
	Page  int   `json:"page"`
	PSize int   `json:"psize"`
	Items int64 `json:"items"`
	Pages int64 `json:"pages"`
}

func (p *Pagination) Process() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PSize <= 0 || p.PSize >= 100 {
		p.PSize = 10
	}
}
