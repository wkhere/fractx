package fractx

func partial[T1, T2, T3 any](f func(T1, T2) T3, v T1) func(T2) T3 {
	return func(x T2) T3 { return f(v, x) }
}
