package misccode

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	ctx, cancel := context.WithCancel(context.Background())

	go handle(ctx, 5*time.Second)
	// 主goroutine会卡在这里不能继续往下执行cancel
	// 如果是计时cancel会等待时间结束,这里会继续执行
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
	cancel()
}

func handle(ctx context.Context, duration time.Duration) {

	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with:", duration)
	}
}
