package pkg

func To[T any](condition bool, t1 T, t2 T) T {
	if condition {
		return t1
	}
	return t2
}

func Do[T any](condition bool, t1 func() T, t2 func() T) T {
	if condition {
		return t1()
	}
	return t2()
}
