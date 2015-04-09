package main

const (
	MongoDBHost = "127.0.0.1"
	DBName      = "inforyoumation"
	KeywordCol  = "keyword"
)

var (
	RssUrls = [...]string{
		"http://bbs.byr.cn/rss/board-ParttimeJob",
		"http://bbs.byr.cn/rss/board-JobInfo",
		"http://www.newsmth.net/nForum/rss/board-Career_Campus",
		"http://www.newsmth.net/nForum/rss/board-Career_PHD",
		"http://www.newsmth.net/nForum/rss/board-Career_Plaza",
		"http://www.newsmth.net/nForum/rss/board-Career_Upgrade",
		"http://www.newsmth.net/nForum/rss/board-ExecutiveSearch",
	}
)
