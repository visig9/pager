// Package pager are useful to retrieve a part of data by paging.
//
// For using, please manual create a Pager instance and call it's methods.
package pager

import "reflect"

// PageItem is a element in a page
type PageItem interface {
	// InPageIndex return a 0 based index in current page.
	InPageIndex() int
	// GlobalIndex return a 0 based index in whole data.
	GlobalIndex() int
	// Data return the correspondent raw element.
	Data() interface{}
}

type pageItem struct {
	inPageIndex int
	globalIndex int
	data        interface{}
}

func (i *pageItem) InPageIndex() int {
	return i.inPageIndex
}

func (i *pageItem) GlobalIndex() int {
	return i.globalIndex
}

func (i *pageItem) Data() interface{} {
	return i.data
}

// Page is single page of Pager
type Page interface {
	// PageNumber return the 1 based page number
	PageNumber() int
	// Size return the size, equal to len(Items())
	Size() int
	// Items return a slice contain PageItem in this page.
	Items() []PageItem
}

type page struct {
	pageNumber int
	items      []PageItem
}

func (p *page) PageNumber() int {
	return p.pageNumber
}

func (p *page) Size() int {
	return len(p.items)
}

func (p *page) Items() []PageItem {
	return p.items
}

// Pager offer a interface to retrieve a oart if data by paging.
type Pager struct {
	// Items must a slice or array want to paging.
	Items interface{}
	// PageSize is the size of page and must > 1
	PageSize int
}

func (p *Pager) getStartEnd(pageNumber int) (start, end int) {
	if p.PageSize < 1 {
		panic("PageSize < 1")
	}

	if pageNumber < 1 {
		panic("pageNumber < 1")
	}

	fullSize := reflect.ValueOf(p.Items).Len()

	start = (pageNumber - 1) * p.PageSize
	if start > fullSize {
		start = fullSize
	}

	end = pageNumber * p.PageSize
	if end > fullSize {
		end = fullSize
	}

	return
}

// PageCount return the total available page.
func (p *Pager) PageCount() int {
	size := reflect.ValueOf(p.Items).Len()
	count := size / p.PageSize
	if size%p.PageSize >= 1 {
		count++
	}

	return count
}

// Page export one page of the internal data with paging metadata.
//
// The pageNumber is 1 based. The given value can over PageCount safely.
func (p *Pager) Page(pageNumber int) Page {
	start, end := p.getStartEnd(pageNumber)
	v := reflect.ValueOf(p.Items)

	pItems := make([]PageItem, 0, end-start)
	for i := start; i < end; i++ {
		inPageIndex := i - start

		pItems = append(
			pItems,
			&pageItem{
				inPageIndex: inPageIndex,
				globalIndex: i,
				data:        v.Index(i).Interface(),
			},
		)
	}

	return &page{
		pageNumber: pageNumber,
		items:      pItems,
	}
}

// RawPage export one page of the internal data without metadata
//
// The pageNumber is 1 based. The given value can over PageCount safely.
func (p *Pager) RawPage(pageNumber int) []interface{} {
	start, end := p.getStartEnd(pageNumber)
	v := reflect.ValueOf(p.Items)

	items := make([]interface{}, 0, end-start)
	for i := start; i < end; i++ {
		items = append(items, v.Index(i).Interface())
	}

	return items
}
