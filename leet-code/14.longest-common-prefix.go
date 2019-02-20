/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (32.93%)
 * Total Accepted:    404.4K
 * Total Submissions: 1.2M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 *
 * If there is no common prefix, return an empty string "".
 *
 * Example 1:
 *
 *
 * Input: ["flower","flow","flight"]
 * Output: "fl"
 *
 *
 * Example 2:
 *
 *
 * Input: ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 *
 *
 * Note:
 *
 * All given inputs are in lowercase letters a-z.
 *
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || strs == nil {
		return ""
	}

	l := len(strs)

	if l == 1 {
		return strs[0]
	}

	strs1 := make([]string, l/2)
	strs2 := make([]string, l-(l/2))

	for i := 0; i < l/2; i++ {
		strs1[i] = strs[i]
	}

	for i := l / 2; i < l; i++ {
		strs2[i-(l/2)] = strs[i]
	}

	s1 := longestCommonPrefix(strs1)
	s2 := longestCommonPrefix(strs2)
	i := 0
	for i < len(s1) && i < len(s2) && s1[i] == s2[i] {
		i++
	}

	return s1[0:i]

}
