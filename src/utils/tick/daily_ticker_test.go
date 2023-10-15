package tick

import "testing"

func TestNineAMTasks(t *testing.T) {
	DailyTicker(22, 19, 0, func() {
		println("触发定时器")
	})
	select {}
}
