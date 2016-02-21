package parser

import (
	"strings"
	"sync"

	"github.com/constar/infor-you-mation-spider/Godeps/_workspace/src/github.com/golang/glog"
	"github.com/constar/infor-you-mation-spider/Godeps/_workspace/src/github.com/qiniu/iconv"
)

var ic iconv.Iconv
var icMutex sync.Mutex

func init() {
	var err error
	ic, err = iconv.Open("utf-8", "gbk")
	if err != nil {
		glog.Fatal(err)
	}
}

// concurrency-safety
func convert(data []byte) []byte {
	icMutex.Lock()
	defer icMutex.Unlock()
	res := ic.ConvString(string(data))
	res = strings.Replace(res, `encoding="gb2312"`, `encoding="utf-8"`, 1)
	return []byte(res)
}
