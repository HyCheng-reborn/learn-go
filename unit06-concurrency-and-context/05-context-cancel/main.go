package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("worker: cancelled,", ctx.Err())

	case <-time.After(time.Second):
		fmt.Println("worker: finished")
	}
}

func main() {
	// ============================================================
	// 实验 a：手动取消（WithCancel）
	// 预期：worker: cancelled, context canceled
	// 想跑实验 b 时，把下面这一整块注释掉即可。
	// ============================================================
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	worker(ctx)
	// }()

	// // 等待 200ms 后主动 cancel()，早于 worker 里的 time.After(1s)
	// time.Sleep(200 * time.Millisecond)
	// cancel()

	// // 等 worker 响应 cancel 并打印完再退出
	// wg.Wait()

	// ============================================================
	// 实验 b：超时取消（WithTimeout）
	// 预期：worker: cancelled, context deadline exceeded
	// 想跑实验 b 时，取消下面这一整块的注释，并把实验 a 那块注释掉。
	// ============================================================

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		worker(ctx)
	}()

	// 不手动 cancel，等 300ms 超时自动触发
	wg.Wait()

}
