package main

import (
	"fmt"
	"github.com/constar/infor-you-mation-spider/Godeps/_workspace/src/github.com/qiniu/iconv"
	"os"
)

func main() {
	cd, err := iconv.Open("gbk", "utf-8") // utf8 => gbk
	if err != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	defer cd.Close()

	autoSync := false
	w := iconv.NewWriter(cd, os.Stdout, 0, autoSync)

	fmt.Fprintln(w,
		`		你好，世界！你好，世界！你好，世界！你好，世界！
		你好，世界！你好，世界！你好，世界！你好，世界！`)

	w.Sync() // call it by yourself if autoSync == false
}
