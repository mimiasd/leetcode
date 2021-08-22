package main

func climbStairs(n int) int {
    if n == 0 || n == 1 || n == 2 {
        return n
    }

    db := make([]int, n+1)
    db[0] = 0
    db[1] = 1
    db[2] = 2

    for i := 3; i <= n; i++ {
        db[i] = db[i-1] + db[i-2]
    }

    return db[n]
}
