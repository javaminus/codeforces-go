// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`[[0,1],[1,0],[4,0],[0,4],[3,3],[2,4]]`, `[0,0]`, 
			`[[0,1],[1,0],[3,3]]`,
		},
		{
			`[[0,0],[1,1],[2,2],[3,4],[3,5],[4,4],[4,5]]`, `[3,3]`, 
			`[[2,2],[3,4],[4,4]]`,
		},
		{
			`[[5,6],[7,7],[2,1],[0,7],[1,6],[5,1],[3,7],[0,3],[4,0],[1,2],[6,3],[5,0],[0,4],[2,2],[1,1],[6,4],[5,4],[0,0],[2,6],[4,5],[5,2],[1,4],[7,5],[2,3],[0,5],[4,2],[1,0],[2,7],[0,1],[4,6],[6,1],[0,6],[4,3],[1,7]]`, `[3,4]`, 
			`[[2,3],[1,4],[1,6],[3,7],[4,3],[5,4],[4,5]]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, queensAttacktheKing, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-158/problems/queens-that-can-attack-the-king/
