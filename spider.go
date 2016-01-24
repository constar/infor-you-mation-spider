package main

import (
	"flag"
	"sync"
	"time"

	"github.com/constar/infor-you-mation-spider/Godeps/_workspace/src/github.com/golang/glog"
	"github.com/constar/infor-you-mation-spider/Godeps/_workspace/src/github.com/yanyiwu/igo"
)

var sleepSeconds = flag.Int("sleep", 60, "sleep seconds")
var isForever = flag.Bool("forever", false, "run forever")

var wg sync.WaitGroup

func Crawl(url string) {
	content := igo.HttpGet(url)
	content = convert(content)
	msgs := Parse(content)
	if msgs == nil {
		glog.Error("Parse failed")
	} else {
		for _, item := range msgs {
			glog.V(3).Info(item)
			j := Job{
				item.GetTitle(),
				item.GetContent(),
				item.GetUrl(),
				igo.GetMd5String(item.GetUrl()),
			}
			jobid, err := SaveJob(j)
			if err != nil {
				glog.V(2).Info(err)
			} else {
				Dispatch(item.GetTitle(), jobid)
			}
			//oid, err := Insert("feeds", item.GetTitle(), item.GetContent(), item.GetUrl())
			//if err == nil {
			//	glog.Info(item.GetTitle(), " ", item.GetUrl())
			//	Dispatch(item.GetTitle(), oid)
			//} else {
			//	glog.V(2).Info(err)
			//}
		}
	}
}

func spiderRunner(url string) {
	defer wg.Done()
	if *isForever {
		for {
			Crawl(url)
			glog.V(3).Info("time.Sleep ", *sleepSeconds, " seconds")
			time.Sleep(time.Duration(*sleepSeconds) * time.Second)
		}
	} else {
		Crawl(url)
	}
}

func main() {
	flag.Parse()
	for _, topic := range TOPICS {
		SaveTopic(topic)
	}
	TopicDispatcherInit()
	for i := 0; i < len(RssUrls); i++ {
		url := RssUrls[i]
		wg.Add(1)
		go spiderRunner(url)
		glog.Info(url)
	}
	wg.Wait()
}
