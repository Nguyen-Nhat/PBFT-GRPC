package math

// TernaryOp represent for e1 ? e2 : e3
func TernaryOp[T any](condition bool, ifOutput T, elseOutput T) T {
	if condition {
		return ifOutput
	}
	return elseOutput
}
