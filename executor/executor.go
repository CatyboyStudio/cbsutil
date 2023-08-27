package executor

import (
	"errors"
	"fmt"
	"log/slog"
	"runtime/debug"
	"sync/atomic"
)

var ErrClosed = errors.New("closed")

type Executor[P any, R any] interface {
	Process(P, func(P) (R, error)) (R, error)

	IsClose() bool
}

type Task[P any, R any] struct {
	Param P
	Func  func(P) (R, error)
}

func (o Task[P, R]) Execute() (r R, err error) {
	defer func() {
		erro := recover()
		if erro != nil {
			s := string(debug.Stack())
			slog.Error("exec error", "erro", erro, "stack", s)
			if e, ok := erro.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", erro)
			}
		}
	}()
	r, err = o.Func(o.Param)
	return
}

type Result[R any] struct {
	Data R
	Err  error
}

func NewResult[R any](v R, err error) Result[R] {
	return Result[R]{v, err}
}

// Inline
type Inline[P any, R any] int

// IsClose implements Executor.
func (Inline[P, R]) IsClose() bool {
	return false
}

// Process implements Executor.
func (Inline[P, R]) Process(p P, f func(P) (R, error)) (R, error) {
	return f(p)
}

func NewInline[P any, R any]() Executor[P, R] {
	return Inline[P, R](0)
}

// Single Queue
type SingleQueueExecutor[P any, R any] struct {
	ch     chan Task[P, R]
	wch    chan Result[R]
	statue int32
}

// IsClose implements Executor.
func (th *SingleQueueExecutor[P, R]) IsClose() bool {
	return atomic.LoadInt32(&th.statue) < 0
}

// Process implements Executor.
func (th *SingleQueueExecutor[P, R]) Process(p P, f func(P) (R, error)) (R, error) {
	req := Task[P, R]{p, f}
	if atomic.LoadInt32(&th.statue) < 0 {
		var def R
		return def, ErrClosed
	}
	th.ch <- req
	r := <-th.wch
	return r.Data, r.Err
}

func NewSingleQueue[P any, R any]() *SingleQueueExecutor[P, R] {
	o := &SingleQueueExecutor[P, R]{
		ch:  make(chan Task[P, R]),
		wch: make(chan Result[R]),
	}
	// var _ (Executor[P, R]) = o
	go o.run()
	return o
}

func (th *SingleQueueExecutor[P, R]) run() {
	for t := range th.ch {
		r, err := t.Execute()
		th.wch <- Result[R]{r, err}
	}
}

func (th *SingleQueueExecutor[P, R]) Close() {
	atomic.StoreInt32(&th.statue, -1)
}
