// Problem: Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.

// You may return any answer array that satisfies this condition.

// Example 1:

// Input: [3,1,2,4]
// Output: [2,4,3,1]
// The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

package main

func sortArrayByParity(A []int) []int {
	var odd_lst []int
	var even_lst []int
	for _, v := range A {
		if v%2 != 0 {
			odd_lst = append(odd_lst, v)
		} else {
			even_lst = append(even_lst, v)
		}
	}
	return append(even_lst, odd_lst...)
}

func main() {

}
