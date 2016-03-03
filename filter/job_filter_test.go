package filter

import (
	"testing"

	"github.com/constar/infor-you-mation-spider/parser"
)

func TestJobFilter(t *testing.T) {
	x := NewJobFilter()
	msgs := []parser.Message{
		{"title", "content", "url"},
		{"招聘工程师", "content", "url"},
	}
	msgs = x.Filter(msgs)
	if len(msgs) != 1 {
		t.Fatal()
	}
	if msgs[0].Title != "招聘工程师" {
		t.Fatal(msgs[0].Title)
	}
}
