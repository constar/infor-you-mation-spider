package filter

import (
	"strings"

	"github.com/constar/infor-you-mation-spider/parser"
)

var WORDS = [...]string{
	"招聘",
	"诚聘",
	"内推",
	"10K",
	"15K",
	"20K",
	"30K",
	"40K",
}

type JobFilter struct {
}

func NewJobFilter() *JobFilter {
	return &JobFilter{}
}

func (this *JobFilter) Filter(msgs []parser.Message) []parser.Message {
	result := make([]parser.Message, 0)
	for _, m := range msgs {
		if this.isJob(m) {
			result = append(result, m)
		}
	}
	return result
}

func (this *JobFilter) isJob(msg parser.Message) bool {
	for _, word := range WORDS {
		if strings.Index(msg.Title, word) != -1 {
			return true
		}
	}
	return false
}
