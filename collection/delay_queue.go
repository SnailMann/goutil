package collection

import (
	"log"
	"sync"
	"time"
)

type DelayQueue struct {
	Name     string
	maxSize  int
	interval int64
	timeout  int64
	modify   int64
	queue    map[string]interface{}
	action   func(map[string]interface{})
	ticker   *time.Ticker
	lock     *sync.Mutex
}

func NewDelayQueue(name string, maxSize int, interval int64, timeout int64, action func(map[string]interface{})) *DelayQueue {
	q := DelayQueue{Name: name, maxSize: maxSize * 2, interval: interval, timeout: timeout}
	q.modify = time.Now().Unix()
	q.queue = make(map[string]interface{}, maxSize)
	q.ticker = time.NewTicker(time.Duration(interval) * time.Second)
	q.lock = &sync.Mutex{}
	q.action = action
	return &q
}

func (q *DelayQueue) Offer(key string, value interface{}) {
	size := len(q.queue)
	if size >= q.maxSize {
		q.action(q.queue)
	}
	q.queue[key] = value
}

func (q *DelayQueue) Poll() {

}

func Start(q *DelayQueue) {
	handler := func() {
		for {
			select {
			case <-q.ticker.C:
				now := time.Now().Unix()
				log.Println()
				if q.modify+q.timeout < now {
					m := CopyMap(q.queue)
					for k := range m {
						delete(q.queue, k)
					}
					q.action(m)
				}
			}
		}
	}
	go handler()
}
