package main

import (
	"flag"
	"sync"
	"time"

	"github.com/constar/infor-you-mation-spider/Godeps/_workspace/src/github.com/golang/glog"
	"github.com/constar/infor-you-mation-spider/Godeps/_workspace/src/github.com/yanyiwu/igo"
	"github.com/constar/infor-you-mation-spider/parser"
)

var BYR_RSS_URLS = [...]string{
	"http://bbs.byr.cn/rss/board-ParttimeJob",
	"http://bbs.byr.cn/rss/board-JobInfo",
}

var SMTH_RSS_URLS = [...]string{
	"http://www.newsmth.net/nForum/rss/board-Career_Campus",
	"http://www.newsmth.net/nForum/rss/board-Career_PHD",
	"http://www.newsmth.net/nForum/rss/board-Career_Plaza",
	"http://www.newsmth.net/nForum/rss/board-Career_Upgrade",
	"http://www.newsmth.net/nForum/rss/board-ExecutiveSearch",
}

var TOPICS = []Topic{
	{"Android", []string{"Android", "安卓"}},
	{"大数据", []string{"大数据", "数据挖掘", "数据分析"}},
	{"人工智能", []string{"人工智能", "机器学习", "自然语言处理"}},
	{"设计", []string{"设计", "交互", "UI", "UE", "美工"}},
	{"云计算", []string{"云计算", "分布式"}},
	{"实习/兼职", []string{"实习", "兼职"}},
	{"Web前端", []string{"前端", "h5", "html", "js", "javascript"}},
	{"创业", []string{"创业"}},
	{"产品", []string{"产品"}},
	{"PHP", []string{"PHP"}},
	{"iOS", []string{"iOS"}},
	{"C++", []string{"C++", "cpp"}},
	{"Java", []string{"Java"}},
	{"Python", []string{"Python"}},
	{"运营/市场", []string{"运营", "市场"}},
	{"Golang", []string{"go"}},
}

var sleepSeconds = flag.Int("sleep", 60, "sleep seconds")
var isForever = flag.Bool("forever", false, "run forever")

var wg sync.WaitGroup

func Crawl(url string, parser parser.Parser) {
	content := igo.HttpGet(url)
	content = convert(content)
	msgs := parser.Parse(content)
	if msgs == nil {
		glog.Error("Parse failed")
	} else {
		for _, item := range msgs {
			glog.V(3).Info(item)
			j := Job{
				item.Title,
				item.Content,
				item.Url,
				igo.GetMd5String(item.Url),
			}
			jobid, err := SaveJob(j)
			if err != nil {
				glog.V(2).Info(err)
			} else {
				Dispatch(item.Title, jobid)
			}
		}
	}
}

func spiderRunner(url string, parser parser.Parser) {
	defer wg.Done()
	if *isForever {
		for {
			Crawl(url, parser)
			glog.V(3).Info("time.Sleep ", *sleepSeconds, " seconds")
			time.Sleep(time.Duration(*sleepSeconds) * time.Second)
		}
	} else {
		Crawl(url, parser)
	}
}

func main() {
	flag.Parse()
	byrparser := parser.NewByrParser()
	smthparser := parser.NewSmthParser()
	for _, topic := range TOPICS {
		SaveTopic(topic)
	}
	TopicDispatcherInit()
	for _, url := range BYR_RSS_URLS {
		wg.Add(1)
		go spiderRunner(url, byrparser)
		glog.Info(url)
	}
	for _, url := range SMTH_RSS_URLS {
		wg.Add(1)
		go spiderRunner(url, smthparser)
		glog.Info(url)
	}
	wg.Wait()
}
