package ratelimit

import (
	"fmt"
	"time"
)

func rateLimiting(requestQueue chan int, duration int64, allowedBurstCount int) {
	//根据允许的突发数量,创建 ch
	//只要该队列中有内容，就可以一直取出来，即便 ch 已经关闭
	burstyLimiter := make(chan time.Time, allowedBurstCount)

	//初始化允许突发的数量
	for i := 0; i < allowedBurstCount; i++ {
		burstyLimiter <- time.Now()
	}

	//控制请求频率的计时器
	go func() {
		for t := range time.Tick(time.Duration(duration) * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	for req := range requestQueue {
		<-burstyLimiter //突发控制器是限流的关键
		fmt.Println("request", req, time.Now())
	}
}

// Another rate limiting example in simple way.
func simpleRateLimiting(duration int64, requestQueueSize int) {
	requests := make(chan int, requestQueueSize)
	for i := 1; i <= requestQueueSize; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Duration(duration) * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
}
