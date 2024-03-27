package domain

type PaginationOptions struct {
	Page    int64
	PerPage int64
}

func NewPaginationOptions(page, perPage int64) PaginationOptions {
	return PaginationOptions{
		Page:    page,
		PerPage: perPage,
	}
}

func (opt *PaginationOptions) ApplyOptions(options PaginationOptions) PaginationOptions {
	if options.Page != 0 {
		opt.Page = options.Page
	}
	if options.PerPage != 0 {
		opt.PerPage = options.PerPage
	}
	return *opt
}

func (opt PaginationOptions) CalculateSkip() int64 {
	return (opt.PerPage * opt.Page) - opt.PerPage
}
