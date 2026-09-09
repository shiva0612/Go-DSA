package main

/*
t(n) = c+2T(n-1)
hence the time complexity is 2^n
we also have to sort the sumArr -> nlogn where n is the arr size which needs to be sorted

	the array size is going to be 2^n ( n elements in a we have 2 options pick, not pick (0,1) using 2*2*2..n => 2^n sub sets we get)
	so, the sumArr size will be 2^n => for sorting [2^n * log(2^n)]

total time complexity is

	2^n for creating the sumArr +
	[2^n * log(2^n)] for sorting the sumArr

space complexity is

	we just need 2^n for storing the sumArr
*/
func subsetsum(i int, a []int, sumArr []int, curSum int) {
	if i == len(a) {
		sumArr = append(sumArr, curSum)
		return
	}

	subsetsum(i+1, a, sumArr, curSum+a[i])
	subsetsum(i, a, sumArr, curSum)
}

/*
similar to above this is also pick and non-pick => time complexity is 2^n
but, as we have to copy (ds) into (ans)
ds at level 1 has len = 1
ds at level k has len = k
so, for copying k elements into (ans) it takes extra k time
final time complexity = 2^n * k
space complexity  = 2^n * k

	at every recursion we are storing (ds) which has len=1,2,3, ... k accroding to the level
*/
func subsets() {

}

/*
 */
func combinationSum(candidates []int, target int) [][]int {
	ds := []int{}
	ans := [][]int{}
	subset_targetsum(0, candidates, &ds, target, &ans)
	return ans
}

func subset_targetsum(i int, a []int, ds *[]int, target int, ans *[][]int) {
	if target == 0 {
		*ans = append(*ans, *ds)
		return
	}
	if i == len(a) {
		return
	}

	if a[i] <= target {
		*ds = append(*ds, a[i])
		subset_targetsum(i, a, ds, target-a[i], ans)
		*ds = (*ds)[:len(*ds)-1]
	}
	subset_targetsum(i+1, a, ds, target, ans)
}
