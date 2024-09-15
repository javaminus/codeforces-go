// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1969/C
// https://codeforces.com/problemset/status/1969/problem/C
func Test_cf1969C(t *testing.T) {
	testCases := [][2]string{
		{
			`4
3 1
3 1 2
1 3
5
4 2
2 2 1 3
6 3
4 1 2 2 4 3`,
			`4
5
5
10`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1969C)
}