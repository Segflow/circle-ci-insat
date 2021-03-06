package main

func Fact(n uint) uint {
	return factRecusive(n)
}

func factIterative(n uint) uint {
	var r uint = 1

	for i := uint(1); i <= n; i++ {
		r = r * i
	}

	return r
}

func factRecusive(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * Fact(n-1)
}
