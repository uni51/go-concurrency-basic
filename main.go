package main

import (
	"fmt"
	"runtime"
	"sync"
)

func memConsumed() uint64 {
	runtime.GC() // GCを強制的に実行する

	var s runtime.MemStats // メモリの使用量を調べるための受け皿

	runtime.ReadMemStats(&s)

	return s.Sys
}

func main() {
	var ch <-chan interface{}
	var wg sync.WaitGroup

	noop := func() {
		wg.Done()
		<-ch // チャネルから値を受信する
	}

	const numGoroutines = 100000

	wg.Add(numGoroutines)

	before := memConsumed()

	for i := 0; i < numGoroutines; i++ {
		go noop()
	}

	wg.Wait()

	after := memConsumed()

	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}
