package queue

import (
	"context"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewConcurrentBlockingQueue[int](10)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//_, err := q.DeQueue(ctx)
	//assert.Equal(t, context.DeadlineExceeded, err)
	q.EnQueue(context.Background(), 1)
	q.EnQueue(context.Background(), 2)
	q.EnQueue(context.Background(), 3)
}
