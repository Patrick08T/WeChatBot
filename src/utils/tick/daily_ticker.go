package tick

import (
	"fmt"
	"time"
)

func NineAMTasks(tasks ...func()) {
	DailyTicker(9, 0, 0, tasks...)
}

func DailyTicker(hour, minute, second int, call ...func()) {
	scheduledTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), hour, minute, second, 0, time.Local)
	durationUntilNextRun := scheduledTime.Sub(time.Now())
	if durationUntilNextRun < 0 {
		// 如果当前时间已经超过了设定的执行时间，就将执行时间推迟到明天的同一时刻
		durationUntilNextRun += 24 * time.Hour
	}
	// 创建一个定时器
	timer := time.NewTimer(durationUntilNextRun)

	// 使用匿名函数和goroutine执行任务
	go func() {
		for {
			<-timer.C // 当定时器触发时
			fmt.Println("执行每日定时任务...")

			// 在这里执行您的任务逻辑
			for _, c := range call {
				c()
			}

			// 计算下一次执行任务的时间
			scheduledTime = scheduledTime.Add(24 * time.Hour)

			// 重新设置定时器
			durationUntilNextRun = scheduledTime.Sub(time.Now())
			timer.Reset(durationUntilNextRun)
		}
	}()
}
