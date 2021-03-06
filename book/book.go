package book

import (
	"sort"

	"github.com/fxtlabs/date"
)

// Book is a data structure that represents a book
type Book struct {
	Title   string
	Author  string
	PubDate date.Date
}

// ByPubDate implements sort.Interface for []Book based
// on the PubDate field in ascending (oldest first) order
type ByPubDate []Book

func (p ByPubDate) Len() int           { return len(p) }
func (p ByPubDate) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByPubDate) Less(i, j int) bool { return p[i].PubDate.Before(p[j].PubDate) }
func (p ByPubDate) SortAscending()     { sort.Sort(p) }
func (p ByPubDate) SortDesceding() {
	sort.Slice(p, func(i, j int) bool { return p[i].PubDate.After(p[j].PubDate) })
}
