package main

func fib(n int) int {

	if n == 0 || n == 1 {
		return n
	}

	db := make([]int, n+1)
	db[0] = 0
	db[1] = 1

	for i := 2; i <= n; i++ {
		db[i] = db[i-1] + db[i-2]
	}

	return db[n]
}

func fib1(n int) int {

	if n == 0 || n == 1 {
		return n
	}

	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
	}

	return curr
}
