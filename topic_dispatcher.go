package main

import (
	"strings"
	"sync"

	"github.com/constar/infor-you-mation-spider/Godeps/_workspace/src/github.com/golang/glog"
	"github.com/constar/infor-you-mation-spider/Godeps/_workspace/src/github.com/yanyiwu/igo"
)

var dispatcher *TopicDispatcher

func TopicDispatcherInit() {
	dispatcher = NewTopicDispatcher()
	if dispatcher == nil {
		panic("NewTopicDispatcher failed")
	}
	for i := 0; i < len(TOPICS); i++ {
		glog.Info("init: topic ", TOPICS[i].Name)
		for j := 0; j < len(TOPICS[i].Words); j++ {
			glog.Info("init: word ", TOPICS[i].Words[j])
			k := TOPICS[i].Words[j]
			dispatcher.Insert(k, TOPICS[i])
		}
	}
}

func Dispatch(text string, jobid string) {
	dispatcher.Dispatch(text, jobid)
}

type TopicDispatcher struct {
	trie *igo.Trie
	lock sync.RWMutex
}

func NewTopicDispatcher() *TopicDispatcher {
	kw := new(TopicDispatcher)
	kw.trie = igo.NewTrie()
	return kw
}

func (kw *TopicDispatcher) Insert(word string, topic Topic) {
	word = strings.ToLower(word)
	kw.lock.Lock()
	defer kw.lock.Unlock()
	glog.Info("trie insert word: ", word)
	if err := kw.trie.Insert(word, topic); err != nil {
		glog.Error(err)
	}
}

func (kw *TopicDispatcher) Dispatch(text string, jobid string) {
	kw.lock.RLock()
	defer kw.lock.RUnlock()
	text = strings.ToLower(text)
	res := kw.trie.Find(text)
	for i := 0; i < len(res); i++ {
		topic := res[i].Data.(Topic)
		err := SaveTopicJobList(topic.Name, jobid)
		if err != nil {
			glog.Info(err)
		}
	}
}
