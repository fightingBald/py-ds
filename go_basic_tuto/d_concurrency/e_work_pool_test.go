package d_concurrency

import (
	"fmt"
	"sync"
	"testing"
)

/*Go 并发里经典的 worker pool（工作池）

1. Worker Pool（工人池 / 线程池）
是什么
固定数量的 worker goroutine 从任务队列里取活干。
“N 个人在流水线前排着，来了任务就一个人伸手去拿”。
常用场景
控制并发度（不能每来一个任务就起一个 goroutine）。
后端服务里消费消息队列（Kafka/RabbitMQ）。
CPU 密集型/外部调用，防止撑爆下游。

核心思想
背压 + 限流：并发数量受控。
复用 goroutine：避免无限 goroutine 堆积。

*/

/*
WaitGroup 在这里的任务是：等 10 个工人 goroutine 全部退出。

为什么要等？因为只有确定没人再往 results 里写了，你才安全地 close(results)。

不用 WaitGroup？行啊，你就得想别的骚操作（比如给每个工人单独一个 done channel，最后再 fan-in 聚合）。

但写起来比 WaitGroup 臭长一截。
*/
func TestWorkerPool(t *testing.T) {

	type job struct{ id int }
	jobs := make(chan job, 100)  // 任务队列
	results := make(chan string) // 结果汇总

	var wg sync.WaitGroup
	workers := 10
	wg.Add(workers) //* 1.你起了 10 个工人：

	//* 2定义worker task 匿名函数
	worker := func(id int) {
		defer wg.Done() //我这个函数里干啥你别管，但只要我退出（正常/异常都行），你一定要帮我执行 wg.Done()。”
		for j := range jobs {
			// 这里模拟处理任务：下载/计算/DB写入...
			results <- fmt.Sprintf("worker-%d done job-%d", id, j.id)
		}

	}
	for i := 1; i <= workers; i++ {
		go worker(i)
	}

	// 派发 25 个任务
	go func() {
		for i := 1; i <= 25; i++ {
			jobs <- job{id: i}
		}
		close(jobs) // 任务发完，关掉
	}()

	// 等工人退出后，关结果通道
	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println(r) // 消耗channel里的结果
	}
}
