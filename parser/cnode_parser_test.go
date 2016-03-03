package parser

import "testing"

func TestCNodeParser(t *testing.T) {
	DATA := `
<?xml version="1.0" encoding="utf-8"?>
<rss version="2.0"><channel><title>CNode：Node.js专业中文社区</title><link>http://cnodejs.org</link><language>zh-cn</language><description>CNode：Node.js专业中文社区</description>
<item><title>新人求职前端</title><link>http://cnodejs.org/topic/56d80c5f820d3c9b3d63e2c1</link><guid>http://cnodejs.org/topic/56d80c5f820d3c9b3d63e2c1</guid><description>&lt;div class=&quot;markdown-text&quot;&gt;&lt;p&gt;最近做了个练手和打算用来求职的项目,演示地址</description><author>suyuanhan</author><pubDate>Thu, 03 Mar 2016 10:05:19 GMT</pubDate></item>
<item><title>深入浅出容器云</title><link>http://cnodejs.org/topic/56d808577947807f3daf8675</link><guid>http://cnodejs.org/topic/56d808577947807f3daf8675</guid><description>d2</description><author>limeflavor</author><pubDate>Thu, 03 Mar 2016 09:48:07 GMT</pubDate></item>
</channel>
</rss>
	`

	parser := NewCNodeParser()
	msgs := parser.Parse([]byte(DATA))
	if len(msgs) != 2 {
		t.Fatal(len(msgs))
	}

	if msgs[0].Url != "http://cnodejs.org/topic/56d80c5f820d3c9b3d63e2c1" {
		t.Fatal(msgs[0].Url)
	}
	if msgs[0].Title != "新人求职前端" {
		t.Fatal(msgs[0].Title)
	}
}
