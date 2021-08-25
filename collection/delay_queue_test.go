package collection

import (
	"fmt"
	"testing"
	"time"
)

func TestDelayQueue(t *testing.T) {
	dq := NewDelayQueue("tom", 10, 2, 5, func(m map[string]interface{}) {
		for k, v := range m {
			fmt.Println(k, v)
		}
	})
	dq.Start()
	for i := 0; i < 15; i++ {
		dq.Offer(fmt.Sprintf("k%d", i), i)
	}
	time.Sleep(time.Minute)
}
