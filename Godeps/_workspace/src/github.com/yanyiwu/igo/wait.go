package igo

import (
	"sync"
)

type WaitGroup struct {
	sync.WaitGroup
}

func (w *WaitGroup) Go(f func()) {
	w.Add(1)
	go func() {
		defer w.Done()
		f()
	}()
}
