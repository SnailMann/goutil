package bloom

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestFilter(t *testing.T) {
	filter := New(100, 0.0001)
	for i := 0; i < 10; i++ {
		go filter.Add([]byte(strconv.Itoa(i)))
	}
	time.Sleep(time.Second)
	for i := 0; i < 20; i++ {
		b := filter.Contain([]byte(strconv.Itoa(i)))
		fmt.Printf("%d contains %v\n", i, b)
	}
}
