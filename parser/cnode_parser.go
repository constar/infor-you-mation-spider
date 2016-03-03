package parser

import (
	"encoding/xml"

	"github.com/constar/infor-you-mation-spider/Godeps/_workspace/src/github.com/golang/glog"
)

type CNodeRss struct {
	Channel CNodeRssChannel `xml:"channel"`
}

type CNodeRssChannel struct {
	Items []CNodeRssItem `xml:"item"`
}

type CNodeRssItem struct {
	Title   string `xml:"title"`
	Content string `xml:"description"`
	Link    string `xml:"link"`
}

type CNodeParser struct {
}

func NewCNodeParser() *CNodeParser {
	return &CNodeParser{}
}

func (this *CNodeParser) Parse(content []byte) []Message {
	var rss CNodeRss
	err := xml.Unmarshal(content, &rss)
	if err != nil {
		glog.Error(err)
		return nil
	}
	items := rss.Channel.Items
	msgs := make([]Message, 0, len(items))
	for i := 0; i < len(items); i++ {
		msg := Message{
			Title:   items[i].Title,
			Content: items[i].Content,
			Url:     items[i].Link,
		}
		msgs = append(msgs, msg)
	}
	return msgs
}
