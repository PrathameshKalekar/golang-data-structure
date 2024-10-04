/*
2491 You are given a positive integer array arr of even length n where skill[i] denotes the skill of the ith player. Divide the players into n / 2 teams of size 2 such that the total skill of each team is equal.

The chemistry of a team is equal to the product of the skills of the players on that team.

Return the sum of the chemistry of all the teams, or return -1 if there is no way to divide the players into teams such that the total skill of each team is equal.

Example 1:

Input: skill = [3,2,5,1,3,4]
Output: 22
Explanation:
Divide the players into the following teams: (1, 5), (2, 4), (3, 3), where each team has a total skill of 6.
The sum of the chemistry of all the teams is: 1 * 5 + 2 * 4 + 3 * 3 = 5 + 8 + 9 = 22.
Example 2:

Input: skill = [3,4]
Output: 12
Explanation:
The two players form a team with a total skill of 7.
The chemistry of the team is 3 * 4 = 12.
Example 3:

Input: skill = [1,1,2,3]
Output: -1
Explanation:
There is no way to divide the players into teams such that the total skill of each team is equal.
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 1, 2, 3}
	answer := dividePlayersIntoTeamsOfEqualSkill(arr)
	fmt.Println(answer)
}

func dividePlayersIntoTeamsOfEqualSkill(arr []int) int64 {
	sort.Ints(arr)

	ans := 0
	l := len(arr)
	sum := arr[0] + arr[l-1]

	for i := 0; i <= (l/2)-1; i++ {
		if arr[i]+arr[l-i-1] != sum {
			return int64(-1)
		}
		ans += arr[i] * arr[l-i-1]
	}

	return int64(ans)
}
