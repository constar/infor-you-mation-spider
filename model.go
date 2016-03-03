package main

import (
	"reflect"
	"strconv"
	"time"

	"github.com/constar/infor-you-mation-spider/Godeps/_workspace/src/github.com/golang/glog"

	"github.com/constar/infor-you-mation-spider/Godeps/_workspace/src/gopkg.in/redis.v3"
)

const (
	NOT_EXPIRE = 0
)

var (
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
)

type Job struct {
	Title   string `redis:"title"`
	Content string `redis:"content"`
	Url     string `redis:"url"`
	Urlmd5  string `redis:"urlmd5"`
	Source  string `redis:"source"`
}

func SaveJob(j Job) (string, error) {
	id, err := getJobId(j.Urlmd5)
	if err != nil {
		return "", err
	}

	t := reflect.TypeOf(j)
	v := reflect.ValueOf(&j).Elem()
	for i := 0; i < v.NumField(); i++ {
		key := "job:" + id + ":" + t.Field(i).Tag.Get("redis")
		value := v.Field(i).Interface().(string)
		_, err := client.Set(key, value, (time.Duration)(*expireHours)*time.Hour).Result()
		if err != nil {
			return "", err
		}
	}
	return id, nil
}

func SaveTopicJobList(topic string, jobid string) error {
	id, err := getTopicId(topic)
	if err != nil {
		glog.Error(err)
		return err
	}

	now := time.Now().Unix()
	key := "topic:" + id + ":joblist"
	value := redis.Z{
		Score:  float64(now),
		Member: jobid,
	}
	_, err = client.ZAdd(key, value).Result()
	return err
}

func PurgeAllTopics() error {
	maxid, err := client.Get("topic:nextid").Result()
	if err != nil {
		glog.Error(err)
		return err
	}
	max, err := strconv.Atoi(maxid)
	if err != nil {
		glog.Error(err)
		return err
	}
	for i := 1; i <= max; i++ {
		if err := PurgeTopicJobList(strconv.Itoa(i)); err != nil {
			glog.Error(err)
			return err
		}
	}
	return nil
}

func PurgeTopicJobList(topicid string) error {
	key := "topic:" + topicid + ":joblist"
	old := time.Now().Add(-1 * (time.Duration)(*expireHours) * time.Hour).Unix()
	min := "0"
	max := strconv.Itoa(int(old))
	_, err := client.ZRemRangeByScore(key, min, max).Result()
	return err
}

func SaveTopic(topic Topic) error {
	id, err := getTopicId(topic.Name)
	if err != nil {
		return err
	}
	var key string
	key = "topic:" + id + ":name"
	_, err = client.Set(key, topic.Name, NOT_EXPIRE).Result()
	if err != nil {
		return err
	}
	key = "topic:" + id + ":words"
	_, err = client.SAdd(key, topic.Words...).Result()
	return err
}

func getTopicId(topic string) (string, error) {
	return getAutoId("topic", topic)
}

func getJobId(urlmd5 string) (string, error) {
	return getAutoId("job", urlmd5)
}

func getAutoId(table string, uniq string) (string, error) {
	val, err := client.Get(table + ":" + uniq + ":id").Result()
	if err == redis.Nil {
		val, err := client.Incr(table + ":nextid").Result()
		if err != nil {
			return "", err
		}
		id := strconv.FormatInt(val, 10)
		glog.V(3).Info(uniq, id)
		client.Set(table+":"+uniq+":id", id, NOT_EXPIRE)
		return id, nil
	}
	if err != nil {
		return "", err
	}
	return val, nil
}
