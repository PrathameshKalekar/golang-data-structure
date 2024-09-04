/*
874 A robot on an infinite XY-plane starts at point (0, 0) facing north. The robot can receive a sequence of these three possible types of commands:

-2: Turn left 90 degrees.
-1: Turn right 90 degrees.
1 <= k <= 9: Move forward k units, one unit at a time.
Some of the grid squares are obstacles. The ith obstacle is at grid point obstacles[i] = (xi, yi). If the robot runs into an obstacle, then it will instead stay in its current location and move on to the next command.

Return the maximum Euclidean distance that the robot ever gets from the origin squared (i.e. if the distance is 5, return 25).

Note:

North means +Y direction.
East means +X direction.
South means -Y direction.
West means -X direction.
There can be obstacle in [0,0].

Example 1:

Input: commands = [4,-1,3], obstacles = []
Output: 25
Explanation: The robot starts at (0, 0):
1. Move north 4 units to (0, 4).
2. Turn right.
3. Move east 3 units to (3, 4).
The furthest point the robot ever gets from the origin is (3, 4), which squared is 32 + 42 = 25 units away.
Example 2:

Input: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
Output: 65
Explanation: The robot starts at (0, 0):
1. Move north 4 units to (0, 4).
2. Turn right.
3. Move east 1 unit and get blocked by the obstacle at (2, 4), robot is at (1, 4).
4. Turn left.
5. Move north 4 units to (1, 8).
The furthest point the robot ever gets from the origin is (1, 8), which squared is 12 + 82 = 65 units away.
Example 3:

Input: commands = [6,-1,-1,6], obstacles = []
Output: 36
Explanation: The robot starts at (0, 0):
1. Move north 6 units to (0, 6).
2. Turn right.
3. Turn right.
4. Move south 6 units to (0, 0).
The furthest point the robot ever gets from the origin is (0, 6), which squared is 62 = 36 units away.
*/
package main

import (
	"fmt"
	"math"
)

func robotSim(commands []int, obstacles [][]int) int {
	// Directions: North, East, South, West
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}

	// Convert obstacles into a set for quick lookup
	obstacleSet := make(map[[2]int]bool)
	for _, obs := range obstacles {
		obstacleSet[[2]int{obs[0], obs[1]}] = true
	}

	// Initial position and direction (facing North)
	x, y, direction := 0, 0, 0
	maxDistSquared := 0

	for _, cmd := range commands {
		if cmd == -2 {
			// Turn left
			direction = (direction + 3) % 4
		} else if cmd == -1 {
			// Turn right
			direction = (direction + 1) % 4
		} else {
			// Move forward k steps
			for k := 0; k < cmd; k++ {
				newX, newY := x+dx[direction], y+dy[direction]
				if obstacleSet[[2]int{newX, newY}] {
					// Stop moving if there's an obstacle
					break
				}
				// Update position
				x, y = newX, newY
				// Update the max distance squared
				maxDistSquared = int(math.Max(float64(maxDistSquared), float64(x*x+y*y)))
			}
		}
	}

	return maxDistSquared
}

func main() {
	fmt.Println(robotSim([]int{4, -1, 3}, [][]int{}))              // Output: 25
	fmt.Println(robotSim([]int{4, -1, 4, -2, 4}, [][]int{{2, 4}})) // Output: 65
}
