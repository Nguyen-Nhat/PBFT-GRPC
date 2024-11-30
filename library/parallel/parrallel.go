package parallel

import (
	"context"
	"errors"
	"sync"
)

const (
	defaultLimit = 10
)

var (
	RegisterCompletedParallelError = errors.New("parallel: registering on completed parallel")
)

type option struct {
	ignoreError   bool
	parallelLimit uint
}

type Option = func(*option)

func WithIngoreError() func(*option) {
	return func(o *option) {
		o.ignoreError = true
	}
}

func WithLimit(limit uint) func(*option) {
	return func(o *option) {
		if o.parallelLimit <= 1 {
			o.parallelLimit = defaultLimit
			return
		}
		o.parallelLimit = limit
	}
}

func New(ctx context.Context, opts ...func(*option)) *Parallel {

	option := &option{
		parallelLimit: defaultLimit,
	}
	for _, opt := range opts {
		opt(option)
	}

	ctx, cancel := context.WithCancel(ctx)
	p := &Parallel{
		ctx:         ctx,
		wg:          &sync.WaitGroup{},
		jobsChannel: make(chan *job),
		cancel:      cancel,
		guard:       make(chan struct{}, option.parallelLimit),
	}
	if !option.ignoreError {
		go p.worker()
	} else {
		go p.workerIgnoreError()
	}
	return p
}

type Func interface {
	Do(context.Context) error
}

type job struct {
	fn Func
}

type Parallel struct {
	ctx         context.Context
	wg          *sync.WaitGroup
	jobsChannel chan *job
	cancel      context.CancelFunc
	errLock     sync.RWMutex
	err         error
	lock        sync.RWMutex
	complete    bool
	guard       chan struct{}
}

func (p *Parallel) Register(fns ...Func) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.complete {
		p.setErrorIfNil(RegisterCompletedParallelError)
		return RegisterCompletedParallelError
	}
	for _, fn := range fns {
		p.wg.Add(1)
		p.jobsChannel <- &job{
			fn: fn,
		}
	}
	return nil
}

func (p *Parallel) Wait() error {
	p.wg.Wait()

	p.lock.Lock()
	defer p.lock.Unlock()

	p.complete = true
	close(p.jobsChannel)
	close(p.guard)
	p.cancel()
	return p.getError()
}

func (p *Parallel) workerIgnoreError() {
	for j := range p.jobsChannel {
		p.allocateGuard()
		go func(job *job) {
			defer p.wg.Done()
			defer p.releaseGuard()
			job.fn.Do(p.ctx)
		}(j)
	}
}

func (p *Parallel) worker() {
	for j := range p.jobsChannel {
		p.allocateGuard()
		go func(job *job) {
			defer p.wg.Done()
			defer p.releaseGuard()

			select {
			case <-p.ctx.Done():
				if p.ctx.Err() != nil {
					p.setErrorIfNil(p.ctx.Err())
				}
			default:
				errR := job.fn.Do(p.ctx)
				if errR != nil {
					p.setErrorIfNil(errR)
					p.cancel()
				}
			}
		}(j)
	}
}

func (p *Parallel) setErrorIfNil(err error) {
	p.errLock.Lock()
	defer p.errLock.Unlock()

	if p.err == nil {
		p.err = err
	}
}

func (p *Parallel) getError() error {
	p.errLock.RLock()
	defer p.errLock.RUnlock()

	return p.err
}

func (p *Parallel) allocateGuard() {
	p.guard <- struct{}{}
}

func (p *Parallel) releaseGuard() {
	<-p.guard
}
