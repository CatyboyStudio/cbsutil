package executor

type ProxyExecutor[P any, R any, P2 any, R2 any] struct {
	target Executor[P2, R2]
	pfunc  func(P) P2
	rfunc  func(R) R2
}

// IsClose implements Executor.
func (th *ProxyExecutor[P, R, P2, R2]) IsClose() bool {
	return th.target.IsClose()
}

// Process implements Executor.
func (th *ProxyExecutor[P, R, P2, R2]) Process(p P, f func(P) (R, error)) (R, error) {
	var p2 P2
	if th.pfunc != nil {
		p2 = th.pfunc(p)
	}
	var r R
	_, err := th.target.Process(p2, func(p2 P2) (R2, error) {
		tr, terr := f(p)
		var tr2 R2
		if terr != nil {
			return tr2, terr
		}
		r = tr
		if th.rfunc != nil {
			tr2 = th.rfunc(tr)
		}
		return tr2, terr
	})
	return r, err
}

func NewProxyExecutor[P any, R any, P2 any, R2 any](exec Executor[P2, R2],
	pfunc func(P) P2, rfunc func(R) R2) *ProxyExecutor[P, R, P2, R2] {
	o := &ProxyExecutor[P, R, P2, R2]{
		target: exec,
		pfunc:  pfunc,
		rfunc:  rfunc,
	}
	return o
}
