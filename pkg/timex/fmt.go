package timex

import (
	"fmt"
	"time"
)

func Println(cost time.Duration, label string) {
	fmt.Printf("%s耗时：%v(秒)， %v(毫秒)，%v(微秒)，%v(纳秒)\n",
		label, cost.Seconds(), cost.Milliseconds(), cost.Microseconds(), cost.Nanoseconds())
}
