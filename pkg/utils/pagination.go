package utils

import (
	"fmt"
	"math"
	"strconv"
)

const (
	_defaultSize = 10
	_defaultPage = 1
)

type Pagination struct {
	Size    int    `json:"size,omitempty"`
	Page    int    `json:"page,omitempty"`
	OrderBy string `json:"order_by,omitempty"`
}

func NewPaginationQuery(size, page int) *Pagination {
	if size == 0 {
		return &Pagination{
			Size: _defaultSize,
			Page: _defaultPage,
		}
	}
	return &Pagination{
		Size: size,
		Page: page,
	}
}

func NewPaginationFromQueryParams(size, page string) *Pagination {
	p := &Pagination{
		Size: _defaultSize,
		Page: 1,
	}

	if sizeNum, err := strconv.Atoi(size); err == nil && sizeNum != 0 {
		p.Size = sizeNum
	}

	if pageNum, err := strconv.Atoi(page); err == nil && pageNum != 0 {
		p.Page = pageNum
	}

	return p
}

func (p *Pagination) SetSize(sizeQuery string) error {
	if sizeQuery == "" {
		p.Size = _defaultPage
		return nil
	}
	n, err := strconv.Atoi(sizeQuery)
	if err != nil {
		return err
	}
	p.Size = n
	return nil
}

func (p *Pagination) SetPage(pageQuery string) error {
	if pageQuery == "" {
		p.Size = 0
		return nil
	}
	n, err := strconv.Atoi(pageQuery)
	if err != nil {
		return err
	}
	p.Page = n

	return nil
}

func (p *Pagination) SetOrderBy(orderByQuery string) {
	p.OrderBy = orderByQuery
}

func (p *Pagination) GetOffset() int {
	if p.Page == 0 {
		return 0
	}
	return (p.Page - 1) * p.Size
}

func (p *Pagination) GetLimit() int {
	return p.Size
}

func (p *Pagination) GetOrderBy() string {
	return p.OrderBy
}

func (p *Pagination) GetPage() int {
	return p.Page
}

func (p *Pagination) GetSize() int {
	return p.Size
}

func (p *Pagination) GetQueryString() string {
	return fmt.Sprintf("page=%v&size=%v&orderBy=%s", p.GetPage(), p.GetSize(), p.GetOrderBy())
}

func (p *Pagination) GetTotalPages(totalCount int) int {
	d := float64(totalCount) / float64(p.GetSize())
	return int(math.Ceil(d))
}

func (p *Pagination) GetHasMore(totalCount int) bool {
	return p.GetPage() < totalCount/p.GetSize()
}
