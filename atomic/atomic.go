package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var totalCount uint64

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 1; i <= 100; i++ {
		atomic.AddUint64(&totalCount, 1)
	}
}

func atomicCount() {
	var wg sync.WaitGroup
	l := 1000
	wg.Add(l)
	for i := 0;i < l; i ++ {
		go worker(&wg)
	}
	wg.Wait()
	fmt.Println(totalCount)
}



var total struct {
	sync.Mutex
	value int
}

func workerLock(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100000; i++ {
		total.Lock()
		total.value += 1
		total.Unlock()
	}
}

func lock() {
	var wg sync.WaitGroup

	l := 1000
	wg.Add(l)
	for i := 0;i < l; i ++ {
		go workerLock(&wg)
	}
	wg.Wait()
	fmt.Println(total.value)
}


func main() {
	start := time.Now().UnixNano()
	//lock() // 6299927000
	//atomicCount() // 2578000
	fmt.Println(time.Now().UnixNano() - start)
}