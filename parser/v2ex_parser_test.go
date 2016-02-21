package parser

import "testing"

func TestV2exParser(t *testing.T) {
	DATA := `
<?xml version="1.0" encoding="utf-8"?>
<feed xmlns="http://www.w3.org/2005/Atom">
<title>V2EX - 酷工作</title>
<subtitle>way to explore</subtitle>
<link rel="alternate" type="text/html" href="https://www.v2ex.com/" />
<link rel="self" type="application/atom+xml" href="https://www.v2ex.com/feed/jobs.xml" />
<id>https://www.v2ex.com/</id>

<updated>2016-02-21T03:39:46Z</updated>

<rights>Copyright © 2010-2012, V2EX</rights>
<entry>
	<title>title1</title>
	<link rel="alternate" type="text/html" href="https://www.v2ex.com/t/257977#reply5" />
	<id>tag:www.v2ex.com,2016-02-21:/t/257977</id>
	<published>2016-02-21T03:09:36Z</published>
	<updated>2016-02-21T03:39:46Z</updated>
	<author>
		<name>Montang</name>
		<uri>https://www.v2ex.com/member/Montang</uri>
	</author>
	<content type="html" xml:base="https://www.v2ex.com/" xml:lang="en"><![CDATA[
	<p>hello</p>
	]]></content>
</entry><entry>
	<title>title2</title>
	<link rel="alternate" type="text/html" href="https://www.v2ex.com/t/257975#reply2" />
	<id>tag:www.v2ex.com,2016-02-21:/t/257975</id>
	<published>2016-02-21T03:07:36Z</published>
	<updated>2016-02-21T03:36:47Z</updated>
	<author>
		<name>zyy93</name>
		<uri>https://www.v2ex.com/member/zyy93</uri>
	</author>
	<content type="html" xml:base="https://www.v2ex.com/" xml:lang="en"><![CDATA[
<p>各位大大，请问我近期做什么才能尽快找到工作呢？</p>
	]]></content>
</entry>
</feed>
	`
	parser := NewV2exParser()
	msgs := parser.Parse([]byte(DATA))
	if len(msgs) != 2 {
		t.Fatal(len(msgs))
	}

	if msgs[0].Url != "https://www.v2ex.com/t/257977#reply5" {
		t.Fatal(msgs[0].Url)
	}
	if msgs[0].Title != "title1" {
		t.Fatal(msgs[0].Title)
	}
}
