package pager

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExamplePager() {
	data := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	p := Pager{
		Items:    data,
		PageSize: 3,
	}
	fmt.Println(p.RawPage(2))
	// Output: [d e f]
}

var data = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

func TestPagerGetStartEnd(t *testing.T) {
	cases := []struct {
		ps int  // pagesize
		pn int  // pagenum
		es int  // expect start
		ee int  // expect end
		ep bool // expect panic
	}{
		{0, 2, 0, 0, true},
		{2, 0, 0, 0, true},
		{2, 1, 0, 2, false},
		{2, 2, 2, 4, false},
		{2, 9, 10, 10, false},
	}

	for _, c := range cases {
		runner := func() {
			p := Pager{Items: data, PageSize: c.ps}
			start, end := p.getStartEnd(c.pn)

			assert.Equal(t, c.es, start)
			assert.Equal(t, c.ee, end)
		}

		if c.ep {
			assert.Panics(t, runner)
		} else {
			assert.NotPanics(t, runner)
		}
	}
}

func TestPagerPage(t *testing.T) {
	cases := []struct {
		ps  int           // pagesize
		pn  int           // pagenum
		epc int           // expected PageCount
		erp []interface{} // expected RawPage
		ep  Page          // expected Page
	}{
		{
			2, 1,
			5,
			[]interface{}{"a", "b"},
			&page{
				pageNumber: 1,
				items: []PageItem{
					&pageItem{
						inPageIndex: 0,
						globalIndex: 0,
						data:        "a",
					},
					&pageItem{
						inPageIndex: 1,
						globalIndex: 1,
						data:        "b",
					},
				},
			},
		},
		{
			2, 2,
			5,
			[]interface{}{"c", "d"},
			&page{
				pageNumber: 2,
				items: []PageItem{
					&pageItem{
						inPageIndex: 0,
						globalIndex: 2,
						data:        "c",
					},
					&pageItem{
						inPageIndex: 1,
						globalIndex: 3,
						data:        "d",
					},
				},
			},
		},
		{
			3, 2,
			4,
			[]interface{}{"d", "e", "f"},
			&page{
				pageNumber: 2,
				items: []PageItem{
					&pageItem{
						inPageIndex: 0,
						globalIndex: 3,
						data:        "d",
					},
					&pageItem{
						inPageIndex: 1,
						globalIndex: 4,
						data:        "e",
					},
					&pageItem{
						inPageIndex: 2,
						globalIndex: 5,
						data:        "f",
					},
				},
			},
		},
	}

	for _, c := range cases {
		pager := Pager{data, c.ps}

		assert.Equal(t, c.erp, pager.RawPage(c.pn))
		assert.Equal(t, c.epc, pager.PageCount())

		p := pager.Page(c.pn)
		assert.Equal(t, c.pn, p.PageNumber())
		assert.Equal(t, c.ep.(*page).items, p.Items())
		assert.Equal(t, len(c.ep.(*page).items), p.Size())

		for _, pi := range c.ep.(*page).items {
			assert.Equal(
				t,
				pi.(*pageItem).inPageIndex,
				pi.InPageIndex(),
			)
			assert.Equal(
				t,
				pi.(*pageItem).globalIndex,
				pi.GlobalIndex(),
			)
			assert.Equal(
				t,
				pi.(*pageItem).data,
				pi.Data(),
			)
		}

		assert.Equal(t, c.ep, p)
	}
}
