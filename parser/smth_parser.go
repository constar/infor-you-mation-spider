package parser

import (
	"encoding/xml"

	"github.com/constar/infor-you-mation-spider/Godeps/_workspace/src/github.com/golang/glog"
)

type SmthRss struct {
	SmthRssChannel SmthRssChannel `xml:"channel"`
}

type SmthRssChannel struct {
	Title         string        `xml:"title"`
	Description   string        `xml:"description"`
	Link          string        `xml:"link"`
	Language      string        `xml:"language"`
	Generator     string        `xml:"generator"`
	LastBuildDate string        `xml:"lastBuildDate"`
	Items         []SmthRssItem `xml:"item"`
}

type SmthRssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Author      string `xml:"author"`
	PubDate     string `xml:"pubDate"`
	Guid        string `xml:"guid"`
	Comments    string `xml:"comments"`
	Description string `xml:"description"`
}

type SmthParser struct {
}

func NewSmthParser() *SmthParser {
	return &SmthParser{}
}

func (this *SmthParser) Parse(content []byte) []Message {
	content = convert(content)
	var rss SmthRss
	err := xml.Unmarshal(content, &rss)
	if err != nil {
		glog.Error(err)
		return nil
	}
	size := len(rss.SmthRssChannel.Items)
	msgs := make([]Message, 0, size)
	for i := 0; i < size; i++ {
		msg := Message{
			Title:   rss.SmthRssChannel.Items[i].Title,
			Content: rss.SmthRssChannel.Items[i].Description,
			Url:     rss.SmthRssChannel.Items[i].Link,
		}
		msgs = append(msgs, msg)
	}
	return msgs
}
