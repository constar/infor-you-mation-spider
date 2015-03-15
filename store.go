package main

import (
	"github.com/yanyiwu/igo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ByrDataItem struct {
	Id           bson.ObjectId "_id"
	Title        string
	Content      string
	Url          string
	UrlMd5       string
	LastModified time.Time
}

type MongoClient struct {
	host   string
	dbSess *mgo.Session
	dbName string
}

var (
	connOption MongoClient
)

func init() {
	Connect(MongoDBHost, DBName)
}

func Connect(host string, dbName string) error {
	var err error
	connOption.host = host
	connOption.dbSess, err = mgo.Dial(host)
	connOption.dbName = dbName
	return err
}

func Close() {
	connOption.dbSess.Close()
}

func Insert(
	colName string,
	title string,
	content string,
	url string) (bson.ObjectId, error) {
	c := connOption.dbSess.DB(connOption.dbName).C(colName)

	oid := bson.NewObjectId()
	last_modified := time.Now()
	last_modified = time.Unix(last_modified.Unix()+8*3600, 0)
	bdi := ByrDataItem{
		oid,
		title,
		content,
		url,
		igo.GetMd5String(url),
		last_modified,
	}

	return oid, c.Insert(&bdi)
}
