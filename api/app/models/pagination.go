package models

type Pagination struct {
	HasPrev    bool `json:"hasPrev"`
	HasNext    bool `json:"hasNext"`
	PerPage    int  `json:"perPage"`
	Total      int  `json:"total"`
	TotalPages int  `json:"totalPages"`
}

func (p *Pagination) Prepare(phones []Phone, page int) {
	p.Total = len(phones)
	p.PerPage = 10
	p.TotalPages = 0
	p.HasPrev = false
	p.HasNext = false

	if p.Total > 0 {
		if p.Total < p.PerPage {
			p.PerPage = p.Total
		}

		p.TotalPages = p.Total / p.PerPage

		if page != 1 {
			p.HasPrev = true
		}

		if page >= p.Total/p.PerPage {
			p.HasNext = false
		} else {
			p.HasNext = true
		}
	}

}
