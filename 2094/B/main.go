// https://codeforces.com/problemset/problem/2094/B

package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n, m, l, r int
		fmt.Scan(&n, &m, &l, &r)

		var rp, lp = 0, 0
		var Inc = true
		for x := 0; x < m; x++ {
			if rp >= r {
				Inc = false
			} else if lp <= l {
				Inc = true
			}
			if Inc {
				rp++
				Inc = false
			} else {
				lp--
				Inc = true
			}
		}
		fmt.Println(lp, rp)
	}
}

