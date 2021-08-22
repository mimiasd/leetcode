package main

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	db := make([]int, n+1)
	db[0] = 0
	db[1] = 1
	db[2] = 1

	for i := 3; i <= n; i++ {
		db[i] = db[i-1] + db[i-2] + db[i-3]
	}

	return db[n]
}
