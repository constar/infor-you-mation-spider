package parser

import (
	"encoding/xml"
	"strings"

	"github.com/constar/infor-you-mation-spider/Godeps/_workspace/src/github.com/golang/glog"
)

type V2exRss struct {
	Items []V2exRssItem `xml:"entry"`
}

type V2exRssItem struct {
	Title   string      `xml:"title"`
	Content string      `xml:"content"`
	Link    V2exRssLink `xml:"link"`
}

type V2exRssLink struct {
	Url string `xml:"href,attr"`
}

type V2exParser struct {
}

func NewV2exParser() *V2exParser {
	return &V2exParser{}
}

func (this *V2exParser) Parse(content []byte) []Message {
	var rss V2exRss
	err := xml.Unmarshal(content, &rss)
	if err != nil {
		glog.Error(err)
		return nil
	}
	size := len(rss.Items)
	msgs := make([]Message, 0, size)
	for i := 0; i < size; i++ {
		segs := strings.Split(rss.Items[i].Link.Url, "#")
		if len(segs) == 0 {
			glog.Error(rss.Items[i].Link.Url, "illegal")
			continue
		}
		url := segs[0]
		msg := Message{
			Title:   rss.Items[i].Title,
			Content: rss.Items[i].Content,
			Url:     url,
		}
		msgs = append(msgs, msg)
	}
	return msgs
}
