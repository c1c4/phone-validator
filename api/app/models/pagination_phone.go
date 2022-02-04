package models

type PaginationPhone struct {
	Pagination Pagination `json:"pagination"`
	Phones     []Phone    `json:"phones"`
}

func (pp *PaginationPhone) Prepare(phones []Phone, pagination Pagination, page int) {
	pp.Pagination = pagination

	if pagination.Total > 0 && pagination.Total > pagination.PerPage {
		if page == 1 {
			pp.Phones = phones[0:pagination.PerPage]
		} else {
			if page >= pagination.Total/pagination.PerPage {
				min := page * pagination.PerPage
				pp.Phones = phones[min:]
			} else {
				max := page * pagination.PerPage
				min := max - pagination.PerPage
				pp.Phones = phones[min:max]
			}
		}
	} else {
		pp.Phones = phones
	}
}
