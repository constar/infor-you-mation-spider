package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/yanyiwu/igo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"sync"
	"time"
)

var dispatcher *TopicDispatcher

func TopicDispatcherInit() {
	dispatcher = NewTopicDispatcher()
	if dispatcher == nil {
		panic("NewTopicDispatcher failed")
	}
	for i := 0; i < len(gOrTopics); i++ {
		glog.Info("init: topic ", gOrTopics[i].Topic)
		for j := 0; j < len(gOrTopics[i].Words); j++ {
			glog.Info("init: word ", gOrTopics[i].Words[j])
			k := gOrTopics[i].Words[j]
			dispatcher.Insert(k, &gOrTopics[i])
		}
	}
}

func Dispatch(text string, feedid bson.ObjectId) {
	dispatcher.Dispatch(text, feedid)
}

type TopicDispatcher struct {
	trie   *igo.Trie
	lock   sync.RWMutex
	dbSess *mgo.Session
}

type KeywordColItem struct {
	Id           bson.ObjectId "_id"
	Keyword      string
	Feedid       bson.ObjectId
	LastModified time.Time
}

func (kci *KeywordColItem) String() string {
	return fmt.Sprintf("%v %s %v", kci.Id, kci.Keyword, kci.Feedid)
}

func NewTopicDispatcher() *TopicDispatcher {
	kw := new(TopicDispatcher)
	kw.trie = igo.NewTrie()
	var err error
	kw.dbSess, err = mgo.Dial(MongoDBHost)
	if err != nil {
		glog.Error(err)
		return nil
	}
	return kw
}

func (kw *TopicDispatcher) Insert(word string, ti TopicInterface) {
	word = strings.ToLower(word)
	kw.lock.Lock()
	defer kw.lock.Unlock()
	glog.Info("trie insert word: ", word)
	if err := kw.trie.Insert(word, ti); err != nil {
		glog.Error(err)
	}
}

func (kw *TopicDispatcher) Dispatch(text string, feedid bson.ObjectId) {
	kw.lock.RLock()
	defer kw.lock.RUnlock()
	text = strings.ToLower(text)
	res := kw.trie.Find(text)
	for i := 0; i < len(res); i++ {
		interf := res[i].Data.(TopicInterface)
		err := kw.dispatchOne(interf.GetTopic(), feedid)
		if err != nil {
			glog.Info(err)
		}
	}
}

func (kw *TopicDispatcher) dispatchOne(topic string, feedid bson.ObjectId) error {
	c := kw.dbSess.DB(DBName).C(KeywordCol)
	last_modified := time.Now()
	last_modified = time.Unix(last_modified.Unix()+8*3600, 0)
	kci := KeywordColItem{
		bson.NewObjectId(),
		topic,
		feedid,
		last_modified,
	}
	glog.Infof("insert %s to %s.%s", kci, DBName, KeywordCol)
	return c.Insert(&kci)
}
