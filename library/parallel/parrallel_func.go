package parallel

import "context"

func NewFunc[TReq, TRes any](req TReq, fn func(ctx context.Context, p TReq) (TRes, error)) *ParallelFunc[TReq, TRes] {
	return &ParallelFunc[TReq, TRes]{
		req: req,
		fn:  fn,
	}
}

// ParallelFunc define a generic struct holding parallel function and its result.
type ParallelFunc[TReq, TRes any] struct {
	req    TReq
	result TRes
	err    error
	fn     func(ctx context.Context, p TReq) (TRes, error)
}

func (p *ParallelFunc[TReq, TRes]) Do(ctx context.Context) error {
	res, err := p.fn(ctx, p.req)
	p.result = res
	p.err = err
	return err
}

func (p *ParallelFunc[TReq, TRes]) GetRequest() TReq {
	return p.req
}

func (p *ParallelFunc[TReq, TRes]) GetResult() (TRes, error) {
	return p.result, p.err
}
