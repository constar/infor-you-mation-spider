package parser

import (
	"encoding/xml"

	"github.com/constar/infor-you-mation-spider/Godeps/_workspace/src/github.com/golang/glog"
)

type ByrRss struct {
	ByrRssChannel ByrRssChannel `xml:"channel"`
}

type ByrRssChannel struct {
	Title         string       `xml:"title"`
	Description   string       `xml:"description"`
	Link          string       `xml:"link"`
	Language      string       `xml:"language"`
	Generator     string       `xml:"generator"`
	LastBuildDate string       `xml:"lastBuildDate"`
	Items         []ByrRssItem `xml:"item"`
}

type ByrRssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Author      string `xml:"author"`
	PubDate     string `xml:"pubDate"`
	Guid        string `xml:"guid"`
	Comments    string `xml:"comments"`
	Description string `xml:"description"`
}

type ByrParser struct {
}

func NewByrParser() *ByrParser {
	return &ByrParser{}
}

func (this *ByrParser) Parse(content []byte) []Message {
	content = convert(content)
	var rss ByrRss
	err := xml.Unmarshal(content, &rss)
	if err != nil {
		glog.Error(err)
		return nil
	}
	size := len(rss.ByrRssChannel.Items)
	msgs := make([]Message, 0, size)
	for i := 0; i < size; i++ {
		msg := Message{
			Title:   rss.ByrRssChannel.Items[i].Title,
			Content: rss.ByrRssChannel.Items[i].Description,
			Url:     rss.ByrRssChannel.Items[i].Link,
		}
		msgs = append(msgs, msg)
	}
	return msgs
}
