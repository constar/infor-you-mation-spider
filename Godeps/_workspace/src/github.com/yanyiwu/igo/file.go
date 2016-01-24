package igo

import (
	"bufio"
	"github.com/constar/infor-you-mation-spider/Godeps/_workspace/src/github.com/golang/glog"
	"io"
	"os"
	"strings"
)

type LineIterator struct {
	fname  string
	f      *os.File
	reader *bufio.Reader
	line   string
}

func NewLineIterator(fname string) (*LineIterator, error) {
	li := new(LineIterator)
	li.fname = fname
	var err error
	li.f, err = os.Open(fname)
	if err != nil {
		return nil, err
	}
	li.reader = bufio.NewReader(li.f)
	return li, nil
}

func (li *LineIterator) Close() {
	li.f.Close()
}

func (li *LineIterator) HasNext() bool {
	line, err := li.reader.ReadString('\n')
	if err != nil {
		if io.EOF != err {
			glog.Error(err)
		}
		return false
	}
	li.line = strings.TrimRight(line, "\n")
	return true
}

func (li *LineIterator) Next() string {
	return li.line
}
