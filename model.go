package main

import (
	"strconv"

	"github.com/golang/glog"

	"gopkg.in/redis.v3"
)

const (
	EXPIRE_TIME = 0
)

var (
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
)

type Job struct {
	title   string
	content string
	url     string
	urlmd5  string
}

func SaveJob(j Job) error {
	id := getJobId(j.urlmd5)
	//client.Set("job:"+id+":title", j.title, EXPIRE_TIME)
	glog.Info(id)
	//client.Set("job" + "")
	return nil
}

func getJobId(urlmd5 string) string {
	val, err := client.Get("job:" + urlmd5 + ":id").Result()
	if err == redis.Nil {
		val, err := client.Incr("job:nextid").Result()
		if err != nil {
			panic(err)
		}
		id := strconv.FormatInt(val, 10)
		glog.V(3).Info(urlmd5, id)
		client.Set("job:"+urlmd5+":id", id, EXPIRE_TIME)
		return id
	}
	if err != nil {
		panic(err)
	}
	return val
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
