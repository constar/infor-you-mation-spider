package main

import (
	"flag"
	"sync"
	"time"

	"github.com/constar/infor-you-mation-spider/Godeps/_workspace/src/github.com/golang/glog"
	"github.com/constar/infor-you-mation-spider/Godeps/_workspace/src/github.com/yanyiwu/igo"
	"github.com/constar/infor-you-mation-spider/filter"
	"github.com/constar/infor-you-mation-spider/parser"
)

type Task struct {
	Name   string
	Url    string
	Parser parser.Parser
	Filter *filter.JobFilter
}

var TASKS = [...]Task{
	{"byr", "http://bbs.byr.cn/rss/board-ParttimeJob", parser.NewByrParser(), nil},
	{"byr", "http://bbs.byr.cn/rss/board-JobInfo", parser.NewByrParser(), nil},
	{"byr", "http://bbs.byr.cn/rss/board-Job", parser.NewByrParser(), filter.NewJobFilter()},
	{"smth", "http://www.newsmth.net/nForum/rss/board-Career_Campus", parser.NewSmthParser(), nil},
	{"smth", "http://www.newsmth.net/nForum/rss/board-Career_PHD", parser.NewSmthParser(), nil},
	{"smth", "http://www.newsmth.net/nForum/rss/board-Career_Plaza", parser.NewSmthParser(), nil},
	{"smth", "http://www.newsmth.net/nForum/rss/board-Career_Upgrade", parser.NewSmthParser(), nil},
	{"smth", "http://www.newsmth.net/nForum/rss/board-ExecutiveSearch", parser.NewSmthParser(), nil},
	{"v2ex", "http://www.v2ex.com/feed/jobs.xml", parser.NewV2exParser(), nil},
	{"cnode", "https://cnodejs.org/rss", parser.NewCNodeParser(), filter.NewJobFilter()},
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
	{"BAT", []string{"BAT"}},
	{"百度", []string{"百度"}},
	{"腾讯", []string{"腾讯"}},
	{"阿里", []string{"阿里"}},
	{"美团", []string{"美团"}},
	{"滴滴", []string{"滴滴", "嘀嘀打车"}},
	{"小米", []string{"小米"}},
	{"校招", []string{"校招"}},
	{"社招", []string{"社招"}},
	{"内推", []string{"内推"}},
	{"北京", []string{"北京", "帝都"}},
	{"上海", []string{"上海"}},
	{"苏杭", []string{"杭州", "苏州"}},
	{"广深", []string{"广州", "深圳"}},
}

var sleepSeconds = flag.Int("sleep", 60, "sleep seconds")
var isForever = flag.Bool("forever", false, "run forever")
var expireHours = flag.Int("expire", 24, "hours for redis expire")

var wg sync.WaitGroup

func Crawl(task Task) {
	content := igo.HttpGet(task.Url)
	msgs := task.Parser.Parse(content)
	if task.Filter != nil {
		msgs = task.Filter.Filter(msgs)
	}
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
				task.Name,
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

func spiderRunner(task Task) {
	defer wg.Done()
	if *isForever {
		for {
			Crawl(task)
			glog.V(3).Info("time.Sleep ", *sleepSeconds, " seconds")
			time.Sleep(time.Duration(*sleepSeconds) * time.Second)
		}
	} else {
		Crawl(task)
	}
}

func main() {
	flag.Parse()
	PurgeAllTopics()
	for _, topic := range TOPICS {
		SaveTopic(topic)
	}
	TopicDispatcherInit()
	for _, task := range TASKS {
		wg.Add(1)
		go spiderRunner(task)
	}
	wg.Wait()
}
