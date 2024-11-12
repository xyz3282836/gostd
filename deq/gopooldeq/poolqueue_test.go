// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gopooldeq

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func pack(head, tail uint32) uint64 {
	const mask = 1<<dequeueBits - 1
	return (uint64(head) << dequeueBits) |
		uint64(tail&mask)
}

func unpack(ptrs uint64) (head, tail uint32) {
	const mask = 1<<dequeueBits - 1
	head = uint32((ptrs >> dequeueBits) & mask)
	tail = uint32(ptrs & mask)
	return
}

func Test_Pool(t *testing.T) {
	// headTail := pack(1<<dequeueBits-500, 1<<dequeueBits-500)
	var headTail uint64
	vallen := 1 << 7
	for i := 0; i < 400; i++ {
		head, _ := unpack(headTail)
		index := head & uint32(vallen-1)
		fmt.Printf("index:%d %b\n", index, index)
		atomic.AddUint64(&headTail, 1<<dequeueBits)
	}
}

func BenchmarkPoolDequeue(b *testing.B) {
	const size = 1024
	pd := NewPoolDequeue(size)
	var wg sync.WaitGroup

	// Producer
	go func() {
		for i := 0; i < b.N; i++ {
			pd.PushHead(i)
		}
		wg.Done()
	}()

	// Consumers
	numConsumers := 10
	wg.Add(numConsumers + 1)
	for i := 0; i < numConsumers; i++ {
		go func() {
			for {
				if _, ok := pd.PopTail(); !ok {
					break
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkPoolChain(b *testing.B) {
	pc := NewPoolChain()
	var wg sync.WaitGroup

	// Producer
	go func() {
		for i := 0; i < b.N; i++ {
			pc.PushHead(i)
		}
		wg.Done()
	}()

	// Consumers
	numConsumers := 10
	wg.Add(numConsumers + 1)
	for i := 0; i < numConsumers; i++ {
		go func() {
			for {
				if _, ok := pc.PopTail(); !ok {
					break
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkChannel(b *testing.B) {
	ch := make(chan interface{}, 1024)
	var wg sync.WaitGroup

	// Producer
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()

	// Consumers
	numConsumers := 10
	wg.Add(numConsumers + 1)
	for i := 0; i < numConsumers; i++ {
		go func() {
			for range ch {
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

// go test 默认会运行所有的单元测试和基准测试。
// 使用 -run=^$ 可以避免运行任何单元测试，只运行基准测试。
// go test -run=^$ -benchmem -bench=^Benchmark -count=5
// go test -v -run Test_Pool
