package mutex

import "sync"

var msg string

func updateMsg(s string, wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()

	mtx.Lock()
	msg = s
	mtx.Unlock()
}

func Main(flag bool) {
	if !flag {
		return
	}

	wg := &sync.WaitGroup{}
	mtx := &sync.Mutex{}

	wg.Add(2)
	go updateMsg("A", wg, mtx)
	go updateMsg("B", wg, mtx)
	wg.Wait()

	println(msg)
}
