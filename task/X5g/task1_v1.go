package main

// Задача 1 Что нужно сделать если нам нужно запускать всего по две за раз горутины вместо всех сразу

import (
	"context"
	"fmt"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type storage struct {
	result map[int32]string
	mux    sync.Mutex
}

func LongWorker(ctx context.Context, wg *sync.WaitGroup, res *storage) {
	done := make(chan struct{}, 1)

	go func(ctx context.Context, ch chan struct{}) {
		time.Sleep(time.Duration(rand.Int31n(3)) * time.Second) // <<
		select {
		case <-ctx.Done():
			close(done)
			return
		case done <- struct{}{}:
			close(done)
		}
	}(ctx, done)

	select {
	case <-ctx.Done():
		fmt.Println("Exit 1") // <<<
		return
	case <-done:
		var key = rand.Int31n(100)
		var value = fmt.Sprintf("some result: %d", rand.Int())
		res.mux.Lock()
		res.result[key] = value
		res.mux.Unlock()
	}
}

type Semaphorer interface {
	Closer()
	Take()
	Put()
}

type Semaphore struct {
	ch chan struct{}
}

func NewSemaphore(limit int) *Semaphore {
	return &Semaphore{
		ch: make(chan struct{}, limit),
	}
}

func (s *Semaphore) Take() {
	s.ch <- struct{}{}
}

func (s *Semaphore) Put() {
	<-s.ch
}

func (s *Semaphore) Close() error {
	close(s.ch)
	return nil
}

func main() {
	max := 10 // <<<
	limit := 2
	semaphore := NewSemaphore(limit)
	store := &storage{result: make(map[int32]string, 100)}

	// WaitGroup
	wg := &sync.WaitGroup{}

	// Gracefull shatdown
	gctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Context
	ctx, cancel := context.WithTimeout(gctx, 10*time.Second)
	defer cancel()

	// Done
	done := make(chan struct{}, 1)

	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-ctx.Done():
				// break
				goto exit
			default:
			}

			wg.Add(1)
			semaphore.Take()

			go func(ctx context.Context) {
				defer func() {
					if r := recover(); r != nil {
						fmt.Println("panic")
					}
					wg.Done()
					semaphore.Put() //
				}()
				LongWorker(ctx, wg, store)
			}(ctx)
		}

	exit:

		//
		wg.Wait()
		semaphore.Close()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Result")
		return

	case <-gctx.Done():
		fmt.Println("Stop Shatdown")
	case <-ctx.Done():
		fmt.Println("Stop Timeout")
	}

	<-done // Ждем когда все горутины остановятся.
}
