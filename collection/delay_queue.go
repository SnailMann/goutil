package collection

import (
	"sync"
	"time"
)

type DelayQueue struct {
	Name     string
	maxSize  int
	interval  int64
	delaytime int64
	modify    int64
	queue    map[string]interface{}
	action   func(map[string]interface{})
	ticker   *time.Ticker
	lock     *sync.Mutex
}

func NewDelayQueue(name string, maxSize int, interval int64, delaytime int64, action func(map[string]interface{})) *DelayQueue {
	q := DelayQueue{Name: name, maxSize: maxSize, interval: interval, delaytime: delaytime}
	q.modify = time.Now().Unix()
	q.queue = make(map[string]interface{}, maxSize * 2)
	q.ticker = time.NewTicker(time.Duration(interval) * time.Second)
	q.lock = &sync.Mutex{}
	q.action = action
	return &q
}

func (q *DelayQueue) Offer(key string, value interface{}) {
	size := len(q.queue)
	if size >= q.maxSize {
		m := q.remove()
		q.action(m)
	}
	q.queue[key] = value
}

func (q *DelayQueue) remove() map[string]interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	m := CopyMap(q.queue)
	for k := range m {
		delete(q.queue, k)
	}
	return m;
}

func (q *DelayQueue) Start() {
	handler := func() {
		for {
			select {
			case <-q.ticker.C:
				now := time.Now().Unix()
				if q.modify+q.delaytime < now {
					if len(q.queue) == 0 {
						return
					}
					m := q.remove()
					q.action(m)
				}
			}
		}
	}
	go handler()
}
