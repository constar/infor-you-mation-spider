package main

import (
	"reflect"
	"strconv"
	"time"

	"github.com/golang/glog"

	"gopkg.in/redis.v3"
)

const (
	// 3 days
	EXPIRE_TIME = 3 * 24 * 60 * 60
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
		_, err := client.Set(key, value, EXPIRE_TIME).Result()
		if err != nil {
			return "", err
		}
	}
	return id, nil
}

func SaveTopicJobList(topic string, jobid string) error {
	id, err := getTopicId(topic)
	if err != nil {
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

func SaveTopic(topic Topic) error {
	id, err := getTopicId(topic.Name)
	if err != nil {
		return err
	}
	var key string
	key = "topic:" + id + ":name"
	_, err = client.Set(key, topic.Name, EXPIRE_TIME).Result()
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
		client.Set(table+":"+uniq+":id", id, EXPIRE_TIME)
		return id, nil
	}
	if err != nil {
		return "", err
	}
	return val, nil
}

//func ExampleClient() {
//pong, err := client.Ping().Result()
//fmt.Println(pong, err)
// Output: PONG <nil>
//err = client.Set("key", "value", 0).Err()
//if err != nil {
//	panic(err)
//}

//val, err := client.Get("key").Result()
//if err != nil {
//	panic(err)
//}
//fmt.Println("key", val)

//val2, err := client.Get("key2").Result()
//if err == redis.Nil {
//	fmt.Println("key2 does not exists")
//} else if err != nil {
//	panic(err)
//} else {
//	fmt.Println("key2", val2)
//}
// Output: key value
// key2 does not exists
//}
