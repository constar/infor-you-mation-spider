package igo

import (
	"github.com/constar/infor-you-mation-spider/Godeps/_workspace/src/github.com/golang/glog"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		glog.Error(err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Error(err)
		return nil
	}
	return body
}
