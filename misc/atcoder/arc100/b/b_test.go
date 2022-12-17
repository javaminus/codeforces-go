// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/arc100/submit?taskScreenName=arc100_b
func Test_run(t *testing.T) {
	t.Log("Current test is [b]")
	testCases := [][2]string{
		{
			`5
3 2 4 1 2`,
			`2`,
		},
		{
			`10
10 71 84 33 6 47 23 25 52 64`,
			`36`,
		},
		{
			`7
1 2 3 1000000000 4 5 6`,
			`999999994`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/abc102/tasks/arc100_b